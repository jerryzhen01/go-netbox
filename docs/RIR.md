# RIR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AggregateCount** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewRIR

`func NewRIR(id int32, url string, display string, name string, slug string, ) *RIR`

NewRIR instantiates a new RIR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRIRWithDefaults

`func NewRIRWithDefaults() *RIR`

NewRIRWithDefaults instantiates a new RIR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RIR) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RIR) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RIR) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *RIR) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RIR) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RIR) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *RIR) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RIR) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RIR) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *RIR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RIR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RIR) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *RIR) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *RIR) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *RIR) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *RIR) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RIR) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RIR) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RIR) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAggregateCount

`func (o *RIR) GetAggregateCount() int64`

GetAggregateCount returns the AggregateCount field if non-nil, zero value otherwise.

### GetAggregateCountOk

`func (o *RIR) GetAggregateCountOk() (*int64, bool)`

GetAggregateCountOk returns a tuple with the AggregateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateCount

`func (o *RIR) SetAggregateCount(v int64)`

SetAggregateCount sets AggregateCount field to given value.

### HasAggregateCount

`func (o *RIR) HasAggregateCount() bool`

HasAggregateCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


