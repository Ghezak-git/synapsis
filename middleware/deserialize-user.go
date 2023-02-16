package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ghezak-git/synapsis/initializers"
	"github.com/Ghezak-git/synapsis/models"
	"github.com/Ghezak-git/synapsis/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "essage": "You are not logged in"})
			return
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "essage": err.Error()})
			return
		}

		var user models.User
		result := initializers.DB.First(&user, "id_user = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "essage": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()

	}
}
