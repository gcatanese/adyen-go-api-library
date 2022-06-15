/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v68/payments ```
 *
 * API version: 68
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
import (
	"time"
)
// CreateCheckoutSessionRequest struct for CreateCheckoutSessionRequest
type CreateCheckoutSessionRequest struct {
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`
	AdditionalAmount *Amount `json:"additionalAmount,omitempty"`
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData map[string]string `json:"additionalData,omitempty"`
	// List of payment methods to be presented to the shopper. To refer to payment methods, use their `paymentMethod.type`from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount Amount `json:"amount"`
	ApplicationInfo *ApplicationInfo `json:"applicationInfo,omitempty"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// List of payment methods to be hidden from the shopper. To refer to payment methods, use their `paymentMethod.type`from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The delay between the authorisation and scheduled auto-capture, specified in hours.
	CaptureDelayHours int32 `json:"captureDelayHours,omitempty"`
	// The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the `sdkVersion` or `token`.  Possible values: * iOS * Android * Web
	Channel string `json:"channel,omitempty"`
	Company *Company `json:"company,omitempty"`
	// The shopper's two-letter country code.
	CountryCode string `json:"countryCode,omitempty"`
	// The shopper's date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty"`
	DeliveryAddress *Address `json:"deliveryAddress,omitempty"`
	// When true and `shopperReference` is provided, the shopper will be asked if the payment details should be stored for future one-click payments.
	EnableOneClick bool `json:"enableOneClick,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be tokenized for payouts.
	EnablePayOut bool `json:"enablePayOut,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be tokenized for recurring payments.
	EnableRecurring bool `json:"enableRecurring,omitempty"`
	// The date the session expires in [ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. When not specified, the expiry date is set to 1 hour after session creation. You cannot set the session expiry to more than 24 hours after session creation.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Price and product information about the purchased items, to be included on the invoice sent to the shopper. > This field is required for 3x 4x Oney, Affirm, Afterpay, Clearpay, Klarna, Ratepay, and Zip.
	LineItems *[]LineItem `json:"lineItems,omitempty"`
	Mandate *Mandate `json:"mandate,omitempty"`
	// The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant.
	Mcc string `json:"mcc,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. > We strongly recommend you send the `merchantOrderReference` value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide `retry.orderAttemptNumber`, `retry.chainAttemptNumber`, and `retry.skipRetry` values in `PaymentRequest.additionalData`.
	MerchantOrderReference string `json:"merchantOrderReference,omitempty"`
	// Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request.* Maximum 20 characters per key. * Maximum 80 characters per value. 
	Metadata map[string]string `json:"metadata,omitempty"`
	MpiData *ThreeDSecureData `json:"mpiData,omitempty"`
	// Date after which no further authorisations shall be performed. Only for 3D Secure 2.
	RecurringExpiry string `json:"recurringExpiry,omitempty"`
	// Minimum number of days between authorisations. Only for 3D Secure 2.
	RecurringFrequency string `json:"recurringFrequency,omitempty"`
	// Defines a recurring payment type. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount. 
	RecurringProcessingModel string `json:"recurringProcessingModel,omitempty"`
	// Specifies the redirect method (GET or POST) when redirecting back from the issuer.
	RedirectFromIssuerMethod string `json:"redirectFromIssuerMethod,omitempty"`
	// Specifies the redirect method (GET or POST) when redirecting to the issuer.
	RedirectToIssuerMethod string `json:"redirectToIssuerMethod,omitempty"`
	// The reference to uniquely identify a payment.
	Reference string `json:"reference"`
	// The URL to return to when a redirect payment is completed.
	ReturnUrl string `json:"returnUrl"`
	RiskData *RiskData `json:"riskData,omitempty"`
	// The shopper's email address.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The shopper's IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). > For 3D Secure 2 transactions, schemes require `shopperIP` for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://support.adyen.com/hc/en-us/requests/new).
	ShopperIP string `json:"shopperIP,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction string `json:"shopperInteraction,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	ShopperName *Name `json:"shopperName,omitempty"`
	// Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. > Your reference must not include personally identifiable information (PII), for example name or email address.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The text to be shown on the shopper's bank statement. To enable this field, contact our [Support Team](https://support.adyen.com/hc/en-us/requests/new).  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.
	ShopperStatement string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`
	// Boolean value indicating whether the card payment method should be split into separate debit and credit options.
	SplitCardFundingSources bool `json:"splitCardFundingSources,omitempty"`
	// An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/manage-funds#split).
	Splits *[]Split `json:"splits,omitempty"`
	// The ecommerce or point-of-sale store that is processing the payment.
	Store string `json:"store,omitempty"`
	// When this is set to **true** and the `shopperReference` is provided, the payment details will be stored.
	StorePaymentMethod bool `json:"storePaymentMethod,omitempty"`
	// The shopper's telephone number.
	TelephoneNumber string `json:"telephoneNumber,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation.
	ThreeDSAuthenticationOnly bool `json:"threeDSAuthenticationOnly,omitempty"`
	// Set to true if the payment should be routed to a trusted MID.
	TrustedShopper bool `json:"trustedShopper,omitempty"`
}