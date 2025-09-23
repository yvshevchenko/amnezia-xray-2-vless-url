package internal

import (
	"net/url"
)

func (c SourceConfig) ComposeConnectionUrl() string {

	var u = url.URL{
		Scheme:   c.OutboundSettings[0].Protocol,
		Host:     c.OutboundSettings[0].Settings.Vnext[0].Address,
		User:     url.User(c.OutboundSettings[0].Settings.Vnext[0].Users[0].Id),
		RawQuery: c.ComposeQueryParams(),
	}

	return u.String()
}

func (c SourceConfig) ComposeQueryParams() string {
	params := make(url.Values)

	params.Add("flow", c.OutboundSettings[0].Settings.Vnext[0].Users[0].Flow)
	params.Add("type", c.OutboundSettings[0].StreamSettings.Network)
	params.Add("security", c.OutboundSettings[0].StreamSettings.Security)
	params.Add("fp", c.OutboundSettings[0].StreamSettings.RealitySettings.Fingerprint)
	params.Add("sni", c.OutboundSettings[0].StreamSettings.RealitySettings.ServerName)
	params.Add("pbk", c.OutboundSettings[0].StreamSettings.RealitySettings.PublicKey)
	params.Add("sid", c.OutboundSettings[0].StreamSettings.RealitySettings.ShortId)

	return params.Encode()
}
