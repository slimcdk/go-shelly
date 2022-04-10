package shelly

import "fmt"

func (c *Client) Status() (Status, error) {

	// Prepare request
	var result Status
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err := req.Get(fmt.Sprintf("%s/status", c.BaseURL))
	return result, err
}
