package client

import (
	"context"
	"time"

	"github.com/weft-finance/vayu-go/openapi"
)

type VayuClient struct {
	accessToken string
	APIKey      string
	Client      *openapi.APIClient
}

func NewVayuClient(APIKey string) *VayuClient {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	cfg := client.GetConfig()
	cfg.UserAgent = "VayuSDK/go"
	cfg.AddDefaultHeader("Authorization", "Bearer"+APIKey)

	return &VayuClient{Client: client, APIKey: APIKey}
}

func (c *VayuClient) Login() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	login := c.Client.AuthAPI.Login(ctx)
	login.LoginRequest(openapi.LoginRequest{RefreshToken: c.APIKey})
	oApiResponse, _, err := login.Execute()

	if err != nil {
		return err
	}

	c.accessToken = oApiResponse.AccessToken
	return nil
}

func (v *VayuClient) IsLoggedIn() bool {
	return v.accessToken != ""
}
