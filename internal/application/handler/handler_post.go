package handler

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
)

type PostHandler struct {
	postService domain2.PostService
}

//
//func NewPostHandler(postService domain2.PostService) *PostHandler {
//	return &PostHandler{postService: postService}
//}
//
//func (h *PostHandler) Route(router fiber.Router, guard *middleware.GuardMiddleware) {
//	router.Post("/", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Create)
//	router.Get("/", guard.RequireAccess(domain2.UserRoleAdmin), h.List)
//	router.Get("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetByID)
//	router.Put("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Update)
//	router.Delete("/:id", guard.RequireAccess(domain2.UserRoleAdmin), h.Delete)
//}
//
//func (h *PostHandler) Create(c *fiber.Ctx) error {
//	body := new(presenter.PostCreateRequest)
//	if err := utils.ParseAndVerify(c, body); err != nil {
//		return err
//	}
//
//	createdPost, err := h.postService.Create(c.Context(), body.GroupID, body.ToDomain())
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusCreated).JSON(presenter.NewPostCreateResponse(createdPost))
//}
//
//func (h *PostHandler) List(c *fiber.Ctx) error {
//	offset := c.QueryInt("offset", DefaultOffset)
//	limit := c.QueryInt("limit", DefaultLimit)
//
//	listPost, err := h.postService.List(c.Context(), offset, limit)
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewPostListResponse(listPost))
//}
//
//func (h *PostHandler) GetByID(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	gotPost, err := h.postService.GetByID(c.Context(), id)
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewPostGetByIDResponse(gotPost))
//}
//
//func (h *PostHandler) GetCommentsByID(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	offset := c.QueryInt("offset", DefaultOffset)
//	limit := c.QueryInt("limit", DefaultLimit)
//
//	listComment, err := h.postService.GetCommentsByID(c.Context(), id, offset, limit)
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewPostGetCommentsByIDResponse(listComment))
//}
//
//func (h *PostHandler) Update(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	body := new(presenter.PostUpdateRequest)
//	if err := utils.ParseAndVerify(c, body); err != nil {
//		return err
//	}
//
//	updatedPost, err := h.postService.Update(c.Context(), id, body.ToDomain())
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewPostUpdateResponse(updatedPost))
//}
//
//func (h *PostHandler) Delete(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	err = h.postService.Delete(c.Context(), id)
//	if err != nil {
//		return err
//	}
//
//	return c.SendStatus(fiber.StatusNoContent)
//}
