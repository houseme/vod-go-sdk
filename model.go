package vod

import v20180717 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"

type VodUploadRequest struct {
	v20180717.ApplyUploadRequest
	MediaFilePath          *string
	CoverFilePath          *string
	ConcurrentUploadNumber *uint64
	MediaUrl               *string
	CoverUrl               *string
}

type VodUploadResponse struct {
	v20180717.CommitUploadResponse
}

func NewVodUploadRequest() (request *VodUploadRequest) {
	return &VodUploadRequest{
		ApplyUploadRequest: *v20180717.NewApplyUploadRequest(),
	}
}
