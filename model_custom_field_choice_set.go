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
	"time"
)

// checks if the CustomFieldChoiceSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldChoiceSet{}

// CustomFieldChoiceSet Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomFieldChoiceSet struct {
	Id           int32                            `json:"id"`
	Url          string                           `json:"url"`
	Display      string                           `json:"display"`
	Name         string                           `json:"name"`
	Description  *string                          `json:"description,omitempty"`
	BaseChoices  *CustomFieldChoiceSetBaseChoices `json:"base_choices,omitempty"`
	ExtraChoices [][]interface{}                  `json:"extra_choices"`
	// Choices are automatically ordered alphabetically
	OrderAlphabetically  *bool        `json:"order_alphabetically,omitempty"`
	ChoicesCount         *string      `json:"choices_count,omitempty"`
	Created              NullableTime `json:"created"`
	LastUpdated          NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldChoiceSet CustomFieldChoiceSet

// NewCustomFieldChoiceSet instantiates a new CustomFieldChoiceSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldChoiceSet(id int32, url string, display string, name string, extraChoices [][]interface{}, created NullableTime, lastUpdated NullableTime) *CustomFieldChoiceSet {
	this := CustomFieldChoiceSet{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.ExtraChoices = extraChoices
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewCustomFieldChoiceSetWithDefaults instantiates a new CustomFieldChoiceSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldChoiceSetWithDefaults() *CustomFieldChoiceSet {
	this := CustomFieldChoiceSet{}
	return &this
}

// GetId returns the Id field value
func (o *CustomFieldChoiceSet) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomFieldChoiceSet) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CustomFieldChoiceSet) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CustomFieldChoiceSet) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *CustomFieldChoiceSet) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *CustomFieldChoiceSet) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *CustomFieldChoiceSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomFieldChoiceSet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomFieldChoiceSet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomFieldChoiceSet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomFieldChoiceSet) SetDescription(v string) {
	o.Description = &v
}

// GetBaseChoices returns the BaseChoices field value if set, zero value otherwise.
func (o *CustomFieldChoiceSet) GetBaseChoices() CustomFieldChoiceSetBaseChoices {
	if o == nil || IsNil(o.BaseChoices) {
		var ret CustomFieldChoiceSetBaseChoices
		return ret
	}
	return *o.BaseChoices
}

// GetBaseChoicesOk returns a tuple with the BaseChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetBaseChoicesOk() (*CustomFieldChoiceSetBaseChoices, bool) {
	if o == nil || IsNil(o.BaseChoices) {
		return nil, false
	}
	return o.BaseChoices, true
}

// HasBaseChoices returns a boolean if a field has been set.
func (o *CustomFieldChoiceSet) HasBaseChoices() bool {
	if o != nil && !IsNil(o.BaseChoices) {
		return true
	}

	return false
}

// SetBaseChoices gets a reference to the given CustomFieldChoiceSetBaseChoices and assigns it to the BaseChoices field.
func (o *CustomFieldChoiceSet) SetBaseChoices(v CustomFieldChoiceSetBaseChoices) {
	o.BaseChoices = &v
}

// GetExtraChoices returns the ExtraChoices field value
func (o *CustomFieldChoiceSet) GetExtraChoices() [][]interface{} {
	if o == nil {
		var ret [][]interface{}
		return ret
	}

	return o.ExtraChoices
}

// GetExtraChoicesOk returns a tuple with the ExtraChoices field value
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetExtraChoicesOk() ([][]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtraChoices, true
}

// SetExtraChoices sets field value
func (o *CustomFieldChoiceSet) SetExtraChoices(v [][]interface{}) {
	o.ExtraChoices = v
}

// GetOrderAlphabetically returns the OrderAlphabetically field value if set, zero value otherwise.
func (o *CustomFieldChoiceSet) GetOrderAlphabetically() bool {
	if o == nil || IsNil(o.OrderAlphabetically) {
		var ret bool
		return ret
	}
	return *o.OrderAlphabetically
}

// GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetOrderAlphabeticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.OrderAlphabetically) {
		return nil, false
	}
	return o.OrderAlphabetically, true
}

// HasOrderAlphabetically returns a boolean if a field has been set.
func (o *CustomFieldChoiceSet) HasOrderAlphabetically() bool {
	if o != nil && !IsNil(o.OrderAlphabetically) {
		return true
	}

	return false
}

// SetOrderAlphabetically gets a reference to the given bool and assigns it to the OrderAlphabetically field.
func (o *CustomFieldChoiceSet) SetOrderAlphabetically(v bool) {
	o.OrderAlphabetically = &v
}

// GetChoicesCount returns the ChoicesCount field value if set, zero value otherwise.
func (o *CustomFieldChoiceSet) GetChoicesCount() string {
	if o == nil || IsNil(o.ChoicesCount) {
		var ret string
		return ret
	}
	return *o.ChoicesCount
}

// GetChoicesCountOk returns a tuple with the ChoicesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldChoiceSet) GetChoicesCountOk() (*string, bool) {
	if o == nil || IsNil(o.ChoicesCount) {
		return nil, false
	}
	return o.ChoicesCount, true
}

// HasChoicesCount returns a boolean if a field has been set.
func (o *CustomFieldChoiceSet) HasChoicesCount() bool {
	if o != nil && !IsNil(o.ChoicesCount) {
		return true
	}

	return false
}

// SetChoicesCount gets a reference to the given string and assigns it to the ChoicesCount field.
func (o *CustomFieldChoiceSet) SetChoicesCount(v string) {
	o.ChoicesCount = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomFieldChoiceSet) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChoiceSet) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *CustomFieldChoiceSet) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomFieldChoiceSet) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChoiceSet) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CustomFieldChoiceSet) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o CustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldChoiceSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BaseChoices) {
		toSerialize["base_choices"] = o.BaseChoices
	}
	toSerialize["extra_choices"] = o.ExtraChoices
	if !IsNil(o.OrderAlphabetically) {
		toSerialize["order_alphabetically"] = o.OrderAlphabetically
	}
	if !IsNil(o.ChoicesCount) {
		toSerialize["choices_count"] = o.ChoicesCount
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomFieldChoiceSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"extra_choices",
		"created",
		"last_updated",
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

	varCustomFieldChoiceSet := _CustomFieldChoiceSet{}

	err = json.Unmarshal(data, &varCustomFieldChoiceSet)

	if err != nil {
		return err
	}

	*o = CustomFieldChoiceSet(varCustomFieldChoiceSet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "base_choices")
		delete(additionalProperties, "extra_choices")
		delete(additionalProperties, "order_alphabetically")
		delete(additionalProperties, "choices_count")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldChoiceSet struct {
	value *CustomFieldChoiceSet
	isSet bool
}

func (v NullableCustomFieldChoiceSet) Get() *CustomFieldChoiceSet {
	return v.value
}

func (v *NullableCustomFieldChoiceSet) Set(val *CustomFieldChoiceSet) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChoiceSet) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChoiceSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChoiceSet(val *CustomFieldChoiceSet) *NullableCustomFieldChoiceSet {
	return &NullableCustomFieldChoiceSet{value: val, isSet: true}
}

func (v NullableCustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChoiceSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
