![logo](https://main.qcloudimg.com/raw/af3d70d4f8161afdb4b3364060ffdec4.jpg)
## Overview
The VOD SDK for Go is an SDK for Go encapsulated based on the upload features of VOD. It provides a rich set of upload capabilities to meet your diversified upload needs. In addition, it encapsulates the APIs of VOD, making it easy for you to integrate the upload capabilities without the need to care about underlying details.

## Features
* [x] General file upload
* [x] HLS file upload
* [x] Upload with cover
* [x] Upload to subapplication
* [x] Upload with task flow
* [x] Upload to specified region
* [x] Upload with temporary key
* [x] Upload with proxy

## Documentation
- [Preparations](https://intl.cloud.tencent.com/document/product/266/33912)
- [API documentation](https://intl.cloud.tencent.com/document/product/266/33919)
- [Feature documentation](https://intl.cloud.tencent.com/document/product/266/33919)
- [Error codes](https://intl.cloud.tencent.com/document/product/266/33919)

## Installation
```
go get -u github.com/tencentcloud/tencentcloud-sdk-go
go get -u github.com/tencentyun/cos-go-sdk-v5
go get -u github.com/tencentyun/vod-go-sdk
```

## Test
The SDK provides a wealth of test cases. You can refer to their call methods. For more information, please see [Test Cases](https://github.com/tencentyun/vod-go-sdk/blob/master/client_test.go).
You can view the execution of test cases by running the following command:
```
go test
```

## Release Notes
The changes of each version are recorded in the release notes. For more information, please see [Release Notes](https://github.com/tencentyun/vod-go-sdk/releases).

## Contributors
We appreciate the great support of the following developers to the project, and you are welcome to join us.

<a href="https://github.com/xujianguo"><img width=50 height=50 src="https://avatars1.githubusercontent.com/u/7297536?s=60&v=4" /></a><a href="https://github.com/soulhdb"><img width=50 height=50 src="https://avatars3.githubusercontent.com/u/5770953?s=60&v=4" /></a>

## License
[MIT](https://github.com/tencentyun/vod-go-sdk/blob/master/LICENSE)
