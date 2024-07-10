package cg

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

type CertificateGeneratorFunc func(string, string) error

// Check of this certificate is valid
func CheckCert(certFilename string) error {
	bytes, err := os.ReadFile(certFilename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return fmt.Errorf("failed to decode pem file")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: %w", err)
	}
	if time.Now().After(cert.NotAfter) {
		return fmt.Errorf("certificate is out of date")
	}
	// Maybe later I will add some other checks
	return nil
}
