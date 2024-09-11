/*
Tender Management API

API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка). 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateEmployeeRequest struct for CreateEmployeeRequest
type CreateEmployeeRequest struct {
	// Уникальный slug пользователя.
	Username string `json:"username"`
}

// NewCreateEmployeeRequest instantiates a new CreateEmployeeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmployeeRequest(username string) *CreateEmployeeRequest {
	this := CreateEmployeeRequest{}
	this.Username = username
	return &this
}

// NewCreateEmployeeRequestWithDefaults instantiates a new CreateEmployeeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmployeeRequestWithDefaults() *CreateEmployeeRequest {
	this := CreateEmployeeRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *CreateEmployeeRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateEmployeeRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateEmployeeRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateEmployeeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEmployeeRequest struct {
	value *CreateEmployeeRequest
	isSet bool
}

func (v NullableCreateEmployeeRequest) Get() *CreateEmployeeRequest {
	return v.value
}

func (v *NullableCreateEmployeeRequest) Set(val *CreateEmployeeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmployeeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmployeeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmployeeRequest(val *CreateEmployeeRequest) *NullableCreateEmployeeRequest {
	return &NullableCreateEmployeeRequest{value: val, isSet: true}
}

func (v NullableCreateEmployeeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmployeeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


