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
	"time"
)

// checks if the DocumentReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentReference{}

// DocumentReference struct for DocumentReference
type DocumentReference struct {
	// Identifies whether the document is active and used for checks.
	Active *bool `json:"active,omitempty"`
	// Your description for the document.
	Description *string `json:"description,omitempty"`
	// Document name.
	FileName *string `json:"fileName,omitempty"`
	// The unique identifier of the resource.
	Id *string `json:"id,omitempty"`
	// The modification date of the document.
	ModificationDate *time.Time `json:"modificationDate,omitempty"`
	// Type of document, used when providing an ID number or uploading a document.
	Type *string `json:"type,omitempty"`
}

// NewDocumentReference instantiates a new DocumentReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentReference() *DocumentReference {
	this := DocumentReference{}
	return &this
}

// NewDocumentReferenceWithDefaults instantiates a new DocumentReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentReferenceWithDefaults() *DocumentReference {
	this := DocumentReference{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DocumentReference) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DocumentReference) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DocumentReference) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DocumentReference) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DocumentReference) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DocumentReference) SetDescription(v string) {
	o.Description = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DocumentReference) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DocumentReference) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DocumentReference) SetFileName(v string) {
	o.FileName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentReference) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentReference) SetId(v string) {
	o.Id = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *DocumentReference) GetModificationDate() time.Time {
	if o == nil || IsNil(o.ModificationDate) {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationDate) {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *DocumentReference) HasModificationDate() bool {
	if o != nil && !IsNil(o.ModificationDate) {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *DocumentReference) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DocumentReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DocumentReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DocumentReference) SetType(v string) {
	o.Type = &v
}

func (o DocumentReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModificationDate) {
		toSerialize["modificationDate"] = o.ModificationDate
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDocumentReference struct {
	value *DocumentReference
	isSet bool
}

func (v NullableDocumentReference) Get() *DocumentReference {
	return v.value
}

func (v *NullableDocumentReference) Set(val *DocumentReference) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentReference) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentReference(val *DocumentReference) *NullableDocumentReference {
	return &NullableDocumentReference{value: val, isSet: true}
}

func (v NullableDocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

