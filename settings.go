package shelly

import (
	"fmt"

	"github.com/sonh/qs"
)

func (c *Client) Settings() (Settings, error) {

	// Prepare request
	var result Settings
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err := req.Get(fmt.Sprintf("%s/settings", c.BaseURL))
	return result, err
}

/** WiFi Access Point (server) **/
func (c *Client) SetWiFiAP(setting WiFiAP) (WiFiAP, error) {

	// Prepare request
	var result WiFiAP
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/settings/ap", c.BaseURL))
	return result, err
}

/** WiFi Access Point (client) **/
func (c *Client) SetWiFiSTA(setting WiFiSTA) (WiFiSTA, error) {

	// Prepare request
	var result WiFiSTA
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/settings/sta", c.BaseURL))
	return result, err
}

/** WiFi Access Point Backup (client) **/
func (c *Client) SetWiFiSTABackup(setting WiFiSTA) (WiFiSTA, error) {

	// Prepare request
	var result WiFiSTA
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/settings/sta1", c.BaseURL))
	return result, err
}

/** MQTT **/
func (c *Client) SetMQTTSettings(setting MQTT) (Settings, error) {

	// Prepare request
	var result Settings
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/settings", c.BaseURL))
	return result, err
}
