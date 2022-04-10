package shelly

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func New(address string) *Client {
	return NewWithClient(address, http.DefaultClient)
}

func NewWithClient(address string, hc *http.Client) *Client {
	rc := *resty.NewWithClient(hc)
	rc.SetScheme("http")
	rc.SetBaseURL(address)
	return &Client{&rc}
}
