package catch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/domain"
)

var ErrorHandler = func(c *fiber.Ctx, caughtErr error) error {
	var status = fiber.StatusInternalServerError

	switch {
	case errors.Is(domain.ErrTokenGenerate, caughtErr):
		status = fiber.StatusInternalServerError

	case errors.Is(domain.ErrTokenParse, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrInvalidTokenIssuer, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrTokenNotActiveYet, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrTokenExpired, caughtErr):
		status = fiber.StatusUnauthorized

	default:
		return fiber.DefaultErrorHandler(c, caughtErr)
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(status).SendString(caughtErr.Error())
}