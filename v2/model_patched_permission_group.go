/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.26.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
)

// checks if the PatchedPermissionGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedPermissionGroup{}

// PatchedPermissionGroup struct for PatchedPermissionGroup
type PatchedPermissionGroup struct {
	Id *int32 `json:"id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewPatchedPermissionGroup instantiates a new PatchedPermissionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPermissionGroup() *PatchedPermissionGroup {
	this := PatchedPermissionGroup{}
	return &this
}

// NewPatchedPermissionGroupWithDefaults instantiates a new PatchedPermissionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPermissionGroupWithDefaults() *PatchedPermissionGroup {
	this := PatchedPermissionGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedPermissionGroup) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionGroup) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedPermissionGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchedPermissionGroup) SetId(v int32) {
	o.Id = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PatchedPermissionGroup) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionGroup) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PatchedPermissionGroup) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *PatchedPermissionGroup) SetPermissions(v []string) {
	o.Permissions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPermissionGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPermissionGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPermissionGroup) SetName(v string) {
	o.Name = &v
}

func (o PatchedPermissionGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedPermissionGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePatchedPermissionGroup struct {
	value *PatchedPermissionGroup
	isSet bool
}

func (v NullablePatchedPermissionGroup) Get() *PatchedPermissionGroup {
	return v.value
}

func (v *NullablePatchedPermissionGroup) Set(val *PatchedPermissionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPermissionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPermissionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPermissionGroup(val *PatchedPermissionGroup) *NullablePatchedPermissionGroup {
	return &NullablePatchedPermissionGroup{value: val, isSet: true}
}

func (v NullablePatchedPermissionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPermissionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


