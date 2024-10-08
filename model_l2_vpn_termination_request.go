/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.9 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// checks if the L2VPNTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2VPNTerminationRequest{}

// L2VPNTerminationRequest Adds support for custom fields and tags.
type L2VPNTerminationRequest struct {
	L2vpn                L2VPNRequest `json:"l2vpn"`
	AdditionalProperties map[string]interface{}
}

type _L2VPNTerminationRequest L2VPNTerminationRequest

// NewL2VPNTerminationRequest instantiates a new L2VPNTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2VPNTerminationRequest(l2vpn L2VPNRequest) *L2VPNTerminationRequest {
	this := L2VPNTerminationRequest{}
	this.L2vpn = l2vpn
	return &this
}

// NewL2VPNTerminationRequestWithDefaults instantiates a new L2VPNTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2VPNTerminationRequestWithDefaults() *L2VPNTerminationRequest {
	this := L2VPNTerminationRequest{}
	return &this
}

// GetL2vpn returns the L2vpn field value
func (o *L2VPNTerminationRequest) GetL2vpn() L2VPNRequest {
	if o == nil {
		var ret L2VPNRequest
		return ret
	}

	return o.L2vpn
}

// GetL2vpnOk returns a tuple with the L2vpn field value
// and a boolean to check if the value has been set.
func (o *L2VPNTerminationRequest) GetL2vpnOk() (*L2VPNRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2vpn, true
}

// SetL2vpn sets field value
func (o *L2VPNTerminationRequest) SetL2vpn(v L2VPNRequest) {
	o.L2vpn = v
}

func (o L2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2VPNTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["l2vpn"] = o.L2vpn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *L2VPNTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"l2vpn",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varL2VPNTerminationRequest := _L2VPNTerminationRequest{}

	err = json.Unmarshal(data, &varL2VPNTerminationRequest)

	if err != nil {
		return err
	}

	*o = L2VPNTerminationRequest(varL2VPNTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "l2vpn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableL2VPNTerminationRequest struct {
	value *L2VPNTerminationRequest
	isSet bool
}

func (v NullableL2VPNTerminationRequest) Get() *L2VPNTerminationRequest {
	return v.value
}

func (v *NullableL2VPNTerminationRequest) Set(val *L2VPNTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableL2VPNTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableL2VPNTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2VPNTerminationRequest(val *L2VPNTerminationRequest) *NullableL2VPNTerminationRequest {
	return &NullableL2VPNTerminationRequest{value: val, isSet: true}
}

func (v NullableL2VPNTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2VPNTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
