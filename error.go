package vod

import "fmt"

// ClientError is the error returned by VOD service.
type ClientError struct {
	Message string
}

// Error formats error.
func (e *ClientError) Error() string {
	return fmt.Sprintf("[VodClientError] Message=%s", e.Message)
}

// GetMessage returns the error message.
func (e *ClientError) GetMessage() string {
	return e.Message
}
