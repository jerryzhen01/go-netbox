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

// checks if the VirtualMachineWithConfigContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineWithConfigContextRequest{}

// VirtualMachineWithConfigContextRequest Adds support for custom fields and tags.
type VirtualMachineWithConfigContextRequest struct {
	Name           string                              `json:"name"`
	Status         *PatchedWritableModuleRequestStatus `json:"status,omitempty"`
	Site           NullableSiteRequest                 `json:"site,omitempty"`
	Cluster        NullableClusterRequest              `json:"cluster,omitempty"`
	Device         NullableDeviceRequest               `json:"device,omitempty"`
	Role           NullableDeviceRoleRequest           `json:"role,omitempty"`
	Tenant         NullableTenantRequest               `json:"tenant,omitempty"`
	Platform       NullablePlatformRequest             `json:"platform,omitempty"`
	PrimaryIp4     NullableIPAddressRequest            `json:"primary_ip4,omitempty"`
	PrimaryIp6     NullableIPAddressRequest            `json:"primary_ip6,omitempty"`
	Vcpus          NullableFloat64                     `json:"vcpus,omitempty"`
	Memory         NullableInt32                       `json:"memory,omitempty"`
	Disk           NullableInt32                       `json:"disk,omitempty"`
	Description    *string                             `json:"description,omitempty"`
	Comments       *string                             `json:"comments,omitempty"`
	ConfigTemplate NullableConfigTemplateRequest       `json:"config_template,omitempty"`
	// Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData     interface{}            `json:"local_context_data,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualMachineWithConfigContextRequest VirtualMachineWithConfigContextRequest

// NewVirtualMachineWithConfigContextRequest instantiates a new VirtualMachineWithConfigContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineWithConfigContextRequest(name string) *VirtualMachineWithConfigContextRequest {
	this := VirtualMachineWithConfigContextRequest{}
	this.Name = name
	return &this
}

// NewVirtualMachineWithConfigContextRequestWithDefaults instantiates a new VirtualMachineWithConfigContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineWithConfigContextRequestWithDefaults() *VirtualMachineWithConfigContextRequest {
	this := VirtualMachineWithConfigContextRequest{}
	return &this
}

// GetName returns the Name field value
func (o *VirtualMachineWithConfigContextRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VirtualMachineWithConfigContextRequest) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContextRequest) GetStatus() PatchedWritableModuleRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret PatchedWritableModuleRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetStatusOk() (*PatchedWritableModuleRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PatchedWritableModuleRequestStatus and assigns it to the Status field.
func (o *VirtualMachineWithConfigContextRequest) SetStatus(v PatchedWritableModuleRequestStatus) {
	o.Status = &v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetSite() SiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret SiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetSiteOk() (*SiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableSiteRequest and assigns it to the Site field.
func (o *VirtualMachineWithConfigContextRequest) SetSite(v SiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetSite() {
	o.Site.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetCluster() ClusterRequest {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret ClusterRequest
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetClusterOk() (*ClusterRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableClusterRequest and assigns it to the Cluster field.
func (o *VirtualMachineWithConfigContextRequest) SetCluster(v ClusterRequest) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetCluster() {
	o.Cluster.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetDevice() DeviceRequest {
	if o == nil || IsNil(o.Device.Get()) {
		var ret DeviceRequest
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetDeviceOk() (*DeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableDeviceRequest and assigns it to the Device field.
func (o *VirtualMachineWithConfigContextRequest) SetDevice(v DeviceRequest) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetRole() DeviceRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret DeviceRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetRoleOk() (*DeviceRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableDeviceRoleRequest and assigns it to the Role field.
func (o *VirtualMachineWithConfigContextRequest) SetRole(v DeviceRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetTenant() TenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret TenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetTenantOk() (*TenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableTenantRequest and assigns it to the Tenant field.
func (o *VirtualMachineWithConfigContextRequest) SetTenant(v TenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetPlatform() PlatformRequest {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret PlatformRequest
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetPlatformOk() (*PlatformRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullablePlatformRequest and assigns it to the Platform field.
func (o *VirtualMachineWithConfigContextRequest) SetPlatform(v PlatformRequest) {
	o.Platform.Set(&v)
}

// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetPrimaryIp4() IPAddressRequest {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret IPAddressRequest
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetPrimaryIp4Ok() (*IPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableIPAddressRequest and assigns it to the PrimaryIp4 field.
func (o *VirtualMachineWithConfigContextRequest) SetPrimaryIp4(v IPAddressRequest) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetPrimaryIp6() IPAddressRequest {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret IPAddressRequest
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetPrimaryIp6Ok() (*IPAddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableIPAddressRequest and assigns it to the PrimaryIp6 field.
func (o *VirtualMachineWithConfigContextRequest) SetPrimaryIp6(v IPAddressRequest) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetVcpus() float64 {
	if o == nil || IsNil(o.Vcpus.Get()) {
		var ret float64
		return ret
	}
	return *o.Vcpus.Get()
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetVcpusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vcpus.Get(), o.Vcpus.IsSet()
}

// HasVcpus returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasVcpus() bool {
	if o != nil && o.Vcpus.IsSet() {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given NullableFloat64 and assigns it to the Vcpus field.
func (o *VirtualMachineWithConfigContextRequest) SetVcpus(v float64) {
	o.Vcpus.Set(&v)
}

// SetVcpusNil sets the value for Vcpus to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetVcpusNil() {
	o.Vcpus.Set(nil)
}

// UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetVcpus() {
	o.Vcpus.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *VirtualMachineWithConfigContextRequest) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetMemory() {
	o.Memory.Unset()
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetDisk() int32 {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret int32
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableInt32 and assigns it to the Disk field.
func (o *VirtualMachineWithConfigContextRequest) SetDisk(v int32) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetDisk() {
	o.Disk.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualMachineWithConfigContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContextRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VirtualMachineWithConfigContextRequest) SetComments(v string) {
	o.Comments = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetConfigTemplate() ConfigTemplateRequest {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret ConfigTemplateRequest
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetConfigTemplateOk() (*ConfigTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableConfigTemplateRequest and assigns it to the ConfigTemplate field.
func (o *VirtualMachineWithConfigContextRequest) SetConfigTemplate(v ConfigTemplateRequest) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *VirtualMachineWithConfigContextRequest) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *VirtualMachineWithConfigContextRequest) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualMachineWithConfigContextRequest) GetLocalContextData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LocalContextData
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualMachineWithConfigContextRequest) GetLocalContextDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LocalContextData) {
		return nil, false
	}
	return &o.LocalContextData, true
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasLocalContextData() bool {
	if o != nil && !IsNil(o.LocalContextData) {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given interface{} and assigns it to the LocalContextData field.
func (o *VirtualMachineWithConfigContextRequest) SetLocalContextData(v interface{}) {
	o.LocalContextData = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContextRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *VirtualMachineWithConfigContextRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VirtualMachineWithConfigContextRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualMachineWithConfigContextRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VirtualMachineWithConfigContextRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VirtualMachineWithConfigContextRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o VirtualMachineWithConfigContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineWithConfigContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Vcpus.IsSet() {
		toSerialize["vcpus"] = o.Vcpus.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if o.LocalContextData != nil {
		toSerialize["local_context_data"] = o.LocalContextData
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

func (o *VirtualMachineWithConfigContextRequest) UnmarshalJSON(data []byte) (err error) {
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

	varVirtualMachineWithConfigContextRequest := _VirtualMachineWithConfigContextRequest{}

	err = json.Unmarshal(data, &varVirtualMachineWithConfigContextRequest)

	if err != nil {
		return err
	}

	*o = VirtualMachineWithConfigContextRequest(varVirtualMachineWithConfigContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "site")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "device")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "vcpus")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "disk")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "local_context_data")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualMachineWithConfigContextRequest struct {
	value *VirtualMachineWithConfigContextRequest
	isSet bool
}

func (v NullableVirtualMachineWithConfigContextRequest) Get() *VirtualMachineWithConfigContextRequest {
	return v.value
}

func (v *NullableVirtualMachineWithConfigContextRequest) Set(val *VirtualMachineWithConfigContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineWithConfigContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineWithConfigContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineWithConfigContextRequest(val *VirtualMachineWithConfigContextRequest) *NullableVirtualMachineWithConfigContextRequest {
	return &NullableVirtualMachineWithConfigContextRequest{value: val, isSet: true}
}

func (v NullableVirtualMachineWithConfigContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineWithConfigContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
