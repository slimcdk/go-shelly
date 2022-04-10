package shelly

import "fmt"

func (c *Client) Reboot() error {
	_, err := c.R().Get(fmt.Sprintf("%s/reboot", c.BaseURL))
	return err
}
