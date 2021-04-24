package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/suss/pkg/susslogger"
	"os"
)

type AccessKey struct {
	ApiKey   string
	Password string
}

// AuthReq middleware
func AuthReq() func(*fiber.Ctx) error {
	// Check for API_KEY and API_PASSWORD in the headers section
	// Get it's values from SERVER_API_KEY and SERVER_API_PASSWORD environment variables
	// Authentication will be ignored if the variables not set
	accessKey := AccessKey{
		ApiKey:   os.Getenv("SERVER_API_KEY"),
		Password: os.Getenv("SERVER_API_PASSWORD"),
	}
	
	err := authenticate(accessKey)
	
	return err
	
}

func authenticate(accessKey AccessKey) fiber.Handler {
	unAuthorized := func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	
	// Return new handler
	return func(c *fiber.Ctx) error {
		
		// Get authorization header
		apiKey := c.Get("API_KEY")
		apiPassword := c.Get("API_PASSWORD")
		
		// Check if the header contains correct API_KEY and API_PASSWORD
		if apiKey != accessKey.ApiKey || apiPassword != accessKey.Password {
			susslogger.Log().Error("API_KEY or API_PASSWORD is not correct")
			return unAuthorized(c)
		}
		
		return c.Next()
		
	}
}
