package sscg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func Gen(certFilename, keyFilename string) error {
	// Generate keys
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return fmt.Errorf("failed to gen key: %w", err)
	}

	// Create serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("failed to create serial number: %w", err)
	}

	// Create certificate itself
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"AE4 inc.", "CGSG"},
		},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(3 * time.Hour),

		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		BasicConstraintsValid: true,
	}

	// Save certificate

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create certificate: %w", err)
	}
	pemCert := pem.EncodeToMemory((&pem.Block{Type: "CERTIFICATE", Bytes: derBytes}))
	if pemCert == nil {
		return fmt.Errorf("failed to encode certificate to pem")
	}
	if err := os.WriteFile(certFilename, pemCert, 0644); err != nil {
		return fmt.Errorf("failed to write cert to file: %w", err)
	}

	// Save private key

	keyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %w", err)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if pemKey == nil {
		return fmt.Errorf("failed to encode private key bytes to pem")
	}
	if err := os.WriteFile(keyFilename, pemKey, 0600); err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}

	return nil
}
