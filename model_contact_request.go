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

// checks if the ContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactRequest{}

// ContactRequest Adds support for custom fields and tags.
type ContactRequest struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactRequest ContactRequest

// NewContactRequest instantiates a new ContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequest(name string) *ContactRequest {
	this := ContactRequest{}
	this.Name = name
	return &this
}

// NewContactRequestWithDefaults instantiates a new ContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestWithDefaults() *ContactRequest {
	this := ContactRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ContactRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContactRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContactRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContactRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContactRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContactRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varContactRequest := _ContactRequest{}

	err = json.Unmarshal(data, &varContactRequest)

	if err != nil {
		return err
	}

	*o = ContactRequest(varContactRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactRequest struct {
	value *ContactRequest
	isSet bool
}

func (v NullableContactRequest) Get() *ContactRequest {
	return v.value
}

func (v *NullableContactRequest) Set(val *ContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequest(val *ContactRequest) *NullableContactRequest {
	return &NullableContactRequest{value: val, isSet: true}
}

func (v NullableContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
