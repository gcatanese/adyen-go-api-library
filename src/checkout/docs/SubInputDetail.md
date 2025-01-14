# SubInputDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]string** | Configuration parameters for the required input. | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) | In case of a select, the items to choose from. | [optional] 
**Key** | Pointer to **string** | The value to provide in the result. | [optional] 
**Optional** | Pointer to **bool** | True if this input is optional to provide. | [optional] 
**Type** | Pointer to **string** | The type of the required input. | [optional] 
**Value** | Pointer to **string** | The value can be pre-filled, if available. | [optional] 

## Methods

### NewSubInputDetail

`func NewSubInputDetail() *SubInputDetail`

NewSubInputDetail instantiates a new SubInputDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubInputDetailWithDefaults

`func NewSubInputDetailWithDefaults() *SubInputDetail`

NewSubInputDetailWithDefaults instantiates a new SubInputDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *SubInputDetail) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SubInputDetail) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SubInputDetail) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SubInputDetail) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetItems

`func (o *SubInputDetail) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubInputDetail) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubInputDetail) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *SubInputDetail) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKey

`func (o *SubInputDetail) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SubInputDetail) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SubInputDetail) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SubInputDetail) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOptional

`func (o *SubInputDetail) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *SubInputDetail) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *SubInputDetail) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *SubInputDetail) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetType

`func (o *SubInputDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubInputDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubInputDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubInputDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SubInputDetail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubInputDetail) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubInputDetail) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SubInputDetail) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


