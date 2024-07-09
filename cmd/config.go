package main

import "flag"

type Config struct {
	Domain          string `json:"domain"`
	CertFilename    string `json:"certFilename"`
	PrivKeyFilename string `json:"privKeyFilename"`
}

func NewConfigFromFlags() Config {
	var config Config

	flag.StringVar(&config.Domain, "domain", "localhost:8080", "Domain address")
	flag.StringVar(&config.CertFilename, "cert-file", "cert.pem", "Certification filename")
	flag.StringVar(&config.PrivKeyFilename, "key-file", "key.pem", "Private key filename")
	flag.Parse()

	return config
}
