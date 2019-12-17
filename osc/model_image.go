/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// Image Information about the OMI.
type Image struct {
	// The account alias of the owner of the OMI.
	AccountAlias string `json:"AccountAlias,omitempty"`
	// The account ID of the owner of the OMI.
	AccountId string `json:"AccountId,omitempty"`
	// The architecture of the OMI (by default, `i386`).
	Architecture string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings []BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// The date and time at which the OMI was created.
	CreationDate string `json:"CreationDate,omitempty"`
	// The description of the OMI.
	Description string `json:"Description,omitempty"`
	// The location where the OMI file is stored on Object Storage Unit (OSU).
	FileLocation string `json:"FileLocation,omitempty"`
	// The ID of the OMI.
	ImageId string `json:"ImageId,omitempty"`
	// The name of the OMI.
	ImageName string `json:"ImageName,omitempty"`
	// The type of the OMI.
	ImageType string `json:"ImageType,omitempty"`
	PermissionsToLaunch PermissionsOnResource `json:"PermissionsToLaunch,omitempty"`
	// The product code associated with the OMI (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes []string `json:"ProductCodes,omitempty"`
	// The name of the root device.
	RootDeviceName string `json:"RootDeviceName,omitempty"`
	// The type of root device used by the OMI (always `bsu`).
	RootDeviceType string `json:"RootDeviceType,omitempty"`
	// The state of the OMI.
	State string `json:"State,omitempty"`
	StateComment StateComment `json:"StateComment,omitempty"`
	// One or more tags associated with the OMI.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
