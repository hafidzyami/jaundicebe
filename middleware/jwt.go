package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafidzyami/jaundicebe/config"

	jwtware "github.com/gofiber/contrib/jwt"
)
   
   func JWTMiddleware(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
	 SigningKey: jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET"))},
	 ContextKey: "jwt",
	 ErrorHandler: func(c *fiber.Ctx, err error) error {
	  // Return status 401 and failed authentication error.
	  return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	   "error": true,
	   "msg":   err.Error(),
	  })
	 },
	})(c)
   }