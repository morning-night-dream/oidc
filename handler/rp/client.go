package rp

import (
	"fmt"
	"net/http"
)

type AuthorizationTransport struct {
	AccessToken string
	Transport   http.RoundTripper
}

func NewAuthorizationTransport(
	accessToken string,
) *AuthorizationTransport {
	return &AuthorizationTransport{
		AccessToken: accessToken,
		Transport:   http.DefaultTransport,
	}
}

func (at *AuthorizationTransport) transport() http.RoundTripper {
	if at.Transport != nil {
		return at.Transport
	}

	return at.Transport
}

func (at *AuthorizationTransport) RoundTrip(
	req *http.Request,
) (*http.Response, error) {
	value := fmt.Sprintf("Bearer %s", at.AccessToken)

	req.Header.Set("Authorization", value)

	resp, err := at.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
