package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch/v5"
)

var (
	flagAddr       = flag.String("addr", "", "The addr to listen")
	flagPort       = flag.Int("port", 8080, "The port to listen")
	flagApiVersion = flag.String("api-version", "2022-09-01", "The API version of the metadata endpoint")
	flagMetadata   = flag.String("metadata", "{}", "The JSON object contains the expected change (in form of JSON patch) on top of the metadata")
	flagCertFile   = flag.String("cert", "localhost.pem", "The certificate file")
	flagKeyFile    = flag.String("key", "localhost-key.pem", "The private key file")
	flagOriginHost = flag.String("origin", "https://management.azure.com", "The origin hostname of the metadata endpoint that is proxied to")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *flagAddr, *flagPort)
	outboundURL, err := url.Parse(*flagOriginHost)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	proxy := httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(outboundURL)
		},
		ModifyResponse: func(r *http.Response) error {
			if r.Request.URL.Path != "/metadata/endpoints" {
				return nil
			}
			apiVersions := r.Request.URL.Query()["api-version"]
			if len(apiVersions) != 1 {
				return nil
			}
			if apiVersions[0] != *flagApiVersion {
				return nil
			}

			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				return fmt.Errorf("new gzip reader: %v", err)
			}

			b, err := io.ReadAll(reader)
			if err != nil {
				return fmt.Errorf("reading the raw response body: %v", err)
			}
			r.Body.Close()

			// Trim away a BOM if present
			b = bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))

			b, err = jsonpatch.MergePatch(b, []byte(*flagMetadata))
			if err != nil {
				return fmt.Errorf("patching the respnse: %v", err)
			}

			var buf bytes.Buffer
			writer := gzip.NewWriter(&buf)
			if _, err := writer.Write(b); err != nil {
				return fmt.Errorf("writing the patched response: %v", err)
			}
			if err := writer.Close(); err != nil {
				return fmt.Errorf("close after writing the patched response: %v", err)
			}
			r.Header.Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
			r.Body = io.NopCloser(&buf)

			return nil
		},
	}

	server := &http.Server{Addr: addr, Handler: &proxy}

	fmt.Printf("Listening at %s\n", addr)
	if err := server.ListenAndServeTLS(*flagCertFile, *flagKeyFile); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
