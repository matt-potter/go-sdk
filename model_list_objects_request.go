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

// ListObjectsRequest struct for ListObjectsRequest
type ListObjectsRequest struct {
	AuthorizationModelId *string              `json:"authorization_model_id,omitempty"`
	Type                 *string              `json:"type,omitempty"`
	Relation             *string              `json:"relation,omitempty"`
	User                 *string              `json:"user,omitempty"`
	ContextualTuples     *ContextualTupleKeys `json:"contextual_tuples,omitempty"`
}

// NewListObjectsRequest instantiates a new ListObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectsRequest() *ListObjectsRequest {
	this := ListObjectsRequest{}
	return &this
}

// NewListObjectsRequestWithDefaults instantiates a new ListObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectsRequestWithDefaults() *ListObjectsRequest {
	this := ListObjectsRequest{}
	return &this
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *ListObjectsRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListObjectsRequest) SetType(v string) {
	o.Type = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *ListObjectsRequest) SetRelation(v string) {
	o.Relation = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ListObjectsRequest) SetUser(v string) {
	o.User = &v
}

// GetContextualTuples returns the ContextualTuples field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetContextualTuples() ContextualTupleKeys {
	if o == nil || o.ContextualTuples == nil {
		var ret ContextualTupleKeys
		return ret
	}
	return *o.ContextualTuples
}

// GetContextualTuplesOk returns a tuple with the ContextualTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool) {
	if o == nil || o.ContextualTuples == nil {
		return nil, false
	}
	return o.ContextualTuples, true
}

// HasContextualTuples returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasContextualTuples() bool {
	if o != nil && o.ContextualTuples != nil {
		return true
	}

	return false
}

// SetContextualTuples gets a reference to the given ContextualTupleKeys and assigns it to the ContextualTuples field.
func (o *ListObjectsRequest) SetContextualTuples(v ContextualTupleKeys) {
	o.ContextualTuples = &v
}

func (o ListObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ContextualTuples != nil {
		toSerialize["contextual_tuples"] = o.ContextualTuples
	}
	return json.Marshal(toSerialize)
}

type NullableListObjectsRequest struct {
	value *ListObjectsRequest
	isSet bool
}

func (v NullableListObjectsRequest) Get() *ListObjectsRequest {
	return v.value
}

func (v *NullableListObjectsRequest) Set(val *ListObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectsRequest(val *ListObjectsRequest) *NullableListObjectsRequest {
	return &NullableListObjectsRequest{value: val, isSet: true}
}

func (v NullableListObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
