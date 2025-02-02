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

// LoggedInUser struct for LoggedInUser
type LoggedInUser struct {
	Id int32 `json:"id"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	// Designates whether the user can log into this admin site.
	IsStaff *bool `json:"is_staff,omitempty"`
	Groups []string `json:"groups"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser *bool `json:"is_superuser,omitempty"`
	Password string `json:"password"`
}

// NewLoggedInUser instantiates a new LoggedInUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggedInUser(id int32, username string, groups []string, password string) *LoggedInUser {
	this := LoggedInUser{}
	this.Id = id
	this.Username = username
	this.Groups = groups
	this.Password = password
	return &this
}

// NewLoggedInUserWithDefaults instantiates a new LoggedInUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggedInUserWithDefaults() *LoggedInUser {
	this := LoggedInUser{}
	return &this
}

// GetId returns the Id field value
func (o *LoggedInUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoggedInUser) SetId(v int32) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *LoggedInUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *LoggedInUser) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LoggedInUser) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LoggedInUser) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LoggedInUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *LoggedInUser) GetFirstName() string {
	if o == nil || isNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetFirstNameOk() (*string, bool) {
	if o == nil || isNil(o.FirstName) {
    return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *LoggedInUser) HasFirstName() bool {
	if o != nil && !isNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *LoggedInUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *LoggedInUser) GetLastName() string {
	if o == nil || isNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetLastNameOk() (*string, bool) {
	if o == nil || isNil(o.LastName) {
    return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *LoggedInUser) HasLastName() bool {
	if o != nil && !isNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *LoggedInUser) SetLastName(v string) {
	o.LastName = &v
}

// GetIsStaff returns the IsStaff field value if set, zero value otherwise.
func (o *LoggedInUser) GetIsStaff() bool {
	if o == nil || isNil(o.IsStaff) {
		var ret bool
		return ret
	}
	return *o.IsStaff
}

// GetIsStaffOk returns a tuple with the IsStaff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetIsStaffOk() (*bool, bool) {
	if o == nil || isNil(o.IsStaff) {
    return nil, false
	}
	return o.IsStaff, true
}

// HasIsStaff returns a boolean if a field has been set.
func (o *LoggedInUser) HasIsStaff() bool {
	if o != nil && !isNil(o.IsStaff) {
		return true
	}

	return false
}

// SetIsStaff gets a reference to the given bool and assigns it to the IsStaff field.
func (o *LoggedInUser) SetIsStaff(v bool) {
	o.IsStaff = &v
}

// GetGroups returns the Groups field value
func (o *LoggedInUser) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetGroupsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *LoggedInUser) SetGroups(v []string) {
	o.Groups = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *LoggedInUser) GetIsSuperuser() bool {
	if o == nil || isNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || isNil(o.IsSuperuser) {
    return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *LoggedInUser) HasIsSuperuser() bool {
	if o != nil && !isNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *LoggedInUser) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetPassword returns the Password field value
func (o *LoggedInUser) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoggedInUser) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoggedInUser) SetPassword(v string) {
	o.Password = v
}

func (o LoggedInUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
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
	if true {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableLoggedInUser struct {
	value *LoggedInUser
	isSet bool
}

func (v NullableLoggedInUser) Get() *LoggedInUser {
	return v.value
}

func (v *NullableLoggedInUser) Set(val *LoggedInUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggedInUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggedInUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggedInUser(val *LoggedInUser) *NullableLoggedInUser {
	return &NullableLoggedInUser{value: val, isSet: true}
}

func (v NullableLoggedInUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggedInUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


