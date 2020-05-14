package vod

import (
	"os"
	"path"
	"strings"
)

func FileExist(target string) bool {
	_, err := os.Stat(target)
	return err == nil
}

func GetFileType(target string) string {
	fileNameWithSuffix := path.Base(target)
	fileSuffix := path.Ext(fileNameWithSuffix)
	if fileSuffix == "" {
		return fileSuffix
	} else {
		return fileSuffix[1:]
	}
}

func GetFileName(target string) string {
	fileNameWithSuffix := path.Base(target)
	fileSuffix := path.Ext(fileNameWithSuffix)
	return strings.TrimSuffix(fileNameWithSuffix, fileSuffix)
}

func IsEmptyStr(target *string) bool {
	return target == nil || *target == ""
}

func NotEmptyStr(target *string) bool {
	return !IsEmptyStr(target)
}

func IsManifestMediaType(mediaType string) bool {
	if mediaType == "m3u8" || mediaType == "mpd" {
		return true
	}

	return false
}
