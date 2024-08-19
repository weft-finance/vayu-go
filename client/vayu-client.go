package client

import (
	"context"
	"time"

	"github.com/weft-finance/vayu-go/openapi"
)

type VayuClient struct {
	accessToken string
	apiKey      string
	Client      *openapi.APIClient
}

func NewVayuClient(APIKey string) *VayuClient {
	client := openapi.NewAPIClient(openapi.NewConfiguration())

	cfg := client.GetConfig()
	cfg.UserAgent = "VayuSDK/go"

	return &VayuClient{Client: client, apiKey: APIKey}
}

func (c *VayuClient) Login() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.Client.AuthAPI.Login(ctx)
	request = request.LoginRequest(openapi.LoginRequest{RefreshToken: c.apiKey})
	oApiResponse, _, err := request.Execute()

	if err != nil {
		return err
	}

	c.accessToken = oApiResponse.AccessToken

	cfg := c.Client.GetConfig()
	cfg.AddDefaultHeader("Authorization", "Bearer "+c.accessToken)

	return nil
}

func (v *VayuClient) IsLoggedIn() bool {
	return v.accessToken != ""
}
