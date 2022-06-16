/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// Difference struct for Difference
type Difference struct {
	Base     *Userset `json:"base,omitempty"`
	Subtract *Userset `json:"subtract,omitempty"`
}

// NewDifference instantiates a new Difference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDifference() *Difference {
	this := Difference{}
	return &this
}

// NewDifferenceWithDefaults instantiates a new Difference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDifferenceWithDefaults() *Difference {
	this := Difference{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *Difference) GetBase() Userset {
	if o == nil || o.Base == nil {
		var ret Userset
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Difference) GetBaseOk() (*Userset, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *Difference) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given Userset and assigns it to the Base field.
func (o *Difference) SetBase(v Userset) {
	o.Base = &v
}

// GetSubtract returns the Subtract field value if set, zero value otherwise.
func (o *Difference) GetSubtract() Userset {
	if o == nil || o.Subtract == nil {
		var ret Userset
		return ret
	}
	return *o.Subtract
}

// GetSubtractOk returns a tuple with the Subtract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Difference) GetSubtractOk() (*Userset, bool) {
	if o == nil || o.Subtract == nil {
		return nil, false
	}
	return o.Subtract, true
}

// HasSubtract returns a boolean if a field has been set.
func (o *Difference) HasSubtract() bool {
	if o != nil && o.Subtract != nil {
		return true
	}

	return false
}

// SetSubtract gets a reference to the given Userset and assigns it to the Subtract field.
func (o *Difference) SetSubtract(v Userset) {
	o.Subtract = &v
}

func (o Difference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Subtract != nil {
		toSerialize["subtract"] = o.Subtract
	}
	return json.Marshal(toSerialize)
}

type NullableDifference struct {
	value *Difference
	isSet bool
}

func (v NullableDifference) Get() *Difference {
	return v.value
}

func (v *NullableDifference) Set(val *Difference) {
	v.value = val
	v.isSet = true
}

func (v NullableDifference) IsSet() bool {
	return v.isSet
}

func (v *NullableDifference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDifference(val *Difference) *NullableDifference {
	return &NullableDifference{value: val, isSet: true}
}

func (v NullableDifference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDifference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}