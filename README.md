## 简介

基于 Go 语言平台的服务端上传的 SDK，通过 SDK 和配合的 Demo，可以将视频和封面文件直接上传到腾讯云点播系统，同时可以指定各项服务端上传的可选参数。

## go get 安装
```
go get -u github.com/tencentcloud/tencentcloud-sdk-go
go get -u github.com/tencentyun/cos-go-sdk-v5
go get -u github.com/tencentyun/vod-go-sdk
```

## 示例

```
package main

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentyun/vod-go-sdk"
	"fmt"
)

func main() {
    client := &vod.VodUploadClient{}
    client.SecretId = os.Getenv("SECRET_ID")
    client.SecretKey = os.Getenv("SECRET_KEY")
    
    req := vod.NewVodUploadRequest()
    req.MediaFilePath = common.StringPtr("video/Wildlife.mp4")
    req.CoverFilePath = common.StringPtr("video/Wildlife-cover.png")
    
    rsp, err := client.Upload("ap-guangzhou", req)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(*rsp.Response.FileId)
    fmt.Println(*rsp.Response.MediaUrl)
    fmt.Println(*rsp.Response.CoverUrl)
}
```
## 设置代理
```
client := &vod.VodUploadClient{}
client.SecretId = os.Getenv("SECRET_ID")
client.SecretKey = os.Getenv("SECRET_KEY")

proxyUrl, err := url.Parse("http://proxyHost:proxyPort")
if err !=nil {
    fmt.Println(err)
    return
}
client.Transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
```