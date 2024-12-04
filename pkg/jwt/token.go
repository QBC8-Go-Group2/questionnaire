package jwt

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWTGenerator interface {
	GenerateJWT(role string, userId uint) (string, error)
}

type JWTService struct{}

// Generate a JWT using the private key
func (j *JWTService) GenerateJWT(role string, userId uint) (string, error) {
	// Read the file
	privateKeyData, err := os.ReadFile("Private_key.pem")
	if err != nil {
		return "error reading private key file:", err
	}

	// Decode the PEM block
	block, _ := pem.Decode(privateKeyData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "invalid private key PEM file", nil
	}

	// Parse the RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "error parsing private key", err
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":   userId,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
