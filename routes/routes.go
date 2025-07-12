package routes

import (
	"net/http"
	"strings"

	"github.com/faisallbhr/gin-boilerplate/config"
	"github.com/faisallbhr/gin-boilerplate/controllers"
	"github.com/faisallbhr/gin-boilerplate/helpers"
	"github.com/faisallbhr/gin-boilerplate/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	corsOrigins := strings.Split(config.GetEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"), ", ")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/refresh", controllers.RefreshToken)

		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			users := auth.Group("/users")
			{
				users.GET("/me", controllers.Me)
				users.GET("/", controllers.GetUsers)
				users.GET("/:id", controllers.GetUser)
				users.PATCH("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
				users.PATCH("/:id/password", controllers.UpdatePassword)
			}
		}
	}

	router.NoRoute(func(c *gin.Context) {
		helpers.ResponseError(c, "Route not found", http.StatusNotFound, nil)
	})

	return router
}
