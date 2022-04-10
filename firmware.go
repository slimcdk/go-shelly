package shelly

import (
	"fmt"

	"github.com/sonh/qs"
)

// Triggers the device to perform a update over the air firmware update
func (c *Client) OTAUpdate(setting FirmwareUpdate) (FirmwareStatus, error) {

	// Prepare request
	var result FirmwareStatus
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/ota", c.BaseURL))
	return result, err
}

// Fetches the current firmware status of the device
func (c *Client) FirmwareStatus() (FirmwareStatus, error) {

	var result FirmwareStatus

	// Prepare request
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err := req.Get(fmt.Sprintf("%s/ota", c.BaseURL))
	if err != nil {
		return result, err
	}

	return result, nil
}
