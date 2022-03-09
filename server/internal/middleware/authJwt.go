package middleware

import (
	"atro/internal/handler"
	"atro/internal/helper"
	"atro/internal/repository"
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

func IsAdmin() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		userID, isExist := ctx.Get("userID")
		fmt.Println("userID: ", userID)
		if isExist == true {
			checkUser, err := repository.NewUserRepository().GetUser(int(userID.(float64)))
			if err == nil {
				role, err := repository.NewRoleRepository().GetRole(checkUser.UserRoleID)
				if err != nil {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find ROLE", err.Error()))
				} else {
					if role.RoleName == "ADMIN" {
						fmt.Println("Admin role session is running!")
					} else {
						ctx.AbortWithStatusJSON(http.StatusForbidden, helper.BuildResponse(-1, "require ADMIN role", ""))
					}
				}
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BuildResponse(-1, "Error when find USER", err.Error()))
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildResponse(-1, "Not Exist session", ""))
		}
	}
}
