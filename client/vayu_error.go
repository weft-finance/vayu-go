package client

import (
	"fmt"

	"net"
	"net/url"

	"github.com/weft-finance/vayu-go/openapi"
)

type VayuError struct {
	error interface{}
	Body  string
	Title string
}

func (ve *VayuError) Error() string {
	genericError := "Internal Error - Unknown error occurred"

	if ve.error == nil {
		return genericError
	}

	return ve.Title + " - " + ve.Body
}

func BuildVayuError(err error) error {
	if err == nil {
		return nil
	}

	vayuError := &VayuError{
		error: err,
		Body:  err.Error(),
		Title: "Internal Error",
	}

	return vayuError
}

func BuildVayuErrorFromGenericOpenAPIError(err interface{}) error {
	if netErr, ok := err.(net.Error); ok {
		return BuildVayuError(fmt.Errorf("network error: %v", netErr))
	}

	if urlErr, ok := err.(*url.Error); ok {
		return BuildVayuError(fmt.Errorf("connection error: %v", urlErr))
	}

	genericErr, ok := err.(*openapi.GenericOpenAPIError)

	if !ok {
		return BuildVayuError(fmt.Errorf("unknown error occurred"))
	}

	bytes := genericErr.Body()

	errorTitle := genericErr.Error()
	errorMessage := "Unknown error occurred"

	if len(bytes) != 0 {
		errorMessage = string(bytes)
	}

	return &VayuError{
		error: genericErr,
		Body:  errorMessage,
		Title: errorTitle,
	}
}
