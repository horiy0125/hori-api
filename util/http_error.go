package util

import "fmt"

type HttpError struct {
	Message string `json:"message"`
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("message=%v", he.Message)
}
