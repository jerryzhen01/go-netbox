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

// IKEPolicyModeValue * `aggressive` - Aggressive * `main` - Main
type IKEPolicyModeValue string

// List of IKEPolicy_mode_value
const (
	IKEPOLICYMODEVALUE_AGGRESSIVE IKEPolicyModeValue = "aggressive"
	IKEPOLICYMODEVALUE_MAIN       IKEPolicyModeValue = "main"
)

// All allowed values of IKEPolicyModeValue enum
var AllowedIKEPolicyModeValueEnumValues = []IKEPolicyModeValue{
	"aggressive",
	"main",
}

func (v *IKEPolicyModeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEPolicyModeValue(value)
	for _, existing := range AllowedIKEPolicyModeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEPolicyModeValue", value)
}

// NewIKEPolicyModeValueFromValue returns a pointer to a valid IKEPolicyModeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEPolicyModeValueFromValue(v string) (*IKEPolicyModeValue, error) {
	ev := IKEPolicyModeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEPolicyModeValue: valid values are %v", v, AllowedIKEPolicyModeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEPolicyModeValue) IsValid() bool {
	for _, existing := range AllowedIKEPolicyModeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEPolicy_mode_value value
func (v IKEPolicyModeValue) Ptr() *IKEPolicyModeValue {
	return &v
}

type NullableIKEPolicyModeValue struct {
	value *IKEPolicyModeValue
	isSet bool
}

func (v NullableIKEPolicyModeValue) Get() *IKEPolicyModeValue {
	return v.value
}

func (v *NullableIKEPolicyModeValue) Set(val *IKEPolicyModeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicyModeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicyModeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicyModeValue(val *IKEPolicyModeValue) *NullableIKEPolicyModeValue {
	return &NullableIKEPolicyModeValue{value: val, isSet: true}
}

func (v NullableIKEPolicyModeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicyModeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
