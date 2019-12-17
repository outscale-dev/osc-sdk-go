/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.16
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// Quota Information about the quota.
type Quota struct {
	// The account ID of the owner of the quotas.
	AccountId string `json:"AccountId,omitempty"`
	// The description of the quota.
	Description string `json:"Description,omitempty"`
	// The maximum value of the quota for the 3DS OUTSCALE user account (if there is no limit, `0`).
	MaxValue int32 `json:"MaxValue,omitempty"`
	// The unique name of the quota.
	Name string `json:"Name,omitempty"`
	// The group name of the quota.
	QuotaCollection string `json:"QuotaCollection,omitempty"`
	// The description of the quota.
	ShortDescription string `json:"ShortDescription,omitempty"`
	// The limit value currently used by the 3DS OUTSCALE user account.
	UsedValue int32 `json:"UsedValue,omitempty"`
}
