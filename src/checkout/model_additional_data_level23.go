/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v52/payments ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// AdditionalDataLevel23 struct for AdditionalDataLevel23
type AdditionalDataLevel23 struct {
	// Customer code, if supplied by a customer. Encoding: ASCII. Max length: 25 characters. > Required for Level 2 and Level 3 data.
	EnhancedSchemeDataCustomerReference string `json:"enhancedSchemeData.customerReference,omitempty"`
	// Total tax amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters. > Required for Level 2 and Level 3 data.
	EnhancedSchemeDataTotalTaxAmount float32 `json:"enhancedSchemeData.totalTaxAmount,omitempty"`
	// Shipping amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters.
	EnhancedSchemeDataFreightAmount float32 `json:"enhancedSchemeData.freightAmount,omitempty"`
	// Duty amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters.
	EnhancedSchemeDataDutyAmount float32 `json:"enhancedSchemeData.dutyAmount,omitempty"`
	// The postal code of a destination address.  Encoding: ASCII. Max length: 10 characters. > Required for American Express.
	EnhancedSchemeDataDestinationPostalCode string `json:"enhancedSchemeData.destinationPostalCode,omitempty"`
	// Destination state or province code.  Encoding: ASCII.Max length: 3 characters.
	EnhancedSchemeDataDestinationStateProvinceCode string `json:"enhancedSchemeData.destinationStateProvinceCode,omitempty"`
	// The postal code of a \"ship-from\" address.  Encoding: ASCII. Max length: 10 characters.
	EnhancedSchemeDataShipFromPostalCode string `json:"enhancedSchemeData.shipFromPostalCode,omitempty"`
	// Destination country code.  Encoding: ASCII. Max length: 3 characters.
	EnhancedSchemeDataDestinationCountryCode string `json:"enhancedSchemeData.destinationCountryCode,omitempty"`
	// Order date. * Format: `ddMMyy`  Encoding: ASCII. Max length: 6 characters.
	EnhancedSchemeDataOrderDate string `json:"enhancedSchemeData.orderDate,omitempty"`
	// Item commodity code. Encoding: ASCII. Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrCommodityCode string `json:"enhancedSchemeData.itemDetailLine[itemNr].commodityCode,omitempty"`
	// Item description. Encoding: ASCII. Max length: 26 characters.
	EnhancedSchemeDataItemDetailLineItemNrDescription string `json:"enhancedSchemeData.itemDetailLine[itemNr].description,omitempty"`
	// Product code. Encoding: ASCII. Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrProductCode string `json:"enhancedSchemeData.itemDetailLine[itemNr].productCode,omitempty"`
	// Quantity, specified as an integer value. Value must be greater than 0. Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrQuantity float32 `json:"enhancedSchemeData.itemDetailLine[itemNr].quantity,omitempty"`
	// Item unit of measurement. Encoding: ASCII. Max length: 3 characters.
	EnhancedSchemeDataItemDetailLineItemNrUnitOfMeasure string `json:"enhancedSchemeData.itemDetailLine[itemNr].unitOfMeasure,omitempty"`
	// Unit price, specified in [minor units](https://docs.adyen.com/development-resources/currency-codes). Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrUnitPrice float32 `json:"enhancedSchemeData.itemDetailLine[itemNr].unitPrice,omitempty"`
	// Discount amount, in minor units.  For example, 2000 means USD 20.00. Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrDiscountAmount float32 `json:"enhancedSchemeData.itemDetailLine[itemNr].discountAmount,omitempty"`
	// Total amount, in minor units. For example, 2000 means USD 20.00. Max length: 12 characters.
	EnhancedSchemeDataItemDetailLineItemNrTotalAmount float32 `json:"enhancedSchemeData.itemDetailLine[itemNr].totalAmount,omitempty"`
}
