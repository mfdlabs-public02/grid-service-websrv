/*
Avatar Api v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avatarv1

import (
	"encoding/json"
)

// checks if the RobloxApiAvatarModelsAssetIdAndTypeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RobloxApiAvatarModelsAssetIdAndTypeModel{}

// RobloxApiAvatarModelsAssetIdAndTypeModel struct for RobloxApiAvatarModelsAssetIdAndTypeModel
type RobloxApiAvatarModelsAssetIdAndTypeModel struct {
	AssetId *int64 `json:"assetId,omitempty"`
	AssetTypeId *int64 `json:"assetTypeId,omitempty"`
}

// NewRobloxApiAvatarModelsAssetIdAndTypeModel instantiates a new RobloxApiAvatarModelsAssetIdAndTypeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRobloxApiAvatarModelsAssetIdAndTypeModel() *RobloxApiAvatarModelsAssetIdAndTypeModel {
	this := RobloxApiAvatarModelsAssetIdAndTypeModel{}
	return &this
}

// NewRobloxApiAvatarModelsAssetIdAndTypeModelWithDefaults instantiates a new RobloxApiAvatarModelsAssetIdAndTypeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRobloxApiAvatarModelsAssetIdAndTypeModelWithDefaults() *RobloxApiAvatarModelsAssetIdAndTypeModel {
	this := RobloxApiAvatarModelsAssetIdAndTypeModel{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) GetAssetId() int64 {
	if o == nil || IsNil(o.AssetId) {
		var ret int64
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) GetAssetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given int64 and assigns it to the AssetId field.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) SetAssetId(v int64) {
	o.AssetId = &v
}

// GetAssetTypeId returns the AssetTypeId field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) GetAssetTypeId() int64 {
	if o == nil || IsNil(o.AssetTypeId) {
		var ret int64
		return ret
	}
	return *o.AssetTypeId
}

// GetAssetTypeIdOk returns a tuple with the AssetTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) GetAssetTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssetTypeId) {
		return nil, false
	}
	return o.AssetTypeId, true
}

// HasAssetTypeId returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) HasAssetTypeId() bool {
	if o != nil && !IsNil(o.AssetTypeId) {
		return true
	}

	return false
}

// SetAssetTypeId gets a reference to the given int64 and assigns it to the AssetTypeId field.
func (o *RobloxApiAvatarModelsAssetIdAndTypeModel) SetAssetTypeId(v int64) {
	o.AssetTypeId = &v
}

func (o RobloxApiAvatarModelsAssetIdAndTypeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RobloxApiAvatarModelsAssetIdAndTypeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	if !IsNil(o.AssetTypeId) {
		toSerialize["assetTypeId"] = o.AssetTypeId
	}
	return toSerialize, nil
}

type NullableRobloxApiAvatarModelsAssetIdAndTypeModel struct {
	value *RobloxApiAvatarModelsAssetIdAndTypeModel
	isSet bool
}

func (v NullableRobloxApiAvatarModelsAssetIdAndTypeModel) Get() *RobloxApiAvatarModelsAssetIdAndTypeModel {
	return v.value
}

func (v *NullableRobloxApiAvatarModelsAssetIdAndTypeModel) Set(val *RobloxApiAvatarModelsAssetIdAndTypeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRobloxApiAvatarModelsAssetIdAndTypeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRobloxApiAvatarModelsAssetIdAndTypeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRobloxApiAvatarModelsAssetIdAndTypeModel(val *RobloxApiAvatarModelsAssetIdAndTypeModel) *NullableRobloxApiAvatarModelsAssetIdAndTypeModel {
	return &NullableRobloxApiAvatarModelsAssetIdAndTypeModel{value: val, isSet: true}
}

func (v NullableRobloxApiAvatarModelsAssetIdAndTypeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRobloxApiAvatarModelsAssetIdAndTypeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

