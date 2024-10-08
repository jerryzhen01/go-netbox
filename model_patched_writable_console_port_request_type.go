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

// PatchedWritableConsolePortRequestType Physical port type  * `de-9` - DE-9 * `db-25` - DB-25 * `rj-11` - RJ-11 * `rj-12` - RJ-12 * `rj-45` - RJ-45 * `mini-din-8` - Mini-DIN 8 * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `other` - Other
type PatchedWritableConsolePortRequestType string

// List of PatchedWritableConsolePortRequest_type
const (
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_DE_9         PatchedWritableConsolePortRequestType = "de-9"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_DB_25        PatchedWritableConsolePortRequestType = "db-25"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_RJ_11        PatchedWritableConsolePortRequestType = "rj-11"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_RJ_12        PatchedWritableConsolePortRequestType = "rj-12"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_RJ_45        PatchedWritableConsolePortRequestType = "rj-45"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_MINI_DIN_8   PatchedWritableConsolePortRequestType = "mini-din-8"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_A        PatchedWritableConsolePortRequestType = "usb-a"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_B        PatchedWritableConsolePortRequestType = "usb-b"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_C        PatchedWritableConsolePortRequestType = "usb-c"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_MINI_A   PatchedWritableConsolePortRequestType = "usb-mini-a"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_MINI_B   PatchedWritableConsolePortRequestType = "usb-mini-b"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_MICRO_A  PatchedWritableConsolePortRequestType = "usb-micro-a"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_MICRO_B  PatchedWritableConsolePortRequestType = "usb-micro-b"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_USB_MICRO_AB PatchedWritableConsolePortRequestType = "usb-micro-ab"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_OTHER        PatchedWritableConsolePortRequestType = "other"
	PATCHEDWRITABLECONSOLEPORTREQUESTTYPE_EMPTY        PatchedWritableConsolePortRequestType = ""
)

// All allowed values of PatchedWritableConsolePortRequestType enum
var AllowedPatchedWritableConsolePortRequestTypeEnumValues = []PatchedWritableConsolePortRequestType{
	"de-9",
	"db-25",
	"rj-11",
	"rj-12",
	"rj-45",
	"mini-din-8",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"other",
	"",
}

func (v *PatchedWritableConsolePortRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableConsolePortRequestType(value)
	for _, existing := range AllowedPatchedWritableConsolePortRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableConsolePortRequestType", value)
}

// NewPatchedWritableConsolePortRequestTypeFromValue returns a pointer to a valid PatchedWritableConsolePortRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableConsolePortRequestTypeFromValue(v string) (*PatchedWritableConsolePortRequestType, error) {
	ev := PatchedWritableConsolePortRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableConsolePortRequestType: valid values are %v", v, AllowedPatchedWritableConsolePortRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableConsolePortRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritableConsolePortRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableConsolePortRequest_type value
func (v PatchedWritableConsolePortRequestType) Ptr() *PatchedWritableConsolePortRequestType {
	return &v
}

type NullablePatchedWritableConsolePortRequestType struct {
	value *PatchedWritableConsolePortRequestType
	isSet bool
}

func (v NullablePatchedWritableConsolePortRequestType) Get() *PatchedWritableConsolePortRequestType {
	return v.value
}

func (v *NullablePatchedWritableConsolePortRequestType) Set(val *PatchedWritableConsolePortRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConsolePortRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConsolePortRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConsolePortRequestType(val *PatchedWritableConsolePortRequestType) *NullablePatchedWritableConsolePortRequestType {
	return &NullablePatchedWritableConsolePortRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritableConsolePortRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConsolePortRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
