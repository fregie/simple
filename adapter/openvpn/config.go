package main

import (
	"encoding/json"
)

type OvpnClientConfig struct {
	Username string
	Password string
	NoGw     bool
	HostID   int
	Config   string
}

func (c *OvpnClientConfig) Index() string {
	return c.Username
}

func (c *OvpnClientConfig) ExportJson() []byte {
	r, _ := json.Marshal(c)
	return r
}
