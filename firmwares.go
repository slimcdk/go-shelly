package shelly

import (
	"fmt"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
)

type Firmware struct {
	File    string `json:"file"`
	Version string `json:"version"`
	OTAUrl  string
}

// AvailableFirmwares fetches available firmwares from shelly-tools.de
func (c *Client) AvailableFirmwares() ([]Firmware, error) {

	// Fetch device information
	info, err := c.Shelly()
	if err != nil {
		return nil, err
	}

	// Prepare request
	var result []Firmware
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fmt.Sprintf("http://archive.shelly-tools.de/archive.php?type=%s", info.Type))
	if err != nil {
		return result, err
	}

	// Sort according to semantic versioning
	sort.Slice(result, func(i, j int) bool {
		return result[j].ToSemVer().LessThan(*result[i].ToSemVer())
	})

	// Construct OTA update link
	for i := range result {
		result[i].OTAUrl = fmt.Sprintf(
			"http://%s/ota?url=http://archive.shelly-tools.de/version/%s/%s",
			strings.Split(c.BaseURL, "//")[1],
			result[i].Version,
			result[i].File,
		)
	}
	return result, nil
}

// Convert string to workable struct
func (fw *Firmware) ToSemVer() *semver.Version {
	return semver.New(fw.Version[1:])
}

func (c *Client) UpdateToLatestFirmware() (Update, error) {

	var result Update

	// Fetch firmwares
	fw, err := c.AvailableFirmwares()
	if err != nil {
		return result, err
	}
	if len(fw) < 1 {
		return result, ErrNoFirmWaresFound
	}

	// Prepare request
	req := c.R()
	req.SetHeader("Accept", "application/json")
	req.SetResult(&result)

	// Perform request
	_, err = req.Get(fw[0].OTAUrl)
	if err != nil {
		return result, err
	}

	return result, nil
}
