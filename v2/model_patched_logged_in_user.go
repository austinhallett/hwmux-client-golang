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

// PatchedLoggedInUser struct for PatchedLoggedInUser
type PatchedLoggedInUser struct {
	Id *int32 `json:"id,omitempty"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	// Designates whether the user can log into this admin site.
	IsStaff *bool `json:"is_staff,omitempty"`
	Groups []string `json:"groups,omitempty"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser *bool `json:"is_superuser,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewPatchedLoggedInUser instantiates a new PatchedLoggedInUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLoggedInUser() *PatchedLoggedInUser {
	this := PatchedLoggedInUser{}
	return &this
}

// NewPatchedLoggedInUserWithDefaults instantiates a new PatchedLoggedInUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLoggedInUserWithDefaults() *PatchedLoggedInUser {
	this := PatchedLoggedInUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PatchedLoggedInUser) SetId(v int32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedLoggedInUser) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedLoggedInUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
    return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PatchedLoggedInUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
    return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PatchedLoggedInUser) SetLastName(v string) {
	o.LastName = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetIsStaff() bool {
	if o == nil || isNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetIsStaffOk() (*bool, bool) {
	if o == nil || isNil(o.IsStaff) {
    return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasIsStaff() bool {
	if o != nil && !isNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *PatchedLoggedInUser) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *PatchedLoggedInUser) SetGroups(v []string) {
	o.Groups = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetIsSuperuser() bool {
	if o == nil || isNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || isNil(o.IsSuperuser) {
    return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasIsSuperuser() bool {
	if o != nil && !isNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *PatchedLoggedInUser) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PatchedLoggedInUser) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLoggedInUser) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedLoggedInUser) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PatchedLoggedInUser) SetPassword(v string) {
	o.Password = &v
}

func (o PatchedLoggedInUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !isNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !isNil(o.IsStaff) {
		toSerialize["is_staff"] = o.IsStaff
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedLoggedInUser struct {
	value *PatchedLoggedInUser
	isSet bool
}

func (v NullablePatchedLoggedInUser) Get() *PatchedLoggedInUser {
	return v.value
}

func (v *NullablePatchedLoggedInUser) Set(val *PatchedLoggedInUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLoggedInUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLoggedInUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLoggedInUser(val *PatchedLoggedInUser) *NullablePatchedLoggedInUser {
	return &NullablePatchedLoggedInUser{value: val, isSet: true}
}

func (v NullablePatchedLoggedInUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLoggedInUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


