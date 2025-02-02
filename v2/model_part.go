/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.24.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// Part Serializer for the Part model
type Part struct {
	PartNo string `json:"part_no"`
	PartFamily PartFamily `json:"part_family"`
	BoardNo NullableString `json:"board_no"`
	Variant *string `json:"variant,omitempty"`
	Revision *string `json:"revision,omitempty"`
	ChipNo NullableString `json:"chip_no,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPart instantiates a new Part object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPart(partNo string, partFamily PartFamily, boardNo NullableString) *Part {
	this := Part{}
	this.PartNo = partNo
	this.PartFamily = partFamily
	this.BoardNo = boardNo
	var variant string = ""
	this.Variant = &variant
	var revision string = ""
	this.Revision = &revision
	return &this
}

// NewPartWithDefaults instantiates a new Part object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartWithDefaults() *Part {
	this := Part{}
	var variant string = ""
	this.Variant = &variant
	var revision string = ""
	this.Revision = &revision
	return &this
}

// GetPartNo returns the PartNo field value
func (o *Part) GetPartNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartNo
}

// GetPartNoOk returns a tuple with the PartNo field value
// and a boolean to check if the value has been set.
func (o *Part) GetPartNoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PartNo, true
}

// SetPartNo sets field value
func (o *Part) SetPartNo(v string) {
	o.PartNo = v
}

// GetPartFamily returns the PartFamily field value
func (o *Part) GetPartFamily() PartFamily {
	if o == nil {
		var ret PartFamily
		return ret
	}

	return o.PartFamily
}

// GetPartFamilyOk returns a tuple with the PartFamily field value
// and a boolean to check if the value has been set.
func (o *Part) GetPartFamilyOk() (*PartFamily, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PartFamily, true
}

// SetPartFamily sets field value
func (o *Part) SetPartFamily(v PartFamily) {
	o.PartFamily = v
}

// GetBoardNo returns the BoardNo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Part) GetBoardNo() string {
	if o == nil || o.BoardNo.Get() == nil {
		var ret string
		return ret
	}

	return *o.BoardNo.Get()
}

// GetBoardNoOk returns a tuple with the BoardNo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Part) GetBoardNoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.BoardNo.Get(), o.BoardNo.IsSet()
}

// SetBoardNo sets field value
func (o *Part) SetBoardNo(v string) {
	o.BoardNo.Set(&v)
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Part) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
    return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Part) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *Part) SetVariant(v string) {
	o.Variant = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Part) GetRevision() string {
	if o == nil || isNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetRevisionOk() (*string, bool) {
	if o == nil || isNil(o.Revision) {
    return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *Part) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *Part) SetRevision(v string) {
	o.Revision = &v
}

// GetChipNo returns the ChipNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Part) GetChipNo() string {
	if o == nil || isNil(o.ChipNo.Get()) {
		var ret string
		return ret
	}
	return *o.ChipNo.Get()
}

// GetChipNoOk returns a tuple with the ChipNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Part) GetChipNoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChipNo.Get(), o.ChipNo.IsSet()
}

// HasChipNo returns a boolean if a field has been set.
func (o *Part) HasChipNo() bool {
	if o != nil && o.ChipNo.IsSet() {
		return true
	}

	return false
}

// SetChipNo gets a reference to the given NullableString and assigns it to the ChipNo field.
func (o *Part) SetChipNo(v string) {
	o.ChipNo.Set(&v)
}
// SetChipNoNil sets the value for ChipNo to be an explicit nil
func (o *Part) SetChipNoNil() {
	o.ChipNo.Set(nil)
}

// UnsetChipNo ensures that no value is present for ChipNo, not even an explicit nil
func (o *Part) UnsetChipNo() {
	o.ChipNo.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Part) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Part) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Part) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Part) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["part_no"] = o.PartNo
	}
	if true {
		toSerialize["part_family"] = o.PartFamily
	}
	if true {
		toSerialize["board_no"] = o.BoardNo.Get()
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if o.ChipNo.IsSet() {
		toSerialize["chip_no"] = o.ChipNo.Get()
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePart struct {
	value *Part
	isSet bool
}

func (v NullablePart) Get() *Part {
	return v.value
}

func (v *NullablePart) Set(val *Part) {
	v.value = val
	v.isSet = true
}

func (v NullablePart) IsSet() bool {
	return v.isSet
}

func (v *NullablePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePart(val *Part) *NullablePart {
	return &NullablePart{value: val, isSet: true}
}

func (v NullablePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


