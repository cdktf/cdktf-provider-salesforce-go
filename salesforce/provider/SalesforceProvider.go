package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-salesforce-go/salesforce/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-salesforce-go/salesforce/v4/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs salesforce}.
type SalesforceProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiVersion() *string
	SetApiVersion(val *string)
	ApiVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	LoginUrl() *string
	SetLoginUrl(val *string)
	LoginUrlInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiVersion()
	ResetClientId()
	ResetLoginUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SalesforceProvider
type jsiiProxy_SalesforceProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_SalesforceProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) ApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) LoginUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) LoginUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SalesforceProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs salesforce} Resource.
func NewSalesforceProvider(scope constructs.Construct, id *string, config *SalesforceProviderConfig) SalesforceProvider {
	_init_.Initialize()

	if err := validateNewSalesforceProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SalesforceProvider{}

	_jsii_.Create(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0/docs salesforce} Resource.
func NewSalesforceProvider_Override(s SalesforceProvider, scope constructs.Construct, id *string, config *SalesforceProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetApiVersion(val *string) {
	_jsii_.Set(
		j,
		"apiVersion",
		val,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetLoginUrl(val *string) {
	_jsii_.Set(
		j,
		"loginUrl",
		val,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_SalesforceProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SalesforceProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSalesforceProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SalesforceProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSalesforceProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SalesforceProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSalesforceProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SalesforceProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-salesforce.provider.SalesforceProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SalesforceProvider) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SalesforceProvider) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		s,
		"resetAlias",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetApiVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetApiVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetLoginUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetLoginUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SalesforceProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SalesforceProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SalesforceProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SalesforceProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

