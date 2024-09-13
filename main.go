package main

import (
	"log"
	"net/http"
	"tes-rssa/database"
	"tes-rssa/ginHandlers"

	_ "tes-rssa/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Crud User
// @version 1
// @Description Rest API CRUD User

// @host localhost:8080

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")

	user := v1.Group("/user")
	{
		user.GET("/", ginHandlers.GetAllUser)
		// user.GET("/")
		// user.GET("/")
		// user.POST("/")
		// user.PUT("/")
		// user.DELETE("/")
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Or specify the domain(s) allowed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Println("Server berjalan di port 8002")
	log.Fatal(http.ListenAndServe(":8002", handler))
}
