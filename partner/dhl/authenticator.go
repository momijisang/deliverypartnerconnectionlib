package dhl

import "fmt"

type dhlAuthenticator struct {
	dhlAuthenticationAPI DHLAuthenticationAPI
	clientID             string
	password             string
}

func NewDHLAuthenticator(
	dhlAuthenticationAPI DHLAuthenticationAPI,
	clientID string,
	password string) Authenticator {
	return &dhlAuthenticator{
		dhlAuthenticationAPI: dhlAuthenticationAPI,
		clientID:             clientID,
		password:             password,
	}
}

func (svc *dhlAuthenticator) Authenticate() (accessToken string, err error) {
	request := DHLAuthenticationAPIRequest{}
	response, err := svc.dhlAuthenticationAPI.Get("/rest/v1/OAuth/AccessToken?clientId="+svc.clientID+"&password="+svc.password, map[string]string{}, request)
	if err != nil {
		return "", fmt.Errorf("failed to authenticate with dhl: %w", err)
	}
	return response.AccessTokenResponse.Token, nil
}
