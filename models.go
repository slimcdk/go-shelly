package shelly

type Shelly struct {
	Type   string `json:"type"`
	Mac    string `json:"mac"`
	Auth   bool   `json:"auth"`
	Fw     string `json:"fw"`
	Longid int64  `json:"longid"`
}

type Device struct {
	*Shelly
	Hostname string `json:"hostname"`
}

type WiFiAP struct {
	Enabled bool   `json:"enabled"`
	SSID    string `json:"ssíd"`
	Key     string `json:"key"`
}

type WiFiSTA struct {
	Enabled    bool   `json:"enabled" qs:"enabled"`
	SSID       string `json:"ssid" qs:"ssid"`
	IPv4Method string `json:"ipv4_method" qs:"ipv4_method"`
	IP         string `json:",omitempty" qs:"ip"`
	Gateway    string `json:",omitempty" qs:"gateway"`
	Mask       string `json:",omitempty" qs:"netmask"`
	Dns        string `json:",omitempty" qs:"dns"`

	// Specific to setting settings
	Key string `json:"-" qs:"key"`
}

type APRoaming struct {
	Enabled   bool  `json:"enabled"`
	Threshold int64 `json:"threshold"`
}

type MQTT struct {
	Enable              bool    `json:"enable" qs:"mqtt_enable,omitempty"`
	Server              string  `json:"server" qs:"mqtt_server,omitempty"`
	User                string  `json:"user" qs:"mqtt_user,omitempty"`
	ID                  string  `json:"id" qs:"mqtt_id,omitempty"`
	ReconnectTimeoutMax float64 `json:"reconnect_timeout_max" qs:"mqtt_reconnect_timeout_max,omitempty" `
	ReconnectTimeoutMin float64 `json:"reconnect_timeout_min" qs:"mqtt_reconnect_timeout_min,omitempty" `
	CleanSession        bool    `json:"clean_session" qs:"mqtt_clean_session,omitempty"`
	KeepAlive           int64   `json:"keep_alive" qs:"mqtt_keep_alive,omitempty" `
	MaxQoS              int     `json:"max_qos" qs:"mqtt_max_qos,omitempty" `
	Retain              bool    `json:"retain" qs:"mqtt_retain,omitempty" `
	UpdatePeriod        int64   `json:"update_period" qs:"mqtt_update_period,omitempty" `

	// Specific to setting settings
	Pass string `json:"-" qs:"mqtt_pass,omitempty"`
}

type CoIOT struct {
	Enabled      bool   `json:"enabled"`
	UpdatePeriod int64  `json:"update_period"`
	Peer         string `json:"peer"`
}

type SNTP struct {
	Server  string `json:"server"`
	Enabled bool   `json:"enabled"`
}

type Login struct {
	Enabled     bool   `json:"enabled"`
	Unprotected bool   `json:"unprotected"`
	Username    string `json:"username"`
}

type BuildInfo struct {
	BuildID        string `json:"build_id"`
	BuildTimestamp string `json:"build_timestamp"`
	BuildVersion   string `json:"build_version"`
}

type Cloud struct {
	Enabled   bool `json:"enabled"`
	Connected bool `json:"connected"`
}

type Settings struct {
	Device                    Device    `json:"device"`
	WiFiAP                    WiFiAP    `json:"wifi_ap"`
	WiFiSTA                   WiFiSTA   `json:"wifi_sta"`
	WiFiSTA1                  WiFiSTA   `json:"wifi_sta1"`
	APRoaming                 APRoaming `json:"ap_roaming"`
	MQTT                      MQTT      `json:"mqtt"`
	CoIOT                     CoIOT     `json:"coiot"`
	SNTP                      SNTP      `json:"sntp"`
	Login                     Login     `json:"login"`
	PinCode                   string    `json:"pin_code"`
	Name                      string    `json:"name"`
	Fw                        string    `json:"fw"`
	Discoverable              bool      `json:"discoverable"`
	BuildInfo                 BuildInfo `json:"build_info"`
	Cloud                     Cloud     `json:"cloud"`
	Timezone                  string    `json:"timezone"`
	Lat                       float64   `json:"lat"`
	Lng                       float64   `json:"lng"`
	Tzautodetect              bool      `json:"tzautodetect"`
	TZUtcOffset               int64     `json:"tz_utc_offset"`
	TZDst                     bool      `json:"tz_dst"`
	TZDstAuto                 bool      `json:"tz_dst_auto"`
	Time                      string    `json:"time"`
	Unixtime                  int64     `json:"unixtime"`
	LedStatusDisable          bool      `json:"led_status_disable"`
	DebugEnable               bool      `json:"debug_enable"`
	AllowCrossOrigin          bool      `json:"allow_cross_origin"`
	WifiRecoveryRebootEnabled bool      `json:"wifirecovery_reboot_enabled"`
}

type Status struct {
	WiFiSTA struct {
		Connected bool   `json:"connected"`
		SSID      string `json:"ssid"`
		IP        string `json:"ip"`
		RSSI      int    `json:"rssi"`
	} `json:"wifi_sta"`

	MQTT struct {
		Connected bool `json:"connected"`
	} `json:"mqtt"`

	Cloud     Cloud  `json:"cloud"`
	Time      string `json:"time"`
	UnixTime  int64  `json:"unixtime"`
	Serial    int    `json:"serial"`
	HasUpdate bool   `json:"has_update"`
	Mac       string `json:"mac"`
	Update    Update `json:"update"`
	RamTotal  int64  `json:"ram_total"`
	RamFree   int64  `json:"ram_free"`
	RamLwm    int64  `json:"ram_lwm"`
	FSSize    int64  `json:"fs_size"`
	FSFree    int64  `json:"fs_free"`
	Uptime    int64  `json:"uptime"`
}

type Update struct {
	Status     string `json:"status"`
	HasUpdate  bool   `json:"has_update"`
	NewVersion string `json:"new_version"`
	OldVersion string `json:"old_version"`
}

type STASettings struct {
	Enabled    bool   `qs:"enabled"`
	SSID       string `qs:"ssid"`
	Key        string `qs:"key"`
	IPv4Method string `qs:"ipv4_method"`
	IP         string `qs:"ip"`
	Netmask    string `qs:"netmask"`
	Gateway    string `qs:"gateway"`
	Dns        string `qs:"dns"`
}
