package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT Authentication --")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		fmt.Println("token not present in the header")
		return fmt.Errorf("unauthorized")
	}

	fmt.Println("token: ", token[0])

	claims, err := validateToken(token[0])
	if err != nil {
		return err
	}
	expiresFloat := claims["expires"].(float64)
	expires := int64(expiresFloat)

	// Check token expiration
	if time.Now().Unix() > (expires) {
		fmt.Println("token expired")
		return fmt.Errorf("unauthenticated")
	}
	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token: ", err)
		return nil, fmt.Errorf("not authenticated")
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("not authenticated")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not authenticated")
	}

	return claims, nil
}
