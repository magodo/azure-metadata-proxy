package main

// A PIT copy from https://github.com/hashicorp/go-azure-sdk/blob/af8da8d2759751639b76b9c8c216d4918b07bccd/sdk/internal/metadata/client.go#L163
type metaDataResponse struct {
	Portal         string `json:"portal"`
	Authentication struct {
		LoginEndpoint    string   `json:"loginEndpoint"`
		Audiences        []string `json:"audiences"`
		Tenant           string   `json:"tenant"`
		IdentityProvider string   `json:"identityProvider"`
	} `json:"authentication"`
	Media         string `json:"media"`
	GraphAudience string `json:"graphAudience"`
	Graph         string `json:"graph"`
	Name          string `json:"name"`
	Suffixes      struct {
		AzureDataLakeStoreFileSystem        string `json:"azureDataLakeStoreFileSystem"`
		AcrLoginServer                      string `json:"acrLoginServer"`
		SqlServerHostname                   string `json:"sqlServerHostname"`
		AzureDataLakeAnalyticsCatalogAndJob string `json:"azureDataLakeAnalyticsCatalogAndJob"`
		KeyVaultDns                         string `json:"keyVaultDns"`
		Storage                             string `json:"storage"`
		AzureFrontDoorEndpointSuffix        string `json:"azureFrontDoorEndpointSuffix"`
		StorageSyncEndpointSuffix           string `json:"storageSyncEndpointSuffix"`
		MhsmDns                             string `json:"mhsmDns"`
		MysqlServerEndpoint                 string `json:"mysqlServerEndpoint"`
		PostgresqlServerEndpoint            string `json:"postgresqlServerEndpoint"`
		MariadbServerEndpoint               string `json:"mariadbServerEndpoint"`
		SynapseAnalytics                    string `json:"synapseAnalytics"`
		AttestationEndpoint                 string `json:"attestationEndpoint"`
	} `json:"suffixes"`
	Batch                                 string `json:"batch"`
	ResourceManager                       string `json:"resourceManager"`
	VmImageAliasDoc                       string `json:"vmImageAliasDoc"`
	ActiveDirectoryDataLake               string `json:"activeDirectoryDataLake"`
	SqlManagement                         string `json:"sqlManagement"`
	MicrosoftGraphResourceId              string `json:"microsoftGraphResourceId"`
	AppInsightsResourceId                 string `json:"appInsightsResourceId"`
	AppInsightsTelemetryChannelResourceId string `json:"appInsightsTelemetryChannelResourceId"`
	AttestationResourceId                 string `json:"attestationResourceId"`
	SynapseAnalyticsResourceId            string `json:"synapseAnalyticsResourceId"`
	LogAnalyticsResourceId                string `json:"logAnalyticsResourceId"`
	OssrDbmsResourceId                    string `json:"ossrDbmsResourceId"`
	Gallery                               string `json:"gallery"`
}
