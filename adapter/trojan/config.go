package main

import (
	"encoding/json"
	"fmt"

	"github.com/p4gefau1t/trojan-go/common"
)

type TrojanServerConfig struct {
	LocalAddr string `json:"local_addr"`
	LocalPort int    `json:"local_port"`
	Sqlite    string `json:"sqlite"`
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

func (c *TrojanClientConfig) ExportURL() string {
	var pwd string
	if len(c.Password) > 0 {
		pwd = c.Password[0]
	}
	url := fmt.Sprintf("trojan-go://%s@%s:%d", pwd, c.RemoteAddr, c.RemotePort)

	return url
}

type CustomOption struct {
	SslVerify      bool   `json:"ssl_verify"`
	SslSni         string `json:"ssl_sni"`
	EnableWS       bool   `json:"enable_websocket"`
	WSPath         string `json:"websocket_path"`
	WSHost         string `json:"websocket_host"`
	EnableMux      bool   `json:"enable_mux"`
	MuxConCurrency int    `json:"mux_concurrency"`
	MuxIdleTimeout int    `json:"mux_idle_timeout"`
}
