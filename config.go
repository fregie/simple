package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr            string   `yaml:"grpc_addr"`
	GrpcGatewayAddr string   `yaml:"grpc_gateway_addr"`
	PromAddr        string   `yaml:"prom_addr"`
	Host            string   `yaml:"host"`
	Sqlite          string   `yaml:"sqlite"`
	Services        []string `yaml:"services"`
}

func parseConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{
		Addr: "127.0.0.1:4433",
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
