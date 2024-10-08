package router

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/adapter/presentation/router/validator"
	"samsamoohooh-go-api/internal/infra/config"
)

const DefaultStartPort = ":8080"

type Router struct {
	*fiber.App
	config        *config.Config
	handlerSet    HandlerSet
	middlewareSet MiddlewareSet
}

func New(config *config.Config, handlerSet HandlerSet, middlewareSet MiddlewareSet) *Router {

	r := &Router{config: config, handlerSet: handlerSet, middlewareSet: middlewareSet, App: fiber.New(fiber.Config{
		AppName:         config.HTTP.Name,
		StructValidator: validator.New(),
		ErrorHandler:    customErrorHandler,
	})}

	r.setMiddleware()
	r.route()

	return r
}

// set middleware
func (r *Router) setMiddleware() {

}

// Route
func (r *Router) route() {
	// users
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{

				google := auth.Group("/google")
				{
					google.Get("/", r.handlerSet.AuthHandler.GoogleLogin)
					google.Get("/callback", r.handlerSet.AuthHandler.GoogleCallback)
				}
			}

			users := v1.Group("/users", r.middlewareSet.GuardMiddleware.ProtectFromTempToken)
			{
				users.Post("/", r.handlerSet.UserHandler.Create)
				users.Get("/:id", r.handlerSet.UserHandler.GetByID)
				users.Get("/:id/groups", r.handlerSet.UserHandler.GetGroupsByID)
				users.Get("/", r.handlerSet.UserHandler.GetAll)
				users.Put("/:id", r.handlerSet.UserHandler.Update)
				users.Delete("/:id", r.handlerSet.UserHandler.Delete)
			}

			groups := v1.Group("/groups")
			{
				groups.Post("/", r.handlerSet.GroupHandler.Create)
				groups.Get("/:id", r.handlerSet.GroupHandler.GetByID)
				groups.Get("/:id/users", r.handlerSet.GroupHandler.GetUsersByID)
			}
		}
	}
}

func (r *Router) Start() error {
	startPort := DefaultStartPort
	if r.config.HTTP.Port != "" {
		startPort = r.config.HTTP.Port
	}

	return r.Listen(startPort)
}
