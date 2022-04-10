package shelly

import (
	"fmt"

	"github.com/sonh/qs"
)

type RelayStatus struct {
	Ison          bool     `json:"ison"`
	HasTimer      bool     `json:"has_timer"`
	Overpower     bool     `json:"overpower"`
	DefaultState  string   `json:"default_state"`
	AutoOn        float64  `json:"auto_on"`
	AutoOff       float64  `json:"auto_off"`
	BtnOnUrl      string   `json:"btn_on_url"`
	OutOnUrl      string   `json:"out_on_url"`
	OutOffUrl     string   `json:"out_off_url"`
	Schedule      bool     `json:"schedule"`
	ScheduleRules []string `json:"schedule_rules"`
}

type RelaySetting struct {
	Index         uint64
	Name          string   `qs:"name,omitempty"`
	ApplianceType string   `qs:"appliance_type,omitempty"`
	Reset         bool     `qs:"reset,omitempty"`
	DefaultState  string   `qs:"default_state,omitempty"`
	AutoOn        float64  `qs:"auto_on,omitempty"`
	AutoOff       float64  `qs:"auto_off,omitempty"`
	Schedule      bool     `qs:"schedule,omitempty"`
	ScheduleRules []string `qs:"schedule_rules,omitempty"`
	MaxPower      int64    `qs:"max_power,omitempty"`
}

func (c *Client) SetRelay(setting RelaySetting) (RelayStatus, error) {

	// Prepare request
	var result RelayStatus
	req := c.R()
	req.SetHeader("Accept", "application/json")

	values, err := qs.NewEncoder().Values(&setting)
	if err != nil {
		return result, err
	}
	req.SetQueryString(values.Encode())
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("%s/settings/relay/%d", c.BaseURL, setting.Index))
	return result, err
}
