package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/slimcdk/go-shelly"
)

func prettyPrint(emp ...interface{}) {
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}

func main() {

	device := shelly.New("http://192.168.33.1")

	// Fetch device information
	info, err := device.Shelly()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(info.Type, info.Mac)

	// Fetch settings
	settings, err := device.Settings()
	if err != nil {
		log.Fatalln(err)
	}
	prettyPrint(settings.WiFiSTA, settings.WiFiSTA1, settings.MQTT)

	// Update device firmware
	prettyPrint(device.OTAUpdate(shelly.FirmwareUpdate{Update: true}))

	// Set MQTT settings
	_, err = device.SetMQTTSettings(shelly.MQTT{
		Enable: true,
		Server: "mqtt://192.168.1.10:1883",
		User:   "",
		Pass:   "",
		MaxQoS: 1,
		ID:     info.Mac,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to WiFi
	_, err = device.SetWiFiSTA(shelly.WiFiSTA{
		Enabled: true,
		SSID:    "",
		Key:     "",
	})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = device.SetRelay(shelly.RelaySetting{DefaultState: "on"})
	if err != nil {
		log.Fatalln(err)
	}

}
