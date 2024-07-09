package main

import "flag"

type Config struct {
	Domain string `json:"domain"`
}

func NewConfigFromFlags() Config {
	var config Config

	flag.StringVar(&config.Domain, "domain", "localhost:8080", "Domain address")
	flag.Parse()

	return config
}
