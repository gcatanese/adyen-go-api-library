/*
 * Adyen for Platforms: Notification Configuration API
 *
 * The Notification Configuration API provides endpoints for setting up and testing notifications that inform you of events on your platform, for example when a KYC check or a payout has been completed.  For more information, refer to our [documentation](https://docs.adyen.com/marketpay/marketpay-notifications). ## Authentication To connect to the Notification Configuration API, you must use basic authentication credentials of your web service user. If you don't have one, contact our [Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Notification Configuration API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Notification/v1/createNotificationConfiguration ```
 *
 * API version: 5
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsnotificationconfiguration
// DeleteNotificationConfigurationRequest struct for DeleteNotificationConfigurationRequest
type DeleteNotificationConfigurationRequest struct {
	// A list of IDs of the notification subscription configurations to be deleted.
	NotificationIds []int64 `json:"notificationIds"`
}
