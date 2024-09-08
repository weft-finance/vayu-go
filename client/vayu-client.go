package client

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/openapi"
)

type VayuClient struct {
	accessToken string
	apiKey      string
	Client      *openapi.APIClient
}

type VayuError struct {
	error interface{}
	Body  string
	Title string
}

func (ve *VayuError) Error() string {
	return ve.Title + " - " + ve.Body
}

func BuildVayuError(err error) *VayuError {
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

func BuildVayuErrorFromGenericOpenAPIError(err interface{}) *VayuError {
	genericErr, ok := err.(*openapi.GenericOpenAPIError)
	if !ok {
		return nil
	}

	errorMessage := string(genericErr.Body())

	return &VayuError{
		error: genericErr,
		Body:  errorMessage,
		Title: genericErr.Error(),
	}
}

func NewVayuClient(APIKey string) *VayuClient {
	client := openapi.NewAPIClient(openapi.NewConfiguration())

	cfg := client.GetConfig()
	cfg.UserAgent = "VayuSDK/go"

	return &VayuClient{Client: client, apiKey: APIKey}
}

func (api *VayuClient) Login() *VayuError {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.Client.AuthAPI.Login(ctx)
	request = request.LoginRequest(openapi.LoginRequest{RefreshToken: api.apiKey})
	oApiResponse, _, err := request.Execute()

	if err != nil {
		return BuildVayuErrorFromGenericOpenAPIError(err)
	}

	api.accessToken = oApiResponse.AccessToken

	cfg := api.Client.GetConfig()
	cfg.AddDefaultHeader("Authorization", "Bearer "+api.accessToken)

	return nil
}

func (v *VayuClient) IsLoggedIn() bool {
	return v.accessToken != ""
}

func (v *VayuClient) ValidateLoggedIn() *VayuError {
	if v.IsLoggedIn() {
		return nil
	}

	return BuildVayuError(fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method"))
}
