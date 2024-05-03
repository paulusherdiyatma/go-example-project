// package middleware represents a middleware layer of the application
package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	// record the start time of the request
	start := time.Now()

	// proceed with the request
	if err := c.Next(); err != nil {
		return err
	}

	// record the end time of the request
	end := time.Now()

	// log the request
	log.Printf("[%s] %s %s - %v", c.Method(), c.Path(), c.IP(), end.Sub(start))
	return nil
}
