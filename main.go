package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMX0.J3geSMCoxASJJzSYBZwGxfAlAhmVPdIeAx6_17o3HhA")
	if err != nil {
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}

	UserHandler := handler.NewUSerHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", UserHandler.RegisterUser)
	api.POST("/sessions", UserHandler.Login)
	api.POST("/email_checkers", UserHandler.CheckEmailAvailability)
	api.POST("/avatars", UserHandler.UploadAvatar)

	router.Run()

}
