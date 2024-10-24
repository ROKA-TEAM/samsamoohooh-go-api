package guard

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/pkg/box"
	"samsamoohooh-go-api/pkg/token"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const (
	TokenKey = "guard_token"
)

type Middleware struct {
	tokenService token.Service
	userService  domain.UserService
}

func NewMiddleware(
	tokenService token.Service,
	userService domain.UserService,
) *Middleware {
	return &Middleware{
		tokenService: tokenService,
		userService:  userService,
	}
}

func (m *Middleware) RequireAuthorization(c fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return box.Wrap(domain.ErrAuthorization, "Missing Authorization header")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	isValid, err := m.tokenService.ValidateToken(tokenString)
	if err != nil {
		return box.Wrap(domain.ErrAuthorization, "Not validation token"+err.Error())
	}

	if !isValid {
		return box.Wrap(domain.ErrAuthorization, "Invalid Authorization header")
	}

	parsedToken, err := m.tokenService.ParseToken(tokenString)
	if err != nil {
		return err
	}

	c.Locals(TokenKey, parsedToken)
	return c.Next()
}

func (m *Middleware) AccessOnly(roleType domain.UserRoleType) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		t, ok := c.Locals(TokenKey).(*token.Token)
		if !ok {
			return box.Wrap(domain.ErrAuthorization, "unable to get token")
		}

		if !(strings.Compare(t.Role, string(roleType)) == 0) {
			return box.Wrap(domain.ErrAuthorization, "invalid role")
		}

		return c.Next()
	}
}

func (m *Middleware) CheckGroupAccess(c fiber.Ctx) error {
	t, ok := c.Locals(TokenKey).(*token.Token)
	if !ok {
		return box.Wrap(domain.ErrAuthorization, "unable to get token")
	}

	gid := fiber.Params[int](c, "gid", -1)
	if gid == -1 {
		var body struct {
			GroupID int `json:"groupID"`
		}

		if err := c.Bind().JSON(&body); err != nil {
			return err
		}

		if body.GroupID < 1 {
			return box.Wrap(domain.ErrAuthorization, "invalid group id")
		}

		gid = body.GroupID
	}

	isValid, err := m.userService.IsUserInGroup(c.Context(), t.Subject, gid)
	if err != nil {
		return err
	}

	if !isValid {
		return box.Wrap(domain.ErrForbidden, "requested user is not in the group")
	}

	return c.Next()
}
