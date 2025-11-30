package client

import (
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	http    *http.Client
	baseURL string
}

func New() *Client {
	return &Client{
		http: &http.Client{
			Transport: newAuthTransport(),
		},
		baseURL: os.Getenv("PROXMOX_VE_ENDPOINT") + "/api2/json",
	}
}

type authTransport struct {
	base  http.RoundTripper
	token string
}

func newAuthTransport() *authTransport {
	return &authTransport{
		base:  http.DefaultTransport,
		token: os.Getenv("PROXMOX_VE_API_TOKEN"),
	}
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s", t.token))
	return t.base.RoundTrip(req)
}

type response struct {
	Data    Volume `json:"data"`
	Message string `json:"message"`
}
