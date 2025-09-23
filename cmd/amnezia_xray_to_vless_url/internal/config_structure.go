package internal

type SourceConfig struct {
	InboundsSettings []InboundEntity `json:"inbounds"`
	LogSettings      `json:"log"`
	OutboundSettings []OutboundEntity `json:"outbounds"`
}

type InboundEntity struct {
	Listen   string                `json:"listen"`
	Port     uint16                `json:"port"`
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
	Port    uint16               `json:"port"`
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
