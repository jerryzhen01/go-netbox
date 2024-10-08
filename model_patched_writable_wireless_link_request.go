/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableWirelessLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableWirelessLinkRequest{}

// PatchedWritableWirelessLinkRequest Adds support for custom fields and tags.
type PatchedWritableWirelessLinkRequest struct {
	InterfaceA           *InterfaceRequest                  `json:"interface_a,omitempty"`
	InterfaceB           *InterfaceRequest                  `json:"interface_b,omitempty"`
	Ssid                 *string                            `json:"ssid,omitempty"`
	Status               *PatchedWritableCableRequestStatus `json:"status,omitempty"`
	Tenant               NullableTenantRequest              `json:"tenant,omitempty"`
	AuthType             *AuthenticationType1               `json:"auth_type,omitempty"`
	AuthCipher           *AuthenticationCipher              `json:"auth_cipher,omitempty"`
	AuthPsk              *string                            `json:"auth_psk,omitempty"`
	Description          *string                            `json:"description,omitempty"`
	Comments             *string                            `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                 `json:"tags,omitempty"`
	CustomFields         map[string]interface{}             `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableWirelessLinkRequest PatchedWritableWirelessLinkRequest

// NewPatchedWritableWirelessLinkRequest instantiates a new PatchedWritableWirelessLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableWirelessLinkRequest() *PatchedWritableWirelessLinkRequest {
	this := PatchedWritableWirelessLinkRequest{}
	return &this
}

// NewPatchedWritableWirelessLinkRequestWithDefaults instantiates a new PatchedWritableWirelessLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableWirelessLinkRequestWithDefaults() *PatchedWritableWirelessLinkRequest {
	this := PatchedWritableWirelessLinkRequest{}
	return &this
}

// GetInterfaceA returns the InterfaceA field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetInterfaceA() InterfaceRequest {
	if o == nil || IsNil(o.InterfaceA) {
		var ret InterfaceRequest
		return ret
	}
	return *o.InterfaceA
}

// GetInterfaceAOk returns a tuple with the InterfaceA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetInterfaceAOk() (*InterfaceRequest, bool) {
	if o == nil || IsNil(o.InterfaceA) {
		return nil, false
	}
	return o.InterfaceA, true
}

// HasInterfaceA returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasInterfaceA() bool {
	if o != nil && !IsNil(o.InterfaceA) {
		return true
	}

	return false
}

// SetInterfaceA gets a reference to the given InterfaceRequest and assigns it to the InterfaceA field.
func (o *PatchedWritableWirelessLinkRequest) SetInterfaceA(v InterfaceRequest) {
	o.InterfaceA = &v
}

// GetInterfaceB returns the InterfaceB field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetInterfaceB() InterfaceRequest {
	if o == nil || IsNil(o.InterfaceB) {
		var ret InterfaceRequest
		return ret
	}
	return *o.InterfaceB
}

// GetInterfaceBOk returns a tuple with the InterfaceB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetInterfaceBOk() (*InterfaceRequest, bool) {
	if o == nil || IsNil(o.InterfaceB) {
		return nil, false
	}
	return o.InterfaceB, true
}

// HasInterfaceB returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasInterfaceB() bool {
	if o != nil && !IsNil(o.InterfaceB) {
		return true
	}

	return false
}

// SetInterfaceB gets a reference to the given InterfaceRequest and assigns it to the InterfaceB field.
func (o *PatchedWritableWirelessLinkRequest) SetInterfaceB(v InterfaceRequest) {
	o.InterfaceB = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *PatchedWritableWirelessLinkRequest) SetSsid(v string) {
	o.Ssid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetStatus() PatchedWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetStatusOk() (*PatchedWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedWritableWirelessLinkRequest) SetStatus(v PatchedWritableCableRequestStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableWirelessLinkRequest) GetTenant() TenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret TenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableWirelessLinkRequest) GetTenantOk() (*TenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableTenantRequest and assigns it to the Tenant field.
func (o *PatchedWritableWirelessLinkRequest) SetTenant(v TenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedWritableWirelessLinkRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedWritableWirelessLinkRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetAuthType() AuthenticationType1 {
	if o == nil || IsNil(o.AuthType) {
		var ret AuthenticationType1
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetAuthTypeOk() (*AuthenticationType1, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthenticationType1 and assigns it to the AuthType field.
func (o *PatchedWritableWirelessLinkRequest) SetAuthType(v AuthenticationType1) {
	o.AuthType = &v
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetAuthCipher() AuthenticationCipher {
	if o == nil || IsNil(o.AuthCipher) {
		var ret AuthenticationCipher
		return ret
	}
	return *o.AuthCipher
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetAuthCipherOk() (*AuthenticationCipher, bool) {
	if o == nil || IsNil(o.AuthCipher) {
		return nil, false
	}
	return o.AuthCipher, true
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasAuthCipher() bool {
	if o != nil && !IsNil(o.AuthCipher) {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given AuthenticationCipher and assigns it to the AuthCipher field.
func (o *PatchedWritableWirelessLinkRequest) SetAuthCipher(v AuthenticationCipher) {
	o.AuthCipher = &v
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *PatchedWritableWirelessLinkRequest) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableWirelessLinkRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableWirelessLinkRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableWirelessLinkRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableWirelessLinkRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableWirelessLinkRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableWirelessLinkRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableWirelessLinkRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableWirelessLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableWirelessLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceA) {
		toSerialize["interface_a"] = o.InterfaceA
	}
	if !IsNil(o.InterfaceB) {
		toSerialize["interface_b"] = o.InterfaceB
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthCipher) {
		toSerialize["auth_cipher"] = o.AuthCipher
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableWirelessLinkRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableWirelessLinkRequest := _PatchedWritableWirelessLinkRequest{}

	err = json.Unmarshal(data, &varPatchedWritableWirelessLinkRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableWirelessLinkRequest(varPatchedWritableWirelessLinkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interface_a")
		delete(additionalProperties, "interface_b")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_cipher")
		delete(additionalProperties, "auth_psk")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableWirelessLinkRequest struct {
	value *PatchedWritableWirelessLinkRequest
	isSet bool
}

func (v NullablePatchedWritableWirelessLinkRequest) Get() *PatchedWritableWirelessLinkRequest {
	return v.value
}

func (v *NullablePatchedWritableWirelessLinkRequest) Set(val *PatchedWritableWirelessLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableWirelessLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableWirelessLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableWirelessLinkRequest(val *PatchedWritableWirelessLinkRequest) *NullablePatchedWritableWirelessLinkRequest {
	return &NullablePatchedWritableWirelessLinkRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableWirelessLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableWirelessLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
