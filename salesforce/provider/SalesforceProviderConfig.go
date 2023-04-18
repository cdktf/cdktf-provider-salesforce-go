package provider


type SalesforceProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#alias SalesforceProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// API version of the salesforce org in the format in the format: MAJOR.MINOR (please omit any leading 'v'). The provider requires at least version 53.0. Can be specified with the environment variable SALESFORCE_API_VERSION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#api_version SalesforceProvider#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Client ID of the connected app.
	//
	// Corresponds to Consumer Key in the user interface. Can be specified with the environment variable SALESFORCE_CLIENT_ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#client_id SalesforceProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Directs the authentication request, defaults to the production endpoint https://login.salesforce.com, should be set to https://test.salesforce.com for sandbox organizations. Can be specified with the environment variable SALESFORCE_LOGIN_URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#login_url SalesforceProvider#login_url}
	LoginUrl *string `field:"optional" json:"loginUrl" yaml:"loginUrl"`
	// Private Key associated to the public certificate that was uploaded to the connected app.
	//
	// This may point to a file location or be set directly. This should not be confused with the Consumer Secret in the user interface. Can be specified with the environment variable SALESFORCE_PRIVATE_KEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#private_key SalesforceProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Salesforce Username of a System Administrator like user for the provider to authenticate as.
	//
	// Can be specified with the environment variable SALESFORCE_USERNAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs#username SalesforceProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

