package main

import (
	"context"
	"log"
	"samsamoohooh-go-api/internal/handler"
	"samsamoohooh-go-api/internal/infra/catcher"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/oauth/google"
	"samsamoohooh-go-api/internal/infra/oauth/kakao"
	"samsamoohooh-go-api/internal/infra/token"
	"samsamoohooh-go-api/internal/repository"
	"samsamoohooh-go-api/internal/repository/database"
	"samsamoohooh-go-api/internal/service"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Panicf("failed to connect to database: %v\n", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Panicf("failed to close database connection: %v\n", err)
		}
	}()

	if err := db.AutoMigration(context.Background()); err != nil {
		log.Panicf("failed to auto migrate: %v\n", err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	jwtService := token.NewJWTService(cfg)
	tokenMiddleware := middleware.NewTokenMiddleware(jwtService)
	_ = tokenMiddleware

	oauthGoogleService := google.NewOauthGoogleService(cfg, userService, jwtService)
	oauthKakaoService := kakao.NewOauthKakaoService(cfg, userService, jwtService)

	authHandler := handler.NewAuthHandler(oauthGoogleService, oauthKakaoService, jwtService)

	app := fiber.New(fiber.Config{
		ErrorHandler: catcher.ErrorHandler,
	})

	// Changing TimeZone & TimeFormat
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	v1 := app.Group("/v1")
	{
		api := v1.Group("/api")
		{
			userHandler.Route(api)
			authHandler.Route(api)
		}
	}

	log.Println(app.Listen(":8080"))

}
