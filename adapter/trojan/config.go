package main

import (
	"encoding/json"

	"github.com/p4gefau1t/trojan-go/common"
)

type TrojanServerConfig struct {
	LocalAddr string `json:"local_addr"`
	LocalPort int    `json:"local_port"`
	SSL       struct {
		SNI string `json:"sni"`
	} `json:"ssl"`
	Mux struct {
		Enable      bool `json:"enabled"`
		Concurrency int  `json:"concurrency"`
		IdleTimeout int  `json:"idle_timeout"`
	} `json:"mux"`
	Websocket struct {
		Enable bool   `json:"enabled"`
		Path   string `json:"path"`
		Host   string `json:"host"`
	} `json:"websocket"`
	API struct {
		Addr string `json:"api_addr"`
		Port int    `json:"api_port"`
	} `json:"api"`
}

type SSL struct {
	Verify bool   `json:"verify"`
	SNI    string `json:"sni"`
}

type Mux struct {
	Enable      bool `json:"enabled"`
	Concurrency int  `json:"concurrency"`
	IdleTimeout int  `json:"idle_timeout"`
}

type Websocket struct {
	Enable bool   `json:"enabled"`
	Path   string `json:"path"`
	Host   string `json:"host"`
}

type TrojanClientConfig struct {
	RunType    string    `json:"run_type"`
	RemoteAddr string    `json:"remote_addr"`
	RemotePort int       `json:"remote_port"`
	Password   []string  `json:"password"`
	SSL        SSL       `json:"ssl"`
	Mux        Mux       `json:"mux"`
	Websocket  Websocket `json:"websocket"`
}

func (c *TrojanClientConfig) hash() string {
	return common.SHA224String(c.Password[0])
}

func (c *TrojanClientConfig) ExportJson() []byte {
	r, _ := json.Marshal(c)
	return r
}
