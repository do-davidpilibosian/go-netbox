/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.2 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the PatchedWritableConfigContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableConfigContextRequest{}

// PatchedWritableConfigContextRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableConfigContextRequest struct {
	Name          *string  `json:"name,omitempty"`
	Weight        *int32   `json:"weight,omitempty"`
	Description   *string  `json:"description,omitempty"`
	IsActive      *bool    `json:"is_active,omitempty"`
	Regions       []int32  `json:"regions,omitempty"`
	SiteGroups    []int32  `json:"site_groups,omitempty"`
	Sites         []int32  `json:"sites,omitempty"`
	Locations     []int32  `json:"locations,omitempty"`
	DeviceTypes   []int32  `json:"device_types,omitempty"`
	Roles         []int32  `json:"roles,omitempty"`
	Platforms     []int32  `json:"platforms,omitempty"`
	ClusterTypes  []int32  `json:"cluster_types,omitempty"`
	ClusterGroups []int32  `json:"cluster_groups,omitempty"`
	Clusters      []int32  `json:"clusters,omitempty"`
	TenantGroups  []int32  `json:"tenant_groups,omitempty"`
	Tenants       []int32  `json:"tenants,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	// Remote data source
	DataSource           NullableInt32          `json:"data_source,omitempty"`
	Data                 map[string]interface{} `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableConfigContextRequest PatchedWritableConfigContextRequest

// NewPatchedWritableConfigContextRequest instantiates a new PatchedWritableConfigContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableConfigContextRequest() *PatchedWritableConfigContextRequest {
	this := PatchedWritableConfigContextRequest{}
	return &this
}

// NewPatchedWritableConfigContextRequestWithDefaults instantiates a new PatchedWritableConfigContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableConfigContextRequestWithDefaults() *PatchedWritableConfigContextRequest {
	this := PatchedWritableConfigContextRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableConfigContextRequest) SetName(v string) {
	o.Name = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedWritableConfigContextRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableConfigContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedWritableConfigContextRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetRegions() []int32 {
	if o == nil || IsNil(o.Regions) {
		var ret []int32
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetRegionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []int32 and assigns it to the Regions field.
func (o *PatchedWritableConfigContextRequest) SetRegions(v []int32) {
	o.Regions = v
}

// GetSiteGroups returns the SiteGroups field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetSiteGroups() []int32 {
	if o == nil || IsNil(o.SiteGroups) {
		var ret []int32
		return ret
	}
	return o.SiteGroups
}

// GetSiteGroupsOk returns a tuple with the SiteGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetSiteGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SiteGroups) {
		return nil, false
	}
	return o.SiteGroups, true
}

// HasSiteGroups returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasSiteGroups() bool {
	if o != nil && !IsNil(o.SiteGroups) {
		return true
	}

	return false
}

// SetSiteGroups gets a reference to the given []int32 and assigns it to the SiteGroups field.
func (o *PatchedWritableConfigContextRequest) SetSiteGroups(v []int32) {
	o.SiteGroups = v
}

// GetSites returns the Sites field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetSites() []int32 {
	if o == nil || IsNil(o.Sites) {
		var ret []int32
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetSitesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Sites) {
		return nil, false
	}
	return o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasSites() bool {
	if o != nil && !IsNil(o.Sites) {
		return true
	}

	return false
}

// SetSites gets a reference to the given []int32 and assigns it to the Sites field.
func (o *PatchedWritableConfigContextRequest) SetSites(v []int32) {
	o.Sites = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetLocations() []int32 {
	if o == nil || IsNil(o.Locations) {
		var ret []int32
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetLocationsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []int32 and assigns it to the Locations field.
func (o *PatchedWritableConfigContextRequest) SetLocations(v []int32) {
	o.Locations = v
}

// GetDeviceTypes returns the DeviceTypes field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetDeviceTypes() []int32 {
	if o == nil || IsNil(o.DeviceTypes) {
		var ret []int32
		return ret
	}
	return o.DeviceTypes
}

// GetDeviceTypesOk returns a tuple with the DeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetDeviceTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.DeviceTypes) {
		return nil, false
	}
	return o.DeviceTypes, true
}

// HasDeviceTypes returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasDeviceTypes() bool {
	if o != nil && !IsNil(o.DeviceTypes) {
		return true
	}

	return false
}

// SetDeviceTypes gets a reference to the given []int32 and assigns it to the DeviceTypes field.
func (o *PatchedWritableConfigContextRequest) SetDeviceTypes(v []int32) {
	o.DeviceTypes = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetRoles() []int32 {
	if o == nil || IsNil(o.Roles) {
		var ret []int32
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetRolesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []int32 and assigns it to the Roles field.
func (o *PatchedWritableConfigContextRequest) SetRoles(v []int32) {
	o.Roles = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetPlatforms() []int32 {
	if o == nil || IsNil(o.Platforms) {
		var ret []int32
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetPlatformsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []int32 and assigns it to the Platforms field.
func (o *PatchedWritableConfigContextRequest) SetPlatforms(v []int32) {
	o.Platforms = v
}

// GetClusterTypes returns the ClusterTypes field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetClusterTypes() []int32 {
	if o == nil || IsNil(o.ClusterTypes) {
		var ret []int32
		return ret
	}
	return o.ClusterTypes
}

// GetClusterTypesOk returns a tuple with the ClusterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetClusterTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ClusterTypes) {
		return nil, false
	}
	return o.ClusterTypes, true
}

// HasClusterTypes returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasClusterTypes() bool {
	if o != nil && !IsNil(o.ClusterTypes) {
		return true
	}

	return false
}

// SetClusterTypes gets a reference to the given []int32 and assigns it to the ClusterTypes field.
func (o *PatchedWritableConfigContextRequest) SetClusterTypes(v []int32) {
	o.ClusterTypes = v
}

// GetClusterGroups returns the ClusterGroups field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetClusterGroups() []int32 {
	if o == nil || IsNil(o.ClusterGroups) {
		var ret []int32
		return ret
	}
	return o.ClusterGroups
}

// GetClusterGroupsOk returns a tuple with the ClusterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetClusterGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ClusterGroups) {
		return nil, false
	}
	return o.ClusterGroups, true
}

// HasClusterGroups returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasClusterGroups() bool {
	if o != nil && !IsNil(o.ClusterGroups) {
		return true
	}

	return false
}

// SetClusterGroups gets a reference to the given []int32 and assigns it to the ClusterGroups field.
func (o *PatchedWritableConfigContextRequest) SetClusterGroups(v []int32) {
	o.ClusterGroups = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetClusters() []int32 {
	if o == nil || IsNil(o.Clusters) {
		var ret []int32
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetClustersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []int32 and assigns it to the Clusters field.
func (o *PatchedWritableConfigContextRequest) SetClusters(v []int32) {
	o.Clusters = v
}

// GetTenantGroups returns the TenantGroups field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetTenantGroups() []int32 {
	if o == nil || IsNil(o.TenantGroups) {
		var ret []int32
		return ret
	}
	return o.TenantGroups
}

// GetTenantGroupsOk returns a tuple with the TenantGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetTenantGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TenantGroups) {
		return nil, false
	}
	return o.TenantGroups, true
}

// HasTenantGroups returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasTenantGroups() bool {
	if o != nil && !IsNil(o.TenantGroups) {
		return true
	}

	return false
}

// SetTenantGroups gets a reference to the given []int32 and assigns it to the TenantGroups field.
func (o *PatchedWritableConfigContextRequest) SetTenantGroups(v []int32) {
	o.TenantGroups = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetTenants() []int32 {
	if o == nil || IsNil(o.Tenants) {
		var ret []int32
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetTenantsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []int32 and assigns it to the Tenants field.
func (o *PatchedWritableConfigContextRequest) SetTenants(v []int32) {
	o.Tenants = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchedWritableConfigContextRequest) SetTags(v []string) {
	o.Tags = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableConfigContextRequest) GetDataSource() int32 {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableConfigContextRequest) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableInt32 and assigns it to the DataSource field.
func (o *PatchedWritableConfigContextRequest) SetDataSource(v int32) {
	o.DataSource.Set(&v)
}

// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *PatchedWritableConfigContextRequest) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *PatchedWritableConfigContextRequest) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PatchedWritableConfigContextRequest) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableConfigContextRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PatchedWritableConfigContextRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *PatchedWritableConfigContextRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o PatchedWritableConfigContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableConfigContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.SiteGroups) {
		toSerialize["site_groups"] = o.SiteGroups
	}
	if !IsNil(o.Sites) {
		toSerialize["sites"] = o.Sites
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.DeviceTypes) {
		toSerialize["device_types"] = o.DeviceTypes
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if !IsNil(o.ClusterTypes) {
		toSerialize["cluster_types"] = o.ClusterTypes
	}
	if !IsNil(o.ClusterGroups) {
		toSerialize["cluster_groups"] = o.ClusterGroups
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.TenantGroups) {
		toSerialize["tenant_groups"] = o.TenantGroups
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableConfigContextRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableConfigContextRequest := _PatchedWritableConfigContextRequest{}

	err = json.Unmarshal(bytes, &varPatchedWritableConfigContextRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableConfigContextRequest(varPatchedWritableConfigContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "description")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "regions")
		delete(additionalProperties, "site_groups")
		delete(additionalProperties, "sites")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "device_types")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "platforms")
		delete(additionalProperties, "cluster_types")
		delete(additionalProperties, "cluster_groups")
		delete(additionalProperties, "clusters")
		delete(additionalProperties, "tenant_groups")
		delete(additionalProperties, "tenants")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableConfigContextRequest struct {
	value *PatchedWritableConfigContextRequest
	isSet bool
}

func (v NullablePatchedWritableConfigContextRequest) Get() *PatchedWritableConfigContextRequest {
	return v.value
}

func (v *NullablePatchedWritableConfigContextRequest) Set(val *PatchedWritableConfigContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableConfigContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableConfigContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableConfigContextRequest(val *PatchedWritableConfigContextRequest) *NullablePatchedWritableConfigContextRequest {
	return &NullablePatchedWritableConfigContextRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableConfigContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableConfigContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
