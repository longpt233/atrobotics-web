package middleware

import (
	"atro/internal/handler"
	"atro/internal/helper"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT -> to authorize JWT Token
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "

		// check header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "header đâu người ae ? ", ""))

		}

		// validate token
		tokenString := authHeader[len(BearerSchema):]
		if token, err := handler.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Not Valid Token", err.Error()))

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Claims fail", ""))

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
				} else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Invalid token", ""))
				}

			}
		}

	}

}
