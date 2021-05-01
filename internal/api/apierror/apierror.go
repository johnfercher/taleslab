package apierror

import (
	"context"
	"fmt"
	"github.com/johnfercher/taleslab/internal/api/apilogger"
)

// ApiError response model
// swagger:response errRes
type ApiError interface {
	// Http Status Code
	ErrorStatusCode() int
	// Error message
	Error() string
	// Set causes
	SetCauses(interface{}) ApiError
}

func New(httpStatusCode int, message string) *apiError {
	return &apiError{
		HttpCode: httpStatusCode,
		Message:  message,
	}
}

type apiError struct {
	Message  string      `json:"message"`
	HttpCode int         `json:"http_code"`
	Causes   interface{} `json:"causes"`
}

func (self apiError) ErrorStatusCode() int {
	return self.HttpCode
}

func (self apiError) Error() string {
	return self.Message
}

func (self apiError) SetCauses(causes interface{}) ApiError {
	self.Causes = causes
	return self
}

func Log(ctx context.Context, err ApiError) {
	apilogger.Error(ctx, fmt.Sprintf("status_code: %d, message: %s", err.ErrorStatusCode(), err.Error()))
}
