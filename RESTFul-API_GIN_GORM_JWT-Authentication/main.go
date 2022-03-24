package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"gin_gorm_jwt/config"
	"gin_gorm_jwt/controller"
	"gin_gorm_jwt/middleware"
	"gin_gorm_jwt/repository"
	"gin_gorm_jwt/service"
)

var (
	db               *gorm.DB                    = config.SetupDBConnection()
	autherRepository repository.AutherRepository = repository.NewAutherRepository(db)
	jwtService       service.JWTService          = service.NewJWTService()
	autherService    service.AutherService       = service.NewAutherService(autherRepository)
	authService      service.AuthService         = service.NewAuthService(autherRepository)
	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
	autherController controller.AutherController = controller.NewAutherController(autherService, jwtService)
)

func main() {
	defer config.CloseDBConnection(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Root route...!",
		})
	})

	// authRoutes := r.Group("api/auth", middleware.AuthorizeJWT(jwtService))
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	autherRoutes := r.Group("api/auther", middleware.AuthorizeJWT(jwtService))
	{
		autherRoutes.GET("/profile", autherController.Profile)
		autherRoutes.PUT("/profile", autherController.Update)
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env")
	}

	PORT := os.Getenv("PORT")

	r.Run(PORT)
}
