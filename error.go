package vod

import "fmt"

// VodClientError is the error returned by VOD service.
type VodClientError struct {
	Message string
}

// Error formats error.
func (e *VodClientError) Error() string {
	return fmt.Sprintf("[VodClientError] Message=%s", e.Message)
}

// GetMessage returns the error message.
func (e *VodClientError) GetMessage() string {
	return e.Message
}
