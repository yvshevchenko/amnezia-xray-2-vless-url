package internal

import (
	"encoding/json"
	"io"
	"net/url"
	"os"
	"strconv"
)

type SourceConfig struct {
	InboundsSettings []InboundEntity `json:"inbounds"`
	LogSettings      `json:"log"`
	OutboundSettings []OutboundEntity `json:"outbounds"`
}

type InboundEntity struct {
	Listen   string                `json:"listen"`
	Port     int                   `json:"port"`
	Protocol string                `json:"protocol"`
	Settings InboundEntitySettings `json:"settings"`
}

type InboundEntitySettings struct {
	Udp bool `json:"udp"`
}

type LogSettings struct {
	Loglevel string `json:"loglevel"`
}

type OutboundEntity struct {
	Protocol       string                       `json:"protocol"`
	Settings       OutboundEntitySettings       `json:"settings"`
	StreamSettings OutboundEntityStreamSettings `json:"streamSettings"`
}

type OutboundEntitySettings struct {
	Vnext []VnextEntity `json:"vnext"`
}

type VnextEntity struct {
	Address string               `json:"address"`
	Port    int                  `json:"port"`
	Users   []UserSettingsEntity `json:"users"`
}

type UserSettingsEntity struct {
	Encryption string `json:"encryption"`
	Flow       string `json:"flow"`
	Id         string `json:"id"`
}

type OutboundEntityStreamSettings struct {
	Network         string          `json:"network"`
	RealitySettings RealitySettings `json:"realitySettings"`
	Security        string          `json:"security"`
}

type RealitySettings struct {
	Fingerprint string `json:"fingerprint"`
	PublicKey   string `json:"publicKey"`
	ServerName  string `json:"serverName"`
	ShortId     string `json:"shortId"`
	SpiderX     string `json:"spiderX"`
}

func (c SourceConfig) ComposeConnectionUrl() string {
	if c.ConfigIsValid() {
		var u = url.URL{
			Scheme:   c.OutboundSettings[0].Protocol,
			Host:     c.OutboundSettings[0].Settings.Vnext[0].Address + ":" + strconv.Itoa(c.OutboundSettings[0].Settings.Vnext[0].Port),
			User:     url.User(c.OutboundSettings[0].Settings.Vnext[0].Users[0].Id),
			RawQuery: c.ComposeQueryParams(),
		}

		return u.String()
	} else {
		return ""
	}

}

func (c SourceConfig) ComposeQueryParams() string {
	if c.ConfigIsValid() {
		params := make(url.Values)

		params.Add("flow", c.OutboundSettings[0].Settings.Vnext[0].Users[0].Flow)
		params.Add("type", c.OutboundSettings[0].StreamSettings.Network)
		params.Add("security", c.OutboundSettings[0].StreamSettings.Security)
		params.Add("fp", c.OutboundSettings[0].StreamSettings.RealitySettings.Fingerprint)
		params.Add("sni", c.OutboundSettings[0].StreamSettings.RealitySettings.ServerName)
		params.Add("pbk", c.OutboundSettings[0].StreamSettings.RealitySettings.PublicKey)
		params.Add("sid", c.OutboundSettings[0].StreamSettings.RealitySettings.ShortId)

		return params.Encode()
	} else {
		return ""
	}
}

func (c SourceConfig) ConfigIsValid() bool {
	if len(c.OutboundSettings) == 0 ||
		len(c.OutboundSettings[0].Settings.Vnext) == 0 ||
		len(c.OutboundSettings[0].Settings.Vnext[0].Users) == 0 {
		return false
	}

	return true
}

func NewSourceConfig(fileName string) (SourceConfig, error) {
	var cfg SourceConfig

	// opening the file
	file, err := os.Open(fileName)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	// reading the file contents
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return cfg, err
	}

	//unmarshalling file contents to struct
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
