package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/facilittei/checkout-listener/adapters"
	"github.com/facilittei/checkout-listener/config"
	"github.com/facilittei/checkout-listener/repositories"
)

const (
	authPath = "/authorization-server/oauth/token?grant_type=client_credentials"
)

// AuthHandler authenticates the application
type AuthHandler struct {
	Config         config.Config
	HTTPHandler    adapters.HTTPContract
	AuthRepository repositories.AuthContract
}

// AuthTokenResponse has authorization server response token
type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// NewAuth creates a new instance of an auth provider
func NewAuth(config config.Config, httpHandler adapters.HTTPContract, authRepository repositories.AuthContract) *AuthHandler {
	return &AuthHandler{
		Config:         config,
		HTTPHandler:    httpHandler,
		AuthRepository: authRepository,
	}
}

// GetCredentials retrive resource headers to be used in subsequent requests
func (auth *AuthHandler) GetCredentials() (map[string]string, error) {
	token, err := auth.AuthRepository.GetToken()
	if err != nil {
		return nil, err
	}

	if token != "" {
		return auth.getResourceHeaders(token), nil
	}

	url := auth.Config.PaymentGatewayURL + authPath
	res, err := auth.HTTPHandler.Post(url, nil, auth.getAuthorizationHeaders())
	if err != nil {
		return nil, err
	}

	var authResponse AuthTokenResponse
	json.Unmarshal(res, &authResponse)

	auth.AuthRepository.StoreToken(authResponse.AccessToken)
	return auth.getResourceHeaders(authResponse.AccessToken), nil
}

// getAuthorizationHeaders get authorization server headers
func (auth *AuthHandler) getAuthorizationHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Authorization"] = fmt.Sprintf("Basic %s", auth.Config.PaymentGatewayAuth)
	return headers
}

// getResourceHeaders get resource server headers
func (auth *AuthHandler) getResourceHeaders(token string) map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	headers["X-Api-Version"] = "2"
	headers["X-Resource-Token"] = auth.Config.PaymentGatewayToken
	return headers
}
