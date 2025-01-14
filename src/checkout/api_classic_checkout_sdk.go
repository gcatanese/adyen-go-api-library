/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	_context "context"
	_nethttp "net/http"
)

/*
Create a payment session
Provides the data object that can be used to start the Checkout SDK. To set up the payment, pass its amount, currency, and other required parameters. We use this to optimise the payment flow and perform better risk assessment of the transaction.  For more information, refer to [How it works](https://docs.adyen.com/online-payments#howitworks).
 * @param req PaymentSetupRequest - reference of PaymentSetupRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentSetupResponse
*/
func (a Checkout) PaymentSession(req *PaymentSetupRequest, ctxs ..._context.Context) (PaymentSetupResponse, *_nethttp.Response, error) {
	res := &PaymentSetupResponse{}
	path := "/paymentSession"
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+path, ctxs...)
	return *res, httpRes, err
}

/*
Verify a payment result
Verifies the payment result using the payload returned from the Checkout SDK.  For more information, refer to [How it works](https://docs.adyen.com/online-payments#howitworks).
 * @param req PaymentVerificationRequest - reference of PaymentVerificationRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentVerificationResponse
*/
func (a Checkout) VerifyPaymentResult(req *PaymentVerificationRequest, ctxs ..._context.Context) (PaymentVerificationResponse, *_nethttp.Response, error) {
	res := &PaymentVerificationResponse{}
	path := "/payments/result"
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+path, ctxs...)
	return *res, httpRes, err
}
