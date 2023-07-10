package vod

import v20180717 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"

// UploadRequest is the request struct for api Upload
type UploadRequest struct {
	v20180717.ApplyUploadRequest
	MediaFilePath          *string
	CoverFilePath          *string
	ConcurrentUploadNumber *uint64
	MediaUrl               *string
	CoverUrl               *string
}

// UploadResponse is the response struct for api Upload
type UploadResponse struct {
	v20180717.CommitUploadResponse
}

// NewVodUploadRequest creates a request to invoke Upload API
func NewVodUploadRequest() *UploadRequest {
	return &UploadRequest{
		ApplyUploadRequest: *v20180717.NewApplyUploadRequest(),
	}
}
