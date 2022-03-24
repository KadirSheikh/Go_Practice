package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"gin_gorm_jwt/config"
	"gin_gorm_jwt/controller"
)

var (
	db             *gorm.DB                  = config.SetupDBConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDBConnection(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Root route...!",
		})
	})

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env")
	}

	PORT := os.Getenv("PORT")

	r.Run(PORT)
}
