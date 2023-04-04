/*
Legal Entity Management API

The Legal Entity Management API enables you to manage legal entities that contain information required for verification.  ## Authentication To connect to the Legal Entity Management API, you must use the basic authentication credentials of your web service user. If you don't have one, contact the [Adyen Support Team](https://www.adyen.help/hc/en-us/requests/new). Use the web service user credentials to authenticate your request, for example:  ``` curl -U \"ws12345@Scope.BalancePlatform_YourBalancePlatform\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Legal Entity Management API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://kyc-test.adyen.com/lem/v2/legalEntities ``` ## Going live When going live, your Adyen contact will provide your API credential for the live environment. You can then use the username and password to send requests to `https://kyc-live.adyen.com/lem/v2`.  

API version: 2
Contact: developer-experience@adyen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"
)

// checks if the CapabilityProblem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityProblem{}

// CapabilityProblem struct for CapabilityProblem
type CapabilityProblem struct {
	Entity *CapabilityProblemEntity `json:"entity,omitempty"`
	VerificationErrors []VerificationError `json:"verificationErrors,omitempty"`
}

// NewCapabilityProblem instantiates a new CapabilityProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityProblem() *CapabilityProblem {
	this := CapabilityProblem{}
	return &this
}

// NewCapabilityProblemWithDefaults instantiates a new CapabilityProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityProblemWithDefaults() *CapabilityProblem {
	this := CapabilityProblem{}
	return &this
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *CapabilityProblem) GetEntity() CapabilityProblemEntity {
	if o == nil || IsNil(o.Entity) {
		var ret CapabilityProblemEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblem) GetEntityOk() (*CapabilityProblemEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *CapabilityProblem) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given CapabilityProblemEntity and assigns it to the Entity field.
func (o *CapabilityProblem) SetEntity(v CapabilityProblemEntity) {
	o.Entity = &v
}

// GetVerificationErrors returns the VerificationErrors field value if set, zero value otherwise.
func (o *CapabilityProblem) GetVerificationErrors() []VerificationError {
	if o == nil || IsNil(o.VerificationErrors) {
		var ret []VerificationError
		return ret
	}
	return o.VerificationErrors
}

// GetVerificationErrorsOk returns a tuple with the VerificationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblem) GetVerificationErrorsOk() ([]VerificationError, bool) {
	if o == nil || IsNil(o.VerificationErrors) {
		return nil, false
	}
	return o.VerificationErrors, true
}

// HasVerificationErrors returns a boolean if a field has been set.
func (o *CapabilityProblem) HasVerificationErrors() bool {
	if o != nil && !IsNil(o.VerificationErrors) {
		return true
	}

	return false
}

// SetVerificationErrors gets a reference to the given []VerificationError and assigns it to the VerificationErrors field.
func (o *CapabilityProblem) SetVerificationErrors(v []VerificationError) {
	o.VerificationErrors = v
}

func (o CapabilityProblem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityProblem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.VerificationErrors) {
		toSerialize["verificationErrors"] = o.VerificationErrors
	}
	return toSerialize, nil
}

type NullableCapabilityProblem struct {
	value *CapabilityProblem
	isSet bool
}

func (v NullableCapabilityProblem) Get() *CapabilityProblem {
	return v.value
}

func (v *NullableCapabilityProblem) Set(val *CapabilityProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityProblem(val *CapabilityProblem) *NullableCapabilityProblem {
	return &NullableCapabilityProblem{value: val, isSet: true}
}

func (v NullableCapabilityProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

