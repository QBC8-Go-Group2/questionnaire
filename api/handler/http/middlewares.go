package http

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var publicKey *rsa.PublicKey

const (
	EmailKey  = "email"
	RoleKey   = "role"
	UserIDKey = "id"
)

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	publicKey = rsaPubKey
	return rsaPubKey, nil
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		log.Printf("Processing JWT middleware. Auth header present: %v", authHeader != "")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		tokenStr := authHeader[7:]

		// Parse and validate the token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil {
			log.Printf("Token parsing error: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Failed to get claims from token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		// Log all claims for debugging
		log.Printf("Token claims: %+v", claims)

		// Handle user ID
		if id, exists := claims["id"]; exists {
			c.Locals(UserIDKey, fmt.Sprintf("%v", id))
			log.Printf("Set user ID: %v", id)
		}

		// Handle email
		if email, exists := claims["Email"]; exists {
			c.Locals(EmailKey, fmt.Sprintf("%v", email))
			log.Printf("Set email: %v", email)
		}

		// Handle role
		if role, exists := claims["role"]; exists {
			c.Locals(RoleKey, fmt.Sprintf("%v", role))
			log.Printf("Set role: %v", role)
		}

		// Verify token is still valid
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				log.Printf("Token has expired")
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Token has expired",
				})
			}
		}

		log.Printf("JWT middleware completed successfully. UserID: %v, Email: %v, Role: %v",
			c.Locals(UserIDKey), c.Locals(EmailKey), c.Locals(RoleKey))

		return c.Next()
	}
}
