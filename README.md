# Azure Metadata Proxy

A reverse proxy that aims to manipulate the Azure metadata endpoint (e.g. [https://management.azure.com/metadata/endpoints?api-version=2022-09-01](https://management.azure.com/metadata/endpoints?api-version=2022-09-01)) response for a certain API version.

The main reason is to allow users to change the `resourceManager` endpoint in the response from a global endpoint to a regional one (See example below).

**NOTE**: The functionality provided by this tool can be achieved by tools like `mitmproxy` as well.

## Install

`go install github.com/magodo/azure-metadata-host@main`

## Precondition

The proxy will run via TLS, which means you'll need to setup CA for the running host (e.g. localhost). Using [`mkcert`](https://github.com/FiloSottile/mkcert) can acheive this easily.

## Example

This example shows how to manipulate the response to change the `resourceManager` to a regional endpoint, in context of Terraform AzureRM provider. 

1. Open a terminal to run the proxy first:

    ```shell
    azure-metadata-proxy -port 9999 -metadata='{"resourceManager": "https://westeurope.management.azure.com/"}'
    ```

2. Run the provider by setting the `ARM_METADATA_HOSTNAME`:

    ```shell
    ARM_METADATA_HOSTNAME=localhost:9999 terraform plan
    ```
