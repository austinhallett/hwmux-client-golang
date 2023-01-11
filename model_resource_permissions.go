/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// ResourcePermissions struct for ResourcePermissions
type ResourcePermissions struct {
	Id int32 `json:"id"`
	Permissions []PermissionsEnum `json:"permissions,omitempty"`
}

// NewResourcePermissions instantiates a new ResourcePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePermissions(id int32) *ResourcePermissions {
	this := ResourcePermissions{}
	this.Id = id
	return &this
}

// NewResourcePermissionsWithDefaults instantiates a new ResourcePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePermissionsWithDefaults() *ResourcePermissions {
	this := ResourcePermissions{}
	return &this
}

// GetId returns the Id field value
func (o *ResourcePermissions) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourcePermissions) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourcePermissions) SetId(v int32) {
	o.Id = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ResourcePermissions) GetPermissions() []PermissionsEnum {
	if o == nil || isNil(o.Permissions) {
		var ret []PermissionsEnum
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePermissions) GetPermissionsOk() ([]PermissionsEnum, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ResourcePermissions) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []PermissionsEnum and assigns it to the Permissions field.
func (o *ResourcePermissions) SetPermissions(v []PermissionsEnum) {
	o.Permissions = v
}

func (o ResourcePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableResourcePermissions struct {
	value *ResourcePermissions
	isSet bool
}

func (v NullableResourcePermissions) Get() *ResourcePermissions {
	return v.value
}

func (v *NullableResourcePermissions) Set(val *ResourcePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePermissions(val *ResourcePermissions) *NullableResourcePermissions {
	return &NullableResourcePermissions{value: val, isSet: true}
}

func (v NullableResourcePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


