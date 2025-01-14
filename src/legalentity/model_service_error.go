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

// checks if the ServiceError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceError{}

// ServiceError struct for ServiceError
type ServiceError struct {
	// The error code mapped to the error message.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The category of the error.
	ErrorType *string `json:"errorType,omitempty"`
	// A short explanation of the issue.
	Message *string `json:"message,omitempty"`
	// The PSP reference of the payment.
	PspReference *string `json:"pspReference,omitempty"`
	// The HTTP response status.
	Status *int32 `json:"status,omitempty"`
}

// NewServiceError instantiates a new ServiceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceError() *ServiceError {
	this := ServiceError{}
	return &this
}

// NewServiceErrorWithDefaults instantiates a new ServiceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceErrorWithDefaults() *ServiceError {
	this := ServiceError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ServiceError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceError) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ServiceError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ServiceError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *ServiceError) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceError) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *ServiceError) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *ServiceError) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServiceError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServiceError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServiceError) SetMessage(v string) {
	o.Message = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *ServiceError) GetPspReference() string {
	if o == nil || IsNil(o.PspReference) {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceError) GetPspReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PspReference) {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *ServiceError) HasPspReference() bool {
	if o != nil && !IsNil(o.PspReference) {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *ServiceError) SetPspReference(v string) {
	o.PspReference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceError) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceError) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ServiceError) SetStatus(v int32) {
	o.Status = &v
}

func (o ServiceError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.PspReference) {
		toSerialize["pspReference"] = o.PspReference
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableServiceError struct {
	value *ServiceError
	isSet bool
}

func (v NullableServiceError) Get() *ServiceError {
	return v.value
}

func (v *NullableServiceError) Set(val *ServiceError) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceError) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceError(val *ServiceError) *NullableServiceError {
	return &NullableServiceError{value: val, isSet: true}
}

func (v NullableServiceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


