package vod

import "fmt"

type VodClientError struct {
	Message string
}

func (e *VodClientError) Error() string {
	return fmt.Sprintf("[VodClientError] Message=%s", e.Message)
}

func (e *VodClientError) GetMessage() string {
	return e.Message
}
