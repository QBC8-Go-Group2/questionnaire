package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type Keys interface {
	GenerateKeys() (*rsa.PrivateKey, []byte, error)
	SavePrivateKeyToFile(privateKey *rsa.PrivateKey, filePath string) error
	SavePublicKeyToFile(publicKeyPEM []byte, filePath string) error
}

type KeyService struct{}

// Generate RSA key pair
func (k *KeyService) GenerateKeys() (*rsa.PrivateKey, []byte, error) {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating private key: %w", err)
	}

	// Marshal public key to PEM
	publicKey := &privateKey.PublicKey
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error marshaling public key: %w", err)
	}

	return privateKey, publicKeyPEM, nil
}

// Save private key to a file
func (k *KeyService) SavePrivateKeyToFile(privateKey *rsa.PrivateKey, filePath string) error {
	// Encode private key to PEM format
	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Write to file
	err := os.WriteFile(filePath, privatePEM, 0600) // Secure permission
	if err != nil {
		return fmt.Errorf("error writing private key to file: %w", err)
	}

	fmt.Printf("Private key saved to %s\n", filePath)
	return nil
}

// Save public key to a file
func (k *KeyService) SavePublicKeyToFile(publicKeyPEM []byte, filePath string) error {
	// Encode public key to PEM format
	publicPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPEM,
	})

	// Write to file
	err := os.WriteFile(filePath, publicPEM, 0644) // Readable by others if needed
	if err != nil {
		return fmt.Errorf("error writing public key to file: %w", err)
	}

	fmt.Printf("Public key saved to %s\n", filePath)
	return nil
}
