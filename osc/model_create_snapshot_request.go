/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.0
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// CreateSnapshotRequest struct for CreateSnapshotRequest
type CreateSnapshotRequest struct {
	// A description for the snapshot.
	Description string `json:"Description,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The pre-signed URL of the snapshot you want to import from the OSU bucket.
	FileLocation string `json:"FileLocation,omitempty"`
	// The size of the snapshot created in your account, in bytes. This size must be exactly the same as the source snapshot one.
	SnapshotSize int64 `json:"SnapshotSize,omitempty"`
	// The name of the source Region, which must be the same as the Region of your account.
	SourceRegionName string `json:"SourceRegionName,omitempty"`
	// The ID of the snapshot you want to copy.
	SourceSnapshotId string `json:"SourceSnapshotId,omitempty"`
	// The ID of the volume you want to create a snapshot of.
	VolumeId string `json:"VolumeId,omitempty"`
}
