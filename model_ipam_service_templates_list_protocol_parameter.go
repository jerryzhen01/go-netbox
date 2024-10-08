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

// IpamServiceTemplatesListProtocolParameter the model 'IpamServiceTemplatesListProtocolParameter'
type IpamServiceTemplatesListProtocolParameter string

// List of ipam_service_templates_list_protocol_parameter
const (
	IPAMSERVICETEMPLATESLISTPROTOCOLPARAMETER_SCTP IpamServiceTemplatesListProtocolParameter = "sctp"
	IPAMSERVICETEMPLATESLISTPROTOCOLPARAMETER_TCP  IpamServiceTemplatesListProtocolParameter = "tcp"
	IPAMSERVICETEMPLATESLISTPROTOCOLPARAMETER_UDP  IpamServiceTemplatesListProtocolParameter = "udp"
)

// All allowed values of IpamServiceTemplatesListProtocolParameter enum
var AllowedIpamServiceTemplatesListProtocolParameterEnumValues = []IpamServiceTemplatesListProtocolParameter{
	"sctp",
	"tcp",
	"udp",
}

func (v *IpamServiceTemplatesListProtocolParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpamServiceTemplatesListProtocolParameter(value)
	for _, existing := range AllowedIpamServiceTemplatesListProtocolParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpamServiceTemplatesListProtocolParameter", value)
}

// NewIpamServiceTemplatesListProtocolParameterFromValue returns a pointer to a valid IpamServiceTemplatesListProtocolParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpamServiceTemplatesListProtocolParameterFromValue(v string) (*IpamServiceTemplatesListProtocolParameter, error) {
	ev := IpamServiceTemplatesListProtocolParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpamServiceTemplatesListProtocolParameter: valid values are %v", v, AllowedIpamServiceTemplatesListProtocolParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpamServiceTemplatesListProtocolParameter) IsValid() bool {
	for _, existing := range AllowedIpamServiceTemplatesListProtocolParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ipam_service_templates_list_protocol_parameter value
func (v IpamServiceTemplatesListProtocolParameter) Ptr() *IpamServiceTemplatesListProtocolParameter {
	return &v
}

type NullableIpamServiceTemplatesListProtocolParameter struct {
	value *IpamServiceTemplatesListProtocolParameter
	isSet bool
}

func (v NullableIpamServiceTemplatesListProtocolParameter) Get() *IpamServiceTemplatesListProtocolParameter {
	return v.value
}

func (v *NullableIpamServiceTemplatesListProtocolParameter) Set(val *IpamServiceTemplatesListProtocolParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamServiceTemplatesListProtocolParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamServiceTemplatesListProtocolParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamServiceTemplatesListProtocolParameter(val *IpamServiceTemplatesListProtocolParameter) *NullableIpamServiceTemplatesListProtocolParameter {
	return &NullableIpamServiceTemplatesListProtocolParameter{value: val, isSet: true}
}

func (v NullableIpamServiceTemplatesListProtocolParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamServiceTemplatesListProtocolParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
