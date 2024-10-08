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

// checks if the WritableFrontPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableFrontPortRequest{}

// WritableFrontPortRequest Adds support for custom fields and tags.
type WritableFrontPortRequest struct {
	Device DeviceRequest         `json:"device"`
	Module NullableModuleRequest `json:"module,omitempty"`
	Name   string                `json:"name"`
	// Physical label
	Label    *string            `json:"label,omitempty"`
	Type     FrontPortTypeValue `json:"type"`
	Color    *string            `json:"color,omitempty"`
	RearPort int32              `json:"rear_port"`
	// Mapped position on corresponding rear port
	RearPortPosition *int32  `json:"rear_port_position,omitempty"`
	Description      *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableFrontPortRequest WritableFrontPortRequest

// NewWritableFrontPortRequest instantiates a new WritableFrontPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableFrontPortRequest(device DeviceRequest, name string, type_ FrontPortTypeValue, rearPort int32) *WritableFrontPortRequest {
	this := WritableFrontPortRequest{}
	this.Device = device
	this.Name = name
	this.Type = type_
	this.RearPort = rearPort
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// NewWritableFrontPortRequestWithDefaults instantiates a new WritableFrontPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableFrontPortRequestWithDefaults() *WritableFrontPortRequest {
	this := WritableFrontPortRequest{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetDevice returns the Device field value
func (o *WritableFrontPortRequest) GetDevice() DeviceRequest {
	if o == nil {
		var ret DeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetDeviceOk() (*DeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableFrontPortRequest) SetDevice(v DeviceRequest) {
	o.Device = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableFrontPortRequest) GetModule() ModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret ModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableFrontPortRequest) GetModuleOk() (*ModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableModuleRequest and assigns it to the Module field.
func (o *WritableFrontPortRequest) SetModule(v ModuleRequest) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *WritableFrontPortRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *WritableFrontPortRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value
func (o *WritableFrontPortRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableFrontPortRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableFrontPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *WritableFrontPortRequest) GetType() FrontPortTypeValue {
	if o == nil {
		var ret FrontPortTypeValue
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetTypeOk() (*FrontPortTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WritableFrontPortRequest) SetType(v FrontPortTypeValue) {
	o.Type = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *WritableFrontPortRequest) SetColor(v string) {
	o.Color = &v
}

// GetRearPort returns the RearPort field value
func (o *WritableFrontPortRequest) GetRearPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetRearPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *WritableFrontPortRequest) SetRearPort(v int32) {
	o.RearPort = v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *WritableFrontPortRequest) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableFrontPortRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *WritableFrontPortRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableFrontPortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableFrontPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableFrontPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableFrontPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableFrontPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableFrontPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableFrontPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	toSerialize["rear_port"] = o.RearPort
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *WritableFrontPortRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
		"name",
		"type",
		"rear_port",
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

	varWritableFrontPortRequest := _WritableFrontPortRequest{}

	err = json.Unmarshal(data, &varWritableFrontPortRequest)

	if err != nil {
		return err
	}

	*o = WritableFrontPortRequest(varWritableFrontPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "rear_port")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableFrontPortRequest struct {
	value *WritableFrontPortRequest
	isSet bool
}

func (v NullableWritableFrontPortRequest) Get() *WritableFrontPortRequest {
	return v.value
}

func (v *NullableWritableFrontPortRequest) Set(val *WritableFrontPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableFrontPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableFrontPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableFrontPortRequest(val *WritableFrontPortRequest) *NullableWritableFrontPortRequest {
	return &NullableWritableFrontPortRequest{value: val, isSet: true}
}

func (v NullableWritableFrontPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableFrontPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
