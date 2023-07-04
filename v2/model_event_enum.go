/*
HW Mux Reservation System

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hwmux

import (
	"encoding/json"
	"fmt"
)

// EventEnum the model 'EventEnum'
type EventEnum string

// List of EventEnum
const (
	RES EventEnum = "RES"
	REL EventEnum = "REL"
	OFF EventEnum = "OFF"
	ON EventEnum = "ON"
	CR EventEnum = "CR"
	MOD EventEnum = "MOD"
	QUE EventEnum = "QUE"
	ERR EventEnum = "ERR"
)

// All allowed values of EventEnum enum
var AllowedEventEnumEnumValues = []EventEnum{
	"RES",
	"REL",
	"OFF",
	"ON",
	"CR",
	"MOD",
	"QUE",
	"ERR",
}

func (v *EventEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventEnum(value)
	for _, existing := range AllowedEventEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventEnum", value)
}

// NewEventEnumFromValue returns a pointer to a valid EventEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventEnumFromValue(v string) (*EventEnum, error) {
	ev := EventEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventEnum: valid values are %v", v, AllowedEventEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventEnum) IsValid() bool {
	for _, existing := range AllowedEventEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventEnum value
func (v EventEnum) Ptr() *EventEnum {
	return &v
}

type NullableEventEnum struct {
	value *EventEnum
	isSet bool
}

func (v NullableEventEnum) Get() *EventEnum {
	return v.value
}

func (v *NullableEventEnum) Set(val *EventEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEventEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEventEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventEnum(val *EventEnum) *NullableEventEnum {
	return &NullableEventEnum{value: val, isSet: true}
}

func (v NullableEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
