package JWT

import (
	"fmt"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/jwt"
	"log"
	"testing"
)

func TestToken(t *testing.T) {
	var keyGen jwt.Keys = &jwt.KeyService{}
	privateKeyPath := "Private_key.pem"
	publicKeyPath := "Public_key.pem"

	// Generate RSA keys
	privateKey, publicKeyPEM, err := keyGen.GenerateKeys()
	if err != nil {
		log.Fatalf("Error generating keys: %v\n", err)
		return
	}

	// Save private key to file
	err = keyGen.SavePrivateKeyToFile(privateKey, privateKeyPath)
	if err != nil {
		log.Fatalf("Error saving private key: %v\n", err)
		return
	}

	// Save public key to file
	err = keyGen.SavePublicKeyToFile(publicKeyPEM, publicKeyPath)
	if err != nil {
		log.Fatalf("Error saving public key: %v\n", err)
		return
	}
	var jwtGen jwt.JWTGenerator = &jwt.JWTService{}
	jwtToken, err := jwtGen.GenerateJWT("admin", 123456)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated JWT: %s\n", jwtToken)
}
