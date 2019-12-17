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
// SnapshotExportTask Information about the snapshot export task.
type SnapshotExportTask struct {
	// If the snapshot export task fails, an error message appears.
	Comment string `json:"Comment,omitempty"`
	OsuExport OsuExport `json:"OsuExport,omitempty"`
	// The progress of the snapshot export task, as a percentage.
	Progress int32 `json:"Progress,omitempty"`
	// The ID of the snapshot to be exported.
	SnapshotId string `json:"SnapshotId,omitempty"`
	// The state of the snapshot export task (`pending` \\| `active` \\| `completed` \\| `failed`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the snapshot export task.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the snapshot export task.
	TaskId string `json:"TaskId,omitempty"`
}
