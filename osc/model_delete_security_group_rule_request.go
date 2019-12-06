/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// DeleteSecurityGroupRuleRequest struct for DeleteSecurityGroupRuleRequest
type DeleteSecurityGroupRuleRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The direction of the flow: `Inbound` or `Outbound`. You can specify `Outbound` for Nets only.
	Flow string `json:"Flow"`
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`) or protocol number. By default, `-1`, which means all protocols.
	IpProtocol string `json:"IpProtocol,omitempty"`
	// The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange,omitempty"`
	// One or more rules you want to delete from the security group.
	Rules []SecurityGroupRule `json:"Rules,omitempty"`
	// The account ID of the owner of the security group you want to delete a rule from.
	SecurityGroupAccountIdToUnlink string `json:"SecurityGroupAccountIdToUnlink,omitempty"`
	// The ID of the security group you want to delete a rule from.
	SecurityGroupId string `json:"SecurityGroupId"`
	// The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.
	SecurityGroupNameToUnlink string `json:"SecurityGroupNameToUnlink,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP type number.
	ToPortRange int32 `json:"ToPortRange,omitempty"`
}
