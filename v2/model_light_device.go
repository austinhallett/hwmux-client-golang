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

// LightDevice A lightweight device representation used in the DeviceGroup serializer
type LightDevice struct {
	Id int32 `json:"id"`
	SnOrName NullableString `json:"sn_or_name,omitempty"`
	Uri NullableString `json:"uri,omitempty"`
	IsWstk *bool `json:"is_wstk,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Online *bool `json:"online,omitempty"`
	Part Part `json:"part"`
	Location int32 `json:"location"`
	WstkPart NullableString `json:"wstk_part,omitempty"`
	Status *StatusEnum `json:"status,omitempty"`
	LocDesc string `json:"loc_desc"`
}

// NewLightDevice instantiates a new LightDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLightDevice(id int32, part Part, location int32, locDesc string) *LightDevice {
	this := LightDevice{}
	this.Id = id
	this.Part = part
	this.Location = location
	this.LocDesc = locDesc
	return &this
}

// NewLightDeviceWithDefaults instantiates a new LightDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLightDeviceWithDefaults() *LightDevice {
	this := LightDevice{}
	return &this
}

// GetId returns the Id field value
func (o *LightDevice) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LightDevice) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LightDevice) SetId(v int32) {
	o.Id = v
}

// GetSnOrName returns the SnOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LightDevice) GetSnOrName() string {
	if o == nil || isNil(o.SnOrName.Get()) {
		var ret string
		return ret
	}
	return *o.SnOrName.Get()
}

// GetSnOrNameOk returns a tuple with the SnOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LightDevice) GetSnOrNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SnOrName.Get(), o.SnOrName.IsSet()
}

// HasSnOrName returns a boolean if a field has been set.
func (o *LightDevice) HasSnOrName() bool {
	if o != nil && o.SnOrName.IsSet() {
		return true
	}

	return false
}

// SetSnOrName gets a reference to the given NullableString and assigns it to the SnOrName field.
func (o *LightDevice) SetSnOrName(v string) {
	o.SnOrName.Set(&v)
}
// SetSnOrNameNil sets the value for SnOrName to be an explicit nil
func (o *LightDevice) SetSnOrNameNil() {
	o.SnOrName.Set(nil)
}

// UnsetSnOrName ensures that no value is present for SnOrName, not even an explicit nil
func (o *LightDevice) UnsetSnOrName() {
	o.SnOrName.Unset()
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LightDevice) GetUri() string {
	if o == nil || isNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LightDevice) GetUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *LightDevice) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *LightDevice) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *LightDevice) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *LightDevice) UnsetUri() {
	o.Uri.Unset()
}

// GetIsWstk returns the IsWstk field value if set, zero value otherwise.
func (o *LightDevice) GetIsWstk() bool {
	if o == nil || isNil(o.IsWstk) {
		var ret bool
		return ret
	}
	return *o.IsWstk
}

// GetIsWstkOk returns a tuple with the IsWstk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LightDevice) GetIsWstkOk() (*bool, bool) {
	if o == nil || isNil(o.IsWstk) {
    return nil, false
	}
	return o.IsWstk, true
}

// HasIsWstk returns a boolean if a field has been set.
func (o *LightDevice) HasIsWstk() bool {
	if o != nil && !isNil(o.IsWstk) {
		return true
	}

	return false
}

// SetIsWstk gets a reference to the given bool and assigns it to the IsWstk field.
func (o *LightDevice) SetIsWstk(v bool) {
	o.IsWstk = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LightDevice) GetMetadata() map[string]interface{} {
	if o == nil || isNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LightDevice) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
    return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LightDevice) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LightDevice) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *LightDevice) GetOnline() bool {
	if o == nil || isNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LightDevice) GetOnlineOk() (*bool, bool) {
	if o == nil || isNil(o.Online) {
    return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *LightDevice) HasOnline() bool {
	if o != nil && !isNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *LightDevice) SetOnline(v bool) {
	o.Online = &v
}

// GetPart returns the Part field value
func (o *LightDevice) GetPart() Part {
	if o == nil {
		var ret Part
		return ret
	}

	return o.Part
}

// GetPartOk returns a tuple with the Part field value
// and a boolean to check if the value has been set.
func (o *LightDevice) GetPartOk() (*Part, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Part, true
}

// SetPart sets field value
func (o *LightDevice) SetPart(v Part) {
	o.Part = v
}

// GetLocation returns the Location field value
func (o *LightDevice) GetLocation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LightDevice) GetLocationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LightDevice) SetLocation(v int32) {
	o.Location = v
}

// GetWstkPart returns the WstkPart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LightDevice) GetWstkPart() string {
	if o == nil || isNil(o.WstkPart.Get()) {
		var ret string
		return ret
	}
	return *o.WstkPart.Get()
}

// GetWstkPartOk returns a tuple with the WstkPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LightDevice) GetWstkPartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.WstkPart.Get(), o.WstkPart.IsSet()
}

// HasWstkPart returns a boolean if a field has been set.
func (o *LightDevice) HasWstkPart() bool {
	if o != nil && o.WstkPart.IsSet() {
		return true
	}

	return false
}

// SetWstkPart gets a reference to the given NullableString and assigns it to the WstkPart field.
func (o *LightDevice) SetWstkPart(v string) {
	o.WstkPart.Set(&v)
}
// SetWstkPartNil sets the value for WstkPart to be an explicit nil
func (o *LightDevice) SetWstkPartNil() {
	o.WstkPart.Set(nil)
}

// UnsetWstkPart ensures that no value is present for WstkPart, not even an explicit nil
func (o *LightDevice) UnsetWstkPart() {
	o.WstkPart.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LightDevice) GetStatus() StatusEnum {
	if o == nil || isNil(o.Status) {
		var ret StatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LightDevice) GetStatusOk() (*StatusEnum, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LightDevice) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusEnum and assigns it to the Status field.
func (o *LightDevice) SetStatus(v StatusEnum) {
	o.Status = &v
}

// GetLocDesc returns the LocDesc field value
func (o *LightDevice) GetLocDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocDesc
}

// GetLocDescOk returns a tuple with the LocDesc field value
// and a boolean to check if the value has been set.
func (o *LightDevice) GetLocDescOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocDesc, true
}

// SetLocDesc sets field value
func (o *LightDevice) SetLocDesc(v string) {
	o.LocDesc = v
}

func (o LightDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.SnOrName.IsSet() {
		toSerialize["sn_or_name"] = o.SnOrName.Get()
	}
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if !isNil(o.IsWstk) {
		toSerialize["is_wstk"] = o.IsWstk
	}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if true {
		toSerialize["part"] = o.Part
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if o.WstkPart.IsSet() {
		toSerialize["wstk_part"] = o.WstkPart.Get()
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["loc_desc"] = o.LocDesc
	}
	return json.Marshal(toSerialize)
}

type NullableLightDevice struct {
	value *LightDevice
	isSet bool
}

func (v NullableLightDevice) Get() *LightDevice {
	return v.value
}

func (v *NullableLightDevice) Set(val *LightDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableLightDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableLightDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLightDevice(val *LightDevice) *NullableLightDevice {
	return &NullableLightDevice{value: val, isSet: true}
}

func (v NullableLightDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLightDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


