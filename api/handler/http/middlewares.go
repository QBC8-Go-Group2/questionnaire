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

type contextKey string

const (
	EmailKey  contextKey = "Email"
	RoleKey   contextKey = "role"
	UserIDKey contextKey = "id"
)

// LoadPublicKey loads the RSA public key from a .pem file
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

	// Set the package-level publicKey variable
	publicKey = rsaPubKey

	return rsaPubKey, nil
}

// JWTMiddleware is the middleware to validate JWTs and pass claims to context
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Starting JWT middleware")

		// Get the Authorization header
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		// Extract the token
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}
		tokenStr := authHeader[7:]

		// Parse and verify the JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			// Ensure the signing method is as expected
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)
			if time.Now().After(expTime) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Token has expired",
				})
			}
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token missing expiration claim",
			})
		}

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}
		// Store values in context
		if email, ok := claims["Email"].(string); ok {
			c.Locals(EmailKey, email)
		}
		if userId, ok := claims["id"].(float64); ok {
			c.Locals(UserIDKey, fmt.Sprintf("%.0f", userId))
		}
		if role, ok := claims["role"].(string); ok {
			c.Locals(RoleKey, role)
		}

		log.Println("JWT middleware completed successfully")

		return c.Next()
	}
}
