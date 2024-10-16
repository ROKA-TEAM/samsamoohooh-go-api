package handler

import (
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler/utils"
	"samsamoohooh-go-api/internal/infra/presenter"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Route(router fiber.Router) {
	router.Post("/users", h.Create)
	router.Get("/users", h.List)
	router.Get("/users/:id", h.GetByID)
	router.Get("/users/:id/groups", h.GetGroupsByID)
	router.Put("/users/:id", h.Update)
	router.Delete("/users/:id", h.Delete)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	body := new(presenter.UserCreateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdUser, err := h.userService.Create(c.Context(), body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewUserCreateResponse(createdUser))
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	limit := c.QueryInt("limit")
	offset := c.QueryInt("offset")

	listUsers, err := h.userService.List(c.Context(), limit, offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewListUserResponse(listUsers))
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt(":id")
	if err != nil {
		return err
	}

	gotUser, err := h.userService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetByIDResponse(gotUser))
}

func (h *UserHandler) GetGroupsByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt(":id")
	if err != nil {
		return err
	}

	gotGroups, err := h.userService.GetGroupsByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetGroupsByIDResponse(gotGroups))
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt(":id")
	if err != nil {
		return err
	}

	body := new(presenter.UserUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	udpatedUser, err := h.userService.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(presenter.NewUserUpdateResponse(udpatedUser))
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt(":id")
	if err != nil {
		return err
	}

	err = h.userService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}