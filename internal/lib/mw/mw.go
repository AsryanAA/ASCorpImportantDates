package mw

import (
	"ASCorpImportantDates/internal/lib/jwt"
	"ASCorpImportantDates/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHeader(login string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": "Вы не авторизованы",
			})
			return
		}

		u := models.User{
			Login:   login,
			Surname: "Asryan",
			Name:    "",
		}
		checkToken, err := jwt.GenerateJWT(u)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": "Не удалось распознать токен",
			})
			return
		}

		// TODO доделать функционал
		jwt.CheckValidJWT(token)

		// TODO узнать до какого символа должно быть совпадение
		if token[:100] != checkToken[:100] {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": "Токены не совпадают, повторите авторизацию",
			})
			return
		}

		ctx.Next()
	}
}
