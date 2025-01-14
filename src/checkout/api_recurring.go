/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	_context "context"
	_nethttp "net/http"
	"strings"
)

/*
Delete a token for stored payment details
Deletes the token identified in the path. The token can no longer be used with payment requests.
 * @param recurringId The unique identifier of the token.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return StoredPaymentMethodResource
*/
func (a Checkout) DeleteTokenForStoredPaymentDetails(recurringId *string, ctxs ..._context.Context) (StoredPaymentMethodResource, *_nethttp.Response, error) {
	res := &StoredPaymentMethodResource{}
	path := "/storedPaymentMethods/{recurringId}"
	path = strings.ReplaceAll(path, "{"+"recurringId"+"}", *recurringId)
	httpRes, err := a.Client.MakeHTTPDeleteRequest(res, a.BasePath()+path, ctxs...)
	return *res, httpRes, err
}

/*
Get tokens for stored payment details
Lists the tokens for stored payment details for the shopper identified in the path, if there are any available. The token ID can be used with payment requests for the shopper&#39;s payment. A summary of the stored details is included.
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ListStoredPaymentMethodsResponse
*/
func (a Checkout) GetTokensForStoredPaymentDetails(ctxs ..._context.Context) (ListStoredPaymentMethodsResponse, *_nethttp.Response, error) {
	res := &ListStoredPaymentMethodsResponse{}
	path := "/storedPaymentMethods"
	httpRes, err := a.Client.MakeHTTPGetRequest(res, a.BasePath()+path, ctxs...)
	return *res, httpRes, err
}
