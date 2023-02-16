package routes

import (
	"github.com/Ghezak-git/synapsis/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/login", rc.authController.Login)
	router.GET("/token", rc.authController.RefreshToken)
	router.GET("/logout", rc.authController.LogoutUser)
	router.POST("/signup", rc.authController.SignUp)

}
