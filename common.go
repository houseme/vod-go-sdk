package vod

import (
	"os"
	"path"
	"strings"
)

// FileExist returns whether the file exists.
func FileExist(target string) bool {
	_, err := os.Stat(target)
	return err == nil
}

// GetFileType returns the file type.
func GetFileType(target string) string {
	fileNameWithSuffix := path.Base(target)
	fileSuffix := path.Ext(fileNameWithSuffix)
	if fileSuffix == "" {
		return fileSuffix
	} else {
		return fileSuffix[1:]
	}
}

// GetFileName returns the file name without suffix.
func GetFileName(target string) string {
	fileNameWithSuffix := path.Base(target)
	fileSuffix := path.Ext(fileNameWithSuffix)
	return strings.TrimSuffix(fileNameWithSuffix, fileSuffix)
}

// IsEmptyStr returns whether the string is empty.
func IsEmptyStr(target *string) bool {
	return target == nil || *target == ""
}

// NotEmptyStr returns whether the string is not empty.
func NotEmptyStr(target *string) bool {
	return !IsEmptyStr(target)
}

// IsManifestMediaType returns whether the media type is manifest.
func IsManifestMediaType(mediaType string) bool {
	if mediaType == "m3u8" || mediaType == "mpd" {
		return true
	}

	return false
}
