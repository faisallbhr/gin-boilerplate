package main

import (
	"github.com/faisallbhr/gin-boilerplate/config"
	"github.com/faisallbhr/gin-boilerplate/database"
	"github.com/faisallbhr/gin-boilerplate/routes"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
