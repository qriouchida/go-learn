package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"king-tool-api/handler"
	"king-tool-api/infra"
	"king-tool-api/service"
	"log"
	"net/http"
)

func main() {

	engine := infra.DBInit()
	factory := service.NewService(engine)
	infra.DBMigrate(engine)

	defer func() {
		log.Println("engine close")
		engine.Close()
	}()

	g := gin.Default()
	g.Use(service.ServiceFactoryMiddleware(factory))

	routes := g.Group("/v1")
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}
	// route Echo用
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	g.Run()
}

func Env_load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
