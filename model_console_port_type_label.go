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

// ConsolePortTypeLabel the model 'ConsolePortTypeLabel'
type ConsolePortTypeLabel string

// List of ConsolePort_type_label
const (
	CONSOLEPORTTYPELABEL_DE_9         ConsolePortTypeLabel = "DE-9"
	CONSOLEPORTTYPELABEL_DB_25        ConsolePortTypeLabel = "DB-25"
	CONSOLEPORTTYPELABEL_RJ_11        ConsolePortTypeLabel = "RJ-11"
	CONSOLEPORTTYPELABEL_RJ_12        ConsolePortTypeLabel = "RJ-12"
	CONSOLEPORTTYPELABEL_RJ_45        ConsolePortTypeLabel = "RJ-45"
	CONSOLEPORTTYPELABEL_MINI_DIN_8   ConsolePortTypeLabel = "Mini-DIN 8"
	CONSOLEPORTTYPELABEL_USB_TYPE_A   ConsolePortTypeLabel = "USB Type A"
	CONSOLEPORTTYPELABEL_USB_TYPE_B   ConsolePortTypeLabel = "USB Type B"
	CONSOLEPORTTYPELABEL_USB_TYPE_C   ConsolePortTypeLabel = "USB Type C"
	CONSOLEPORTTYPELABEL_USB_MINI_A   ConsolePortTypeLabel = "USB Mini A"
	CONSOLEPORTTYPELABEL_USB_MINI_B   ConsolePortTypeLabel = "USB Mini B"
	CONSOLEPORTTYPELABEL_USB_MICRO_A  ConsolePortTypeLabel = "USB Micro A"
	CONSOLEPORTTYPELABEL_USB_MICRO_B  ConsolePortTypeLabel = "USB Micro B"
	CONSOLEPORTTYPELABEL_USB_MICRO_AB ConsolePortTypeLabel = "USB Micro AB"
	CONSOLEPORTTYPELABEL_OTHER        ConsolePortTypeLabel = "Other"
)

// All allowed values of ConsolePortTypeLabel enum
var AllowedConsolePortTypeLabelEnumValues = []ConsolePortTypeLabel{
	"DE-9",
	"DB-25",
	"RJ-11",
	"RJ-12",
	"RJ-45",
	"Mini-DIN 8",
	"USB Type A",
	"USB Type B",
	"USB Type C",
	"USB Mini A",
	"USB Mini B",
	"USB Micro A",
	"USB Micro B",
	"USB Micro AB",
	"Other",
}

func (v *ConsolePortTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsolePortTypeLabel(value)
	for _, existing := range AllowedConsolePortTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsolePortTypeLabel", value)
}

// NewConsolePortTypeLabelFromValue returns a pointer to a valid ConsolePortTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsolePortTypeLabelFromValue(v string) (*ConsolePortTypeLabel, error) {
	ev := ConsolePortTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsolePortTypeLabel: valid values are %v", v, AllowedConsolePortTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsolePortTypeLabel) IsValid() bool {
	for _, existing := range AllowedConsolePortTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsolePort_type_label value
func (v ConsolePortTypeLabel) Ptr() *ConsolePortTypeLabel {
	return &v
}

type NullableConsolePortTypeLabel struct {
	value *ConsolePortTypeLabel
	isSet bool
}

func (v NullableConsolePortTypeLabel) Get() *ConsolePortTypeLabel {
	return v.value
}

func (v *NullableConsolePortTypeLabel) Set(val *ConsolePortTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortTypeLabel(val *ConsolePortTypeLabel) *NullableConsolePortTypeLabel {
	return &NullableConsolePortTypeLabel{value: val, isSet: true}
}

func (v NullableConsolePortTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
