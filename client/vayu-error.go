package client

import (
	"fmt"

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
	genericErr, ok := err.(*openapi.GenericOpenAPIError)
	if !ok {
		return BuildVayuError(fmt.Errorf("Unknown error occurred"))
	}

	bytes := genericErr.Body()

	errorTitle := genericErr.Error()
	errorMessage := "Unknown error occurred"

	if bytes != nil && len(bytes) != 0 {
		errorMessage = string(bytes)
	}

	return &VayuError{
		error: genericErr,
		Body:  errorMessage,
		Title: errorTitle,
	}
}
