package vod

import v20180717 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"

// VodUploadRequest is the request struct for api Upload
type VodUploadRequest struct {
	v20180717.ApplyUploadRequest
	MediaFilePath          *string
	CoverFilePath          *string
	ConcurrentUploadNumber *uint64
	MediaUrl               *string
	CoverUrl               *string
}

// VodUploadResponse is the response struct for api Upload
type VodUploadResponse struct {
	v20180717.CommitUploadResponse
}

// NewVodUploadRequest creates a request to invoke Upload API
func NewVodUploadRequest() *VodUploadRequest {
	return &VodUploadRequest{
		ApplyUploadRequest: *v20180717.NewApplyUploadRequest(),
	}
}
