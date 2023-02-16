package main

import (
	"log"
	"net/http"

	"github.com/Ghezak-git/synapsis/controllers"
	"github.com/Ghezak-git/synapsis/initializers"
	"github.com/Ghezak-git/synapsis/middleware"
	"github.com/Ghezak-git/synapsis/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	ProductController      controllers.ProductController
	ProductRouteController routes.ProductRouteController

	TransactionController      controllers.TransactionController
	TransactionRouteController routes.TransactionRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables1", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	ProductController = controllers.NewProductController(initializers.DB)
	ProductRouteController = routes.NewProductRouteController(ProductController)

	TransactionController = controllers.NewTransactionController(initializers.DB)
	TransactionRouteController = routes.NewTransactionRouteController(TransactionController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	server.Use(cors.New(corsConfig))
	server.GET("/", func(ctx *gin.Context) {
		message := "Hiii!!! Your place not here!!!!"
		ctx.JSON(http.StatusOK, gin.H{"message": message})
	})
	router := server.Group("/api")
	router.GET("/", middleware.DeserializeUser(), func(ctx *gin.Context) {
		message := "Welcome to My GHZ API"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	ProductRouteController.ProductRoute(router)
	TransactionRouteController.TransactionRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

}
