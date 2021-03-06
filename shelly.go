package shelly

import "fmt"

func (c *Client) Shelly() (Shelly, error) {

	// Prepare request
	var result Shelly
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err := req.Get(fmt.Sprintf("%s/shelly", c.BaseURL))
	return result, err
}

func (c *Client) Ping() error {
	return nil
}
