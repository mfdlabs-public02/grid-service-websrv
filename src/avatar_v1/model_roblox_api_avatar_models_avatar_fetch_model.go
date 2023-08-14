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

// checks if the RobloxApiAvatarModelsAvatarFetchModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RobloxApiAvatarModelsAvatarFetchModel{}

// RobloxApiAvatarModelsAvatarFetchModel struct for RobloxApiAvatarModelsAvatarFetchModel
type RobloxApiAvatarModelsAvatarFetchModel struct {
	ResolvedAvatarType *string `json:"resolvedAvatarType,omitempty"`
	EquippedGearVersionIds []int64 `json:"equippedGearVersionIds,omitempty"`
	BackpackGearVersionIds []int64 `json:"backpackGearVersionIds,omitempty"`
	AssetAndAssetTypeIds []RobloxApiAvatarModelsAssetIdAndTypeModel `json:"assetAndAssetTypeIds,omitempty"`
	AnimationAssetIds *map[string]int64 `json:"animationAssetIds,omitempty"`
	BodyColors *RobloxApiAvatarModelsBodyColorsModel `json:"bodyColors,omitempty"`
	Scales *RobloxWebResponsesAvatarScaleModel `json:"scales,omitempty"`
}

// NewRobloxApiAvatarModelsAvatarFetchModel instantiates a new RobloxApiAvatarModelsAvatarFetchModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRobloxApiAvatarModelsAvatarFetchModel() *RobloxApiAvatarModelsAvatarFetchModel {
	this := RobloxApiAvatarModelsAvatarFetchModel{}
	return &this
}

// NewRobloxApiAvatarModelsAvatarFetchModelWithDefaults instantiates a new RobloxApiAvatarModelsAvatarFetchModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRobloxApiAvatarModelsAvatarFetchModelWithDefaults() *RobloxApiAvatarModelsAvatarFetchModel {
	this := RobloxApiAvatarModelsAvatarFetchModel{}
	return &this
}

// GetResolvedAvatarType returns the ResolvedAvatarType field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetResolvedAvatarType() string {
	if o == nil || IsNil(o.ResolvedAvatarType) {
		var ret string
		return ret
	}
	return *o.ResolvedAvatarType
}

// GetResolvedAvatarTypeOk returns a tuple with the ResolvedAvatarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetResolvedAvatarTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResolvedAvatarType) {
		return nil, false
	}
	return o.ResolvedAvatarType, true
}

// HasResolvedAvatarType returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasResolvedAvatarType() bool {
	if o != nil && !IsNil(o.ResolvedAvatarType) {
		return true
	}

	return false
}

// SetResolvedAvatarType gets a reference to the given string and assigns it to the ResolvedAvatarType field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetResolvedAvatarType(v string) {
	o.ResolvedAvatarType = &v
}

// GetEquippedGearVersionIds returns the EquippedGearVersionIds field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetEquippedGearVersionIds() []int64 {
	if o == nil || IsNil(o.EquippedGearVersionIds) {
		var ret []int64
		return ret
	}
	return o.EquippedGearVersionIds
}

// GetEquippedGearVersionIdsOk returns a tuple with the EquippedGearVersionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetEquippedGearVersionIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.EquippedGearVersionIds) {
		return nil, false
	}
	return o.EquippedGearVersionIds, true
}

// HasEquippedGearVersionIds returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasEquippedGearVersionIds() bool {
	if o != nil && !IsNil(o.EquippedGearVersionIds) {
		return true
	}

	return false
}

// SetEquippedGearVersionIds gets a reference to the given []int64 and assigns it to the EquippedGearVersionIds field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetEquippedGearVersionIds(v []int64) {
	o.EquippedGearVersionIds = v
}

// GetBackpackGearVersionIds returns the BackpackGearVersionIds field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetBackpackGearVersionIds() []int64 {
	if o == nil || IsNil(o.BackpackGearVersionIds) {
		var ret []int64
		return ret
	}
	return o.BackpackGearVersionIds
}

// GetBackpackGearVersionIdsOk returns a tuple with the BackpackGearVersionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetBackpackGearVersionIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.BackpackGearVersionIds) {
		return nil, false
	}
	return o.BackpackGearVersionIds, true
}

// HasBackpackGearVersionIds returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasBackpackGearVersionIds() bool {
	if o != nil && !IsNil(o.BackpackGearVersionIds) {
		return true
	}

	return false
}

// SetBackpackGearVersionIds gets a reference to the given []int64 and assigns it to the BackpackGearVersionIds field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetBackpackGearVersionIds(v []int64) {
	o.BackpackGearVersionIds = v
}

// GetAssetAndAssetTypeIds returns the AssetAndAssetTypeIds field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetAssetAndAssetTypeIds() []RobloxApiAvatarModelsAssetIdAndTypeModel {
	if o == nil || IsNil(o.AssetAndAssetTypeIds) {
		var ret []RobloxApiAvatarModelsAssetIdAndTypeModel
		return ret
	}
	return o.AssetAndAssetTypeIds
}

// GetAssetAndAssetTypeIdsOk returns a tuple with the AssetAndAssetTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetAssetAndAssetTypeIdsOk() ([]RobloxApiAvatarModelsAssetIdAndTypeModel, bool) {
	if o == nil || IsNil(o.AssetAndAssetTypeIds) {
		return nil, false
	}
	return o.AssetAndAssetTypeIds, true
}

// HasAssetAndAssetTypeIds returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasAssetAndAssetTypeIds() bool {
	if o != nil && !IsNil(o.AssetAndAssetTypeIds) {
		return true
	}

	return false
}

// SetAssetAndAssetTypeIds gets a reference to the given []RobloxApiAvatarModelsAssetIdAndTypeModel and assigns it to the AssetAndAssetTypeIds field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetAssetAndAssetTypeIds(v []RobloxApiAvatarModelsAssetIdAndTypeModel) {
	o.AssetAndAssetTypeIds = v
}

// GetAnimationAssetIds returns the AnimationAssetIds field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetAnimationAssetIds() map[string]int64 {
	if o == nil || IsNil(o.AnimationAssetIds) {
		var ret map[string]int64
		return ret
	}
	return *o.AnimationAssetIds
}

// GetAnimationAssetIdsOk returns a tuple with the AnimationAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetAnimationAssetIdsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.AnimationAssetIds) {
		return nil, false
	}
	return o.AnimationAssetIds, true
}

// HasAnimationAssetIds returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasAnimationAssetIds() bool {
	if o != nil && !IsNil(o.AnimationAssetIds) {
		return true
	}

	return false
}

// SetAnimationAssetIds gets a reference to the given map[string]int64 and assigns it to the AnimationAssetIds field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetAnimationAssetIds(v map[string]int64) {
	o.AnimationAssetIds = &v
}

// GetBodyColors returns the BodyColors field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetBodyColors() RobloxApiAvatarModelsBodyColorsModel {
	if o == nil || IsNil(o.BodyColors) {
		var ret RobloxApiAvatarModelsBodyColorsModel
		return ret
	}
	return *o.BodyColors
}

// GetBodyColorsOk returns a tuple with the BodyColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetBodyColorsOk() (*RobloxApiAvatarModelsBodyColorsModel, bool) {
	if o == nil || IsNil(o.BodyColors) {
		return nil, false
	}
	return o.BodyColors, true
}

// HasBodyColors returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasBodyColors() bool {
	if o != nil && !IsNil(o.BodyColors) {
		return true
	}

	return false
}

// SetBodyColors gets a reference to the given RobloxApiAvatarModelsBodyColorsModel and assigns it to the BodyColors field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetBodyColors(v RobloxApiAvatarModelsBodyColorsModel) {
	o.BodyColors = &v
}

// GetScales returns the Scales field value if set, zero value otherwise.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetScales() RobloxWebResponsesAvatarScaleModel {
	if o == nil || IsNil(o.Scales) {
		var ret RobloxWebResponsesAvatarScaleModel
		return ret
	}
	return *o.Scales
}

// GetScalesOk returns a tuple with the Scales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) GetScalesOk() (*RobloxWebResponsesAvatarScaleModel, bool) {
	if o == nil || IsNil(o.Scales) {
		return nil, false
	}
	return o.Scales, true
}

// HasScales returns a boolean if a field has been set.
func (o *RobloxApiAvatarModelsAvatarFetchModel) HasScales() bool {
	if o != nil && !IsNil(o.Scales) {
		return true
	}

	return false
}

// SetScales gets a reference to the given RobloxWebResponsesAvatarScaleModel and assigns it to the Scales field.
func (o *RobloxApiAvatarModelsAvatarFetchModel) SetScales(v RobloxWebResponsesAvatarScaleModel) {
	o.Scales = &v
}

func (o RobloxApiAvatarModelsAvatarFetchModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RobloxApiAvatarModelsAvatarFetchModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResolvedAvatarType) {
		toSerialize["resolvedAvatarType"] = o.ResolvedAvatarType
	}
	if !IsNil(o.EquippedGearVersionIds) {
		toSerialize["equippedGearVersionIds"] = o.EquippedGearVersionIds
	}
	if !IsNil(o.BackpackGearVersionIds) {
		toSerialize["backpackGearVersionIds"] = o.BackpackGearVersionIds
	}
	if !IsNil(o.AssetAndAssetTypeIds) {
		toSerialize["assetAndAssetTypeIds"] = o.AssetAndAssetTypeIds
	}
	if !IsNil(o.AnimationAssetIds) {
		toSerialize["animationAssetIds"] = o.AnimationAssetIds
	}
	if !IsNil(o.BodyColors) {
		toSerialize["bodyColors"] = o.BodyColors
	}
	if !IsNil(o.Scales) {
		toSerialize["scales"] = o.Scales
	}
	return toSerialize, nil
}

type NullableRobloxApiAvatarModelsAvatarFetchModel struct {
	value *RobloxApiAvatarModelsAvatarFetchModel
	isSet bool
}

func (v NullableRobloxApiAvatarModelsAvatarFetchModel) Get() *RobloxApiAvatarModelsAvatarFetchModel {
	return v.value
}

func (v *NullableRobloxApiAvatarModelsAvatarFetchModel) Set(val *RobloxApiAvatarModelsAvatarFetchModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRobloxApiAvatarModelsAvatarFetchModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRobloxApiAvatarModelsAvatarFetchModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRobloxApiAvatarModelsAvatarFetchModel(val *RobloxApiAvatarModelsAvatarFetchModel) *NullableRobloxApiAvatarModelsAvatarFetchModel {
	return &NullableRobloxApiAvatarModelsAvatarFetchModel{value: val, isSet: true}
}

func (v NullableRobloxApiAvatarModelsAvatarFetchModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRobloxApiAvatarModelsAvatarFetchModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

