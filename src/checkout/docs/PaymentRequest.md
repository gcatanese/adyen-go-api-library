# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ApplicationInfo** | Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**AuthenticationData** | Pointer to [**AuthenticationData**](AuthenticationData.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BrowserInfo** | Pointer to [**BrowserInfo**](BrowserInfo.md) |  | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**Channel** | Pointer to **string** | The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the &#x60;sdkVersion&#x60; or &#x60;token&#x60;.  Possible values: * iOS * Android * Web | [optional] 
**CheckoutAttemptId** | Pointer to **string** | Checkout attempt ID that corresponds to the Id generated for tracking user payment journey. | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**ConversionId** | Pointer to **string** | Conversion ID that corresponds to the Id generated for tracking user payment journey. | [optional] 
**CountryCode** | Pointer to **string** | The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** | Pointer to [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**DeliveryDate** | Pointer to **time.Time** | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**DeviceFingerprint** | Pointer to **string** | A string containing the shopper&#39;s device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting). | [optional] 
**EnableOneClick** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the shopper will be asked if the payment details should be stored for future one-click payments. | [optional] 
**EnablePayOut** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for payouts. | [optional] 
**EnableRecurring** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for recurring payments. | [optional] 
**EntityType** | Pointer to **string** | The type of the entity the payment is processed for. | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**IndustryUsage** | Pointer to **string** | The reason for the amount update. Possible values:  * **delayedCharge**  * **noShow**  * **installment** | [optional] 
**Installments** | Pointer to [**Installments**](Installments.md) |  | [optional] 
**LineItems** | Pointer to [**[]LineItem**](LineItem.md) | Price and product information about the purchased items, to be included on the invoice sent to the shopper. &gt; This field is required for 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, Zip and Atome. | [optional] 
**LocalizedShopperStatement** | Pointer to **map[string]string** | This field allows merchants to use dynamic shopper statement in local character sets. The local shopper statement field can be supplied in markets where localized merchant descriptors are used. Currently, Adyen only supports this in the Japanese market .The available character sets at the moment are: * Processing in Japan: **ja-Kana** The character set **ja-Kana** supports UTF-8 based Katakana and alphanumeric and special characters. Merchants should send the Katakana shopperStatement in full-width characters.  An example request would be: &gt; {   \&quot;shopperStatement\&quot; : \&quot;ADYEN - SELLER-A\&quot;,   \&quot;localizedShopperStatement\&quot; : {     \&quot;ja-Kana\&quot; : \&quot;ADYEN - セラーA\&quot;   } } We recommend merchants to always supply the field localizedShopperStatement in addition to the field shopperStatement.It is issuer dependent whether the localized shopper statement field is supported. In the case of non-domestic transactions (e.g. US-issued cards processed in JP) the field &#x60;shopperStatement&#x60; is used to modify the statement of the shopper. Adyen handles the complexity of ensuring the correct descriptors are assigned. | [optional] 
**Mandate** | Pointer to [**Mandate**](Mandate.md) |  | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**MerchantRiskIndicator** | Pointer to [**MerchantRiskIndicator**](MerchantRiskIndicator.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. * Maximum 20 characters per key. * Maximum 80 characters per value.  | [optional] 
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**Order** | Pointer to [**CheckoutOrder**](CheckoutOrder.md) |  | [optional] 
**OrderReference** | Pointer to **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**Origin** | Pointer to **string** | Required for the 3D Secure 2 &#x60;channel&#x60; **Web** integration.  Set this parameter to the origin URL of the page that you are loading the 3D Secure Component from. | [optional] 
**PaymentMethod** | [**CheckoutPaymentMethod**](CheckoutPaymentMethod.md) |  | 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**RecurringExpiry** | Pointer to **string** | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | Pointer to **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**RedirectFromIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting back from the issuer. | [optional] 
**RedirectToIssuerMethod** | Pointer to **string** | Specifies the redirect method (GET or POST) when redirecting to the issuer. | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ReturnUrl** | **string** | The URL to return to in case of a redirection. The format depends on the channel. This URL can have a maximum of 1024 characters. * For web, include the protocol &#x60;http://&#x60; or &#x60;https://&#x60;. You can also include your own additional query parameters, for example, shopper ID or order reference number. Example: &#x60;https://your-company.com/checkout?shopperOrder&#x3D;12xy&#x60; * For iOS, use the custom URL for your app. To know more about setting custom URL schemes, refer to the [Apple Developer documentation](https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/defining_a_custom_url_scheme_for_your_app). Example: &#x60;my-app://&#x60; * For Android, use a custom URL handled by an Activity on your app. You can configure it with an [intent filter](https://developer.android.com/guide/components/intents-filters). Example: &#x60;my-app://your.package.name&#x60; | 
**RiskData** | Pointer to [**RiskData**](RiskData.md) |  | [optional] 
**SessionValidity** | Pointer to **string** | The date and time until when the session remains valid, in [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format.  For example: 2020-07-18T15:42:40.428+01:00 | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperEmail&#x60; for all browser-based and mobile implementations. | [optional] 
**ShopperIP** | Pointer to **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperIP&#x60; for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | Pointer to **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**ShopperStatement** | Pointer to **string** | The text to be shown on the shopper&#39;s bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , &#39; _ - ? + * /_**. | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/add-manage-funds#split). | [optional] 
**Store** | Pointer to **string** | The ecommerce or point-of-sale store that is processing the payment. Used in [partner model integrations](https://docs.adyen.com/marketplaces-and-platforms/classic/platforms-for-partners#route-payments) for Adyen for Platforms. | [optional] 
**StorePaymentMethod** | Pointer to **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be stored. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDS2RequestData** | Pointer to [**ThreeDS2RequestData**](ThreeDS2RequestData.md) |  | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**TrustedShopper** | Pointer to **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest(amount Amount, merchantAccount string, paymentMethod CheckoutPaymentMethod, reference string, returnUrl string, ) *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInfo

`func (o *PaymentRequest) GetAccountInfo() AccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *PaymentRequest) GetAccountInfoOk() (*AccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *PaymentRequest) SetAccountInfo(v AccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *PaymentRequest) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAdditionalData

`func (o *PaymentRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetApplicationInfo

`func (o *PaymentRequest) GetApplicationInfo() ApplicationInfo`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *PaymentRequest) GetApplicationInfoOk() (*ApplicationInfo, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *PaymentRequest) SetApplicationInfo(v ApplicationInfo)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *PaymentRequest) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetAuthenticationData

`func (o *PaymentRequest) GetAuthenticationData() AuthenticationData`

GetAuthenticationData returns the AuthenticationData field if non-nil, zero value otherwise.

### GetAuthenticationDataOk

`func (o *PaymentRequest) GetAuthenticationDataOk() (*AuthenticationData, bool)`

GetAuthenticationDataOk returns a tuple with the AuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationData

`func (o *PaymentRequest) SetAuthenticationData(v AuthenticationData)`

SetAuthenticationData sets AuthenticationData field to given value.

### HasAuthenticationData

`func (o *PaymentRequest) HasAuthenticationData() bool`

HasAuthenticationData returns a boolean if a field has been set.

### GetBillingAddress

`func (o *PaymentRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBrowserInfo

`func (o *PaymentRequest) GetBrowserInfo() BrowserInfo`

GetBrowserInfo returns the BrowserInfo field if non-nil, zero value otherwise.

### GetBrowserInfoOk

`func (o *PaymentRequest) GetBrowserInfoOk() (*BrowserInfo, bool)`

GetBrowserInfoOk returns a tuple with the BrowserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserInfo

`func (o *PaymentRequest) SetBrowserInfo(v BrowserInfo)`

SetBrowserInfo sets BrowserInfo field to given value.

### HasBrowserInfo

`func (o *PaymentRequest) HasBrowserInfo() bool`

HasBrowserInfo returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *PaymentRequest) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *PaymentRequest) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *PaymentRequest) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *PaymentRequest) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetChannel

`func (o *PaymentRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PaymentRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PaymentRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PaymentRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCheckoutAttemptId

`func (o *PaymentRequest) GetCheckoutAttemptId() string`

GetCheckoutAttemptId returns the CheckoutAttemptId field if non-nil, zero value otherwise.

### GetCheckoutAttemptIdOk

`func (o *PaymentRequest) GetCheckoutAttemptIdOk() (*string, bool)`

GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAttemptId

`func (o *PaymentRequest) SetCheckoutAttemptId(v string)`

SetCheckoutAttemptId sets CheckoutAttemptId field to given value.

### HasCheckoutAttemptId

`func (o *PaymentRequest) HasCheckoutAttemptId() bool`

HasCheckoutAttemptId returns a boolean if a field has been set.

### GetCompany

`func (o *PaymentRequest) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PaymentRequest) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PaymentRequest) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PaymentRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConversionId

`func (o *PaymentRequest) GetConversionId() string`

GetConversionId returns the ConversionId field if non-nil, zero value otherwise.

### GetConversionIdOk

`func (o *PaymentRequest) GetConversionIdOk() (*string, bool)`

GetConversionIdOk returns a tuple with the ConversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionId

`func (o *PaymentRequest) SetConversionId(v string)`

SetConversionId sets ConversionId field to given value.

### HasConversionId

`func (o *PaymentRequest) HasConversionId() bool`

HasConversionId returns a boolean if a field has been set.

### GetCountryCode

`func (o *PaymentRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PaymentRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PaymentRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PaymentRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PaymentRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PaymentRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PaymentRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PaymentRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDccQuote

`func (o *PaymentRequest) GetDccQuote() ForexQuote`

GetDccQuote returns the DccQuote field if non-nil, zero value otherwise.

### GetDccQuoteOk

`func (o *PaymentRequest) GetDccQuoteOk() (*ForexQuote, bool)`

GetDccQuoteOk returns a tuple with the DccQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccQuote

`func (o *PaymentRequest) SetDccQuote(v ForexQuote)`

SetDccQuote sets DccQuote field to given value.

### HasDccQuote

`func (o *PaymentRequest) HasDccQuote() bool`

HasDccQuote returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *PaymentRequest) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *PaymentRequest) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *PaymentRequest) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *PaymentRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *PaymentRequest) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *PaymentRequest) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *PaymentRequest) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *PaymentRequest) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *PaymentRequest) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *PaymentRequest) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *PaymentRequest) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *PaymentRequest) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetEnableOneClick

`func (o *PaymentRequest) GetEnableOneClick() bool`

GetEnableOneClick returns the EnableOneClick field if non-nil, zero value otherwise.

### GetEnableOneClickOk

`func (o *PaymentRequest) GetEnableOneClickOk() (*bool, bool)`

GetEnableOneClickOk returns a tuple with the EnableOneClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOneClick

`func (o *PaymentRequest) SetEnableOneClick(v bool)`

SetEnableOneClick sets EnableOneClick field to given value.

### HasEnableOneClick

`func (o *PaymentRequest) HasEnableOneClick() bool`

HasEnableOneClick returns a boolean if a field has been set.

### GetEnablePayOut

`func (o *PaymentRequest) GetEnablePayOut() bool`

GetEnablePayOut returns the EnablePayOut field if non-nil, zero value otherwise.

### GetEnablePayOutOk

`func (o *PaymentRequest) GetEnablePayOutOk() (*bool, bool)`

GetEnablePayOutOk returns a tuple with the EnablePayOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePayOut

`func (o *PaymentRequest) SetEnablePayOut(v bool)`

SetEnablePayOut sets EnablePayOut field to given value.

### HasEnablePayOut

`func (o *PaymentRequest) HasEnablePayOut() bool`

HasEnablePayOut returns a boolean if a field has been set.

### GetEnableRecurring

`func (o *PaymentRequest) GetEnableRecurring() bool`

GetEnableRecurring returns the EnableRecurring field if non-nil, zero value otherwise.

### GetEnableRecurringOk

`func (o *PaymentRequest) GetEnableRecurringOk() (*bool, bool)`

GetEnableRecurringOk returns a tuple with the EnableRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurring

`func (o *PaymentRequest) SetEnableRecurring(v bool)`

SetEnableRecurring sets EnableRecurring field to given value.

### HasEnableRecurring

`func (o *PaymentRequest) HasEnableRecurring() bool`

HasEnableRecurring returns a boolean if a field has been set.

### GetEntityType

`func (o *PaymentRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PaymentRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PaymentRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PaymentRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFraudOffset

`func (o *PaymentRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *PaymentRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *PaymentRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *PaymentRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetIndustryUsage

`func (o *PaymentRequest) GetIndustryUsage() string`

GetIndustryUsage returns the IndustryUsage field if non-nil, zero value otherwise.

### GetIndustryUsageOk

`func (o *PaymentRequest) GetIndustryUsageOk() (*string, bool)`

GetIndustryUsageOk returns a tuple with the IndustryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryUsage

`func (o *PaymentRequest) SetIndustryUsage(v string)`

SetIndustryUsage sets IndustryUsage field to given value.

### HasIndustryUsage

`func (o *PaymentRequest) HasIndustryUsage() bool`

HasIndustryUsage returns a boolean if a field has been set.

### GetInstallments

`func (o *PaymentRequest) GetInstallments() Installments`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *PaymentRequest) GetInstallmentsOk() (*Installments, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *PaymentRequest) SetInstallments(v Installments)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *PaymentRequest) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetLineItems

`func (o *PaymentRequest) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PaymentRequest) GetLineItemsOk() (*[]LineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PaymentRequest) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PaymentRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLocalizedShopperStatement

`func (o *PaymentRequest) GetLocalizedShopperStatement() map[string]string`

GetLocalizedShopperStatement returns the LocalizedShopperStatement field if non-nil, zero value otherwise.

### GetLocalizedShopperStatementOk

`func (o *PaymentRequest) GetLocalizedShopperStatementOk() (*map[string]string, bool)`

GetLocalizedShopperStatementOk returns a tuple with the LocalizedShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedShopperStatement

`func (o *PaymentRequest) SetLocalizedShopperStatement(v map[string]string)`

SetLocalizedShopperStatement sets LocalizedShopperStatement field to given value.

### HasLocalizedShopperStatement

`func (o *PaymentRequest) HasLocalizedShopperStatement() bool`

HasLocalizedShopperStatement returns a boolean if a field has been set.

### GetMandate

`func (o *PaymentRequest) GetMandate() Mandate`

GetMandate returns the Mandate field if non-nil, zero value otherwise.

### GetMandateOk

`func (o *PaymentRequest) GetMandateOk() (*Mandate, bool)`

GetMandateOk returns a tuple with the Mandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandate

`func (o *PaymentRequest) SetMandate(v Mandate)`

SetMandate sets Mandate field to given value.

### HasMandate

`func (o *PaymentRequest) HasMandate() bool`

HasMandate returns a boolean if a field has been set.

### GetMcc

`func (o *PaymentRequest) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *PaymentRequest) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *PaymentRequest) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *PaymentRequest) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *PaymentRequest) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *PaymentRequest) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *PaymentRequest) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *PaymentRequest) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMerchantRiskIndicator

`func (o *PaymentRequest) GetMerchantRiskIndicator() MerchantRiskIndicator`

GetMerchantRiskIndicator returns the MerchantRiskIndicator field if non-nil, zero value otherwise.

### GetMerchantRiskIndicatorOk

`func (o *PaymentRequest) GetMerchantRiskIndicatorOk() (*MerchantRiskIndicator, bool)`

GetMerchantRiskIndicatorOk returns a tuple with the MerchantRiskIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRiskIndicator

`func (o *PaymentRequest) SetMerchantRiskIndicator(v MerchantRiskIndicator)`

SetMerchantRiskIndicator sets MerchantRiskIndicator field to given value.

### HasMerchantRiskIndicator

`func (o *PaymentRequest) HasMerchantRiskIndicator() bool`

HasMerchantRiskIndicator returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMpiData

`func (o *PaymentRequest) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *PaymentRequest) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *PaymentRequest) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *PaymentRequest) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetOrder

`func (o *PaymentRequest) GetOrder() CheckoutOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaymentRequest) GetOrderOk() (*CheckoutOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaymentRequest) SetOrder(v CheckoutOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaymentRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderReference

`func (o *PaymentRequest) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *PaymentRequest) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *PaymentRequest) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.

### HasOrderReference

`func (o *PaymentRequest) HasOrderReference() bool`

HasOrderReference returns a boolean if a field has been set.

### GetOrigin

`func (o *PaymentRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PaymentRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PaymentRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PaymentRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentRequest) GetPaymentMethod() CheckoutPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequest) GetPaymentMethodOk() (*CheckoutPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequest) SetPaymentMethod(v CheckoutPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPlatformChargebackLogic

`func (o *PaymentRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *PaymentRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *PaymentRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *PaymentRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetRecurringExpiry

`func (o *PaymentRequest) GetRecurringExpiry() string`

GetRecurringExpiry returns the RecurringExpiry field if non-nil, zero value otherwise.

### GetRecurringExpiryOk

`func (o *PaymentRequest) GetRecurringExpiryOk() (*string, bool)`

GetRecurringExpiryOk returns a tuple with the RecurringExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpiry

`func (o *PaymentRequest) SetRecurringExpiry(v string)`

SetRecurringExpiry sets RecurringExpiry field to given value.

### HasRecurringExpiry

`func (o *PaymentRequest) HasRecurringExpiry() bool`

HasRecurringExpiry returns a boolean if a field has been set.

### GetRecurringFrequency

`func (o *PaymentRequest) GetRecurringFrequency() string`

GetRecurringFrequency returns the RecurringFrequency field if non-nil, zero value otherwise.

### GetRecurringFrequencyOk

`func (o *PaymentRequest) GetRecurringFrequencyOk() (*string, bool)`

GetRecurringFrequencyOk returns a tuple with the RecurringFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFrequency

`func (o *PaymentRequest) SetRecurringFrequency(v string)`

SetRecurringFrequency sets RecurringFrequency field to given value.

### HasRecurringFrequency

`func (o *PaymentRequest) HasRecurringFrequency() bool`

HasRecurringFrequency returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *PaymentRequest) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *PaymentRequest) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *PaymentRequest) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *PaymentRequest) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetRedirectFromIssuerMethod

`func (o *PaymentRequest) GetRedirectFromIssuerMethod() string`

GetRedirectFromIssuerMethod returns the RedirectFromIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectFromIssuerMethodOk

`func (o *PaymentRequest) GetRedirectFromIssuerMethodOk() (*string, bool)`

GetRedirectFromIssuerMethodOk returns a tuple with the RedirectFromIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFromIssuerMethod

`func (o *PaymentRequest) SetRedirectFromIssuerMethod(v string)`

SetRedirectFromIssuerMethod sets RedirectFromIssuerMethod field to given value.

### HasRedirectFromIssuerMethod

`func (o *PaymentRequest) HasRedirectFromIssuerMethod() bool`

HasRedirectFromIssuerMethod returns a boolean if a field has been set.

### GetRedirectToIssuerMethod

`func (o *PaymentRequest) GetRedirectToIssuerMethod() string`

GetRedirectToIssuerMethod returns the RedirectToIssuerMethod field if non-nil, zero value otherwise.

### GetRedirectToIssuerMethodOk

`func (o *PaymentRequest) GetRedirectToIssuerMethodOk() (*string, bool)`

GetRedirectToIssuerMethodOk returns a tuple with the RedirectToIssuerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToIssuerMethod

`func (o *PaymentRequest) SetRedirectToIssuerMethod(v string)`

SetRedirectToIssuerMethod sets RedirectToIssuerMethod field to given value.

### HasRedirectToIssuerMethod

`func (o *PaymentRequest) HasRedirectToIssuerMethod() bool`

HasRedirectToIssuerMethod returns a boolean if a field has been set.

### GetReference

`func (o *PaymentRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetReturnUrl

`func (o *PaymentRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetRiskData

`func (o *PaymentRequest) GetRiskData() RiskData`

GetRiskData returns the RiskData field if non-nil, zero value otherwise.

### GetRiskDataOk

`func (o *PaymentRequest) GetRiskDataOk() (*RiskData, bool)`

GetRiskDataOk returns a tuple with the RiskData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskData

`func (o *PaymentRequest) SetRiskData(v RiskData)`

SetRiskData sets RiskData field to given value.

### HasRiskData

`func (o *PaymentRequest) HasRiskData() bool`

HasRiskData returns a boolean if a field has been set.

### GetSessionValidity

`func (o *PaymentRequest) GetSessionValidity() string`

GetSessionValidity returns the SessionValidity field if non-nil, zero value otherwise.

### GetSessionValidityOk

`func (o *PaymentRequest) GetSessionValidityOk() (*string, bool)`

GetSessionValidityOk returns a tuple with the SessionValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidity

`func (o *PaymentRequest) SetSessionValidity(v string)`

SetSessionValidity sets SessionValidity field to given value.

### HasSessionValidity

`func (o *PaymentRequest) HasSessionValidity() bool`

HasSessionValidity returns a boolean if a field has been set.

### GetShopperEmail

`func (o *PaymentRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *PaymentRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *PaymentRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *PaymentRequest) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperIP

`func (o *PaymentRequest) GetShopperIP() string`

GetShopperIP returns the ShopperIP field if non-nil, zero value otherwise.

### GetShopperIPOk

`func (o *PaymentRequest) GetShopperIPOk() (*string, bool)`

GetShopperIPOk returns a tuple with the ShopperIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperIP

`func (o *PaymentRequest) SetShopperIP(v string)`

SetShopperIP sets ShopperIP field to given value.

### HasShopperIP

`func (o *PaymentRequest) HasShopperIP() bool`

HasShopperIP returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *PaymentRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *PaymentRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *PaymentRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *PaymentRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *PaymentRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *PaymentRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *PaymentRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *PaymentRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *PaymentRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *PaymentRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *PaymentRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *PaymentRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *PaymentRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *PaymentRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *PaymentRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *PaymentRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *PaymentRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *PaymentRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *PaymentRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *PaymentRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *PaymentRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplits

`func (o *PaymentRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PaymentRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PaymentRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *PaymentRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *PaymentRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetStorePaymentMethod

`func (o *PaymentRequest) GetStorePaymentMethod() bool`

GetStorePaymentMethod returns the StorePaymentMethod field if non-nil, zero value otherwise.

### GetStorePaymentMethodOk

`func (o *PaymentRequest) GetStorePaymentMethodOk() (*bool, bool)`

GetStorePaymentMethodOk returns a tuple with the StorePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePaymentMethod

`func (o *PaymentRequest) SetStorePaymentMethod(v bool)`

SetStorePaymentMethod sets StorePaymentMethod field to given value.

### HasStorePaymentMethod

`func (o *PaymentRequest) HasStorePaymentMethod() bool`

HasStorePaymentMethod returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *PaymentRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *PaymentRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *PaymentRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *PaymentRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThreeDS2RequestData

`func (o *PaymentRequest) GetThreeDS2RequestData() ThreeDS2RequestData`

GetThreeDS2RequestData returns the ThreeDS2RequestData field if non-nil, zero value otherwise.

### GetThreeDS2RequestDataOk

`func (o *PaymentRequest) GetThreeDS2RequestDataOk() (*ThreeDS2RequestData, bool)`

GetThreeDS2RequestDataOk returns a tuple with the ThreeDS2RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2RequestData

`func (o *PaymentRequest) SetThreeDS2RequestData(v ThreeDS2RequestData)`

SetThreeDS2RequestData sets ThreeDS2RequestData field to given value.

### HasThreeDS2RequestData

`func (o *PaymentRequest) HasThreeDS2RequestData() bool`

HasThreeDS2RequestData returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *PaymentRequest) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *PaymentRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *PaymentRequest) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *PaymentRequest) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.

### GetTrustedShopper

`func (o *PaymentRequest) GetTrustedShopper() bool`

GetTrustedShopper returns the TrustedShopper field if non-nil, zero value otherwise.

### GetTrustedShopperOk

`func (o *PaymentRequest) GetTrustedShopperOk() (*bool, bool)`

GetTrustedShopperOk returns a tuple with the TrustedShopper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedShopper

`func (o *PaymentRequest) SetTrustedShopper(v bool)`

SetTrustedShopper sets TrustedShopper field to given value.

### HasTrustedShopper

`func (o *PaymentRequest) HasTrustedShopper() bool`

HasTrustedShopper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


