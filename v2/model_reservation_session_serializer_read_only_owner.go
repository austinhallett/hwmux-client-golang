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

// ReservationSessionSerializerReadOnlyOwner struct for ReservationSessionSerializerReadOnlyOwner
type ReservationSessionSerializerReadOnlyOwner struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
}

// NewReservationSessionSerializerReadOnlyOwner instantiates a new ReservationSessionSerializerReadOnlyOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReservationSessionSerializerReadOnlyOwner(username string) *ReservationSessionSerializerReadOnlyOwner {
	this := ReservationSessionSerializerReadOnlyOwner{}
	this.Username = username
	return &this
}

// NewReservationSessionSerializerReadOnlyOwnerWithDefaults instantiates a new ReservationSessionSerializerReadOnlyOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReservationSessionSerializerReadOnlyOwnerWithDefaults() *ReservationSessionSerializerReadOnlyOwner {
	this := ReservationSessionSerializerReadOnlyOwner{}
	return &this
}

// GetUsername returns the Username field value
func (o *ReservationSessionSerializerReadOnlyOwner) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ReservationSessionSerializerReadOnlyOwner) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ReservationSessionSerializerReadOnlyOwner) SetUsername(v string) {
	o.Username = v
}

func (o ReservationSessionSerializerReadOnlyOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableReservationSessionSerializerReadOnlyOwner struct {
	value *ReservationSessionSerializerReadOnlyOwner
	isSet bool
}

func (v NullableReservationSessionSerializerReadOnlyOwner) Get() *ReservationSessionSerializerReadOnlyOwner {
	return v.value
}

func (v *NullableReservationSessionSerializerReadOnlyOwner) Set(val *ReservationSessionSerializerReadOnlyOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationSessionSerializerReadOnlyOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationSessionSerializerReadOnlyOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationSessionSerializerReadOnlyOwner(val *ReservationSessionSerializerReadOnlyOwner) *NullableReservationSessionSerializerReadOnlyOwner {
	return &NullableReservationSessionSerializerReadOnlyOwner{value: val, isSet: true}
}

func (v NullableReservationSessionSerializerReadOnlyOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationSessionSerializerReadOnlyOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


