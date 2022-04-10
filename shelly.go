package shelly

import "fmt"

type Device struct {
	Type   string `json:"type"`
	Mac    string `json:"mac"`
	Auth   bool   `json:"auth"`
	Fw     string `json:"fw"`
	Longid int64  `json:"longid"`
}

func (c *Client) Device() (Device, error) {

	var result Device

	_, err := c.R().
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(fmt.Sprintf("%s/shelly", c.BaseURL))

	if err != nil {
		return result, err
	}
	return result, nil
}
