package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (api *VayuClient) extractJWTExpiryDate() error {
	token, _, err := new(jwt.Parser).ParseUnverified(api.accessToken, jwt.MapClaims{})
	if err != nil {
		return BuildVayuError(fmt.Errorf("failed to parse JWT"))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			api.expiresAt = time.Unix(int64(exp), 0)
		} else {
			return BuildVayuError(fmt.Errorf("expiry date not found in JWT"))
		}
	}

	return nil
}

func (api *VayuClient) buildClientConfig() {
	cfg := api.Client.GetConfig()
	cfg.AddDefaultHeader("Authorization", "Bearer "+api.accessToken)

	if cfg.HTTPClient.Transport == nil {
		cfg.HTTPClient.Transport = http.DefaultTransport
	}

	cfg.HTTPClient.Transport = &vayuRoundTripper{
		rt:  cfg.HTTPClient.Transport,
		api: api,
	}
}

type vayuRoundTripper struct {
	rt  http.RoundTripper
	api *VayuClient
}

const TokenExpiryThreshold = 5 * time.Minute

func (c *vayuRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path == "/login" {
		return c.rt.RoundTrip(req)
	}

	if err := c.api.ValidateLoggedIn(); err != nil {
		return nil, err
	}

	if time.Now().Add(TokenExpiryThreshold).After(c.api.expiresAt) {
		err := c.api.Login()
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Authorization", "Bearer "+c.api.accessToken)

	return c.rt.RoundTrip(req)
}
