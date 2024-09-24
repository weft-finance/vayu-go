package client

import (
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/openapi"
)

type VayuClient struct {
	accessToken string
	expiresAt   time.Time
	apiKey      string
	Client      *openapi.APIClient
}

func NewVayuClient(APIKey string) *VayuClient {
	client := openapi.NewAPIClient(openapi.NewConfiguration())

	cfg := client.GetConfig()
	cfg.UserAgent = "VayuSDK/go"

	return &VayuClient{Client: client, apiKey: APIKey}
}

func (api *VayuClient) Login() error {
	ctx, cancel := GenerateContextWithTimeout()
	defer cancel()

	request := api.Client.AuthAPI.Login(ctx)
	request = request.LoginRequest(openapi.LoginRequest{RefreshToken: api.apiKey})
	oApiResponse, _, err := request.Execute()

	if err != nil {
		return BuildVayuErrorFromGenericOpenAPIError(err)
	}

	api.accessToken = oApiResponse.AccessToken

	err = api.extractJWTExpiryDate()
	if err != nil {
		return BuildVayuError(err)
	}

	api.buildClientConfig()

	return nil
}

func (api *VayuClient) IsLoggedIn() bool {
	return api.accessToken != ""
}

func (api *VayuClient) ValidateLoggedIn() error {
	if api.IsLoggedIn() {
		return nil
	}

	return BuildVayuError(fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method"))
}
