package handlers

import (
	"ASCorpImportantDates/internal/http_server/services"
	"ASCorpImportantDates/internal/models"
	"ASCorpImportantDates/internal/storage/sqlite"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Создание новой записи Пользователя
// @Desription Создает новую запись в таблице
// @Tags Пользователь (users)
// @Success 200 {object} models.User
// @Failure 404
// @Router /create [post]
func CreateUser(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUser models.User
		if err := ctx.BindJSON(&newUser); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Некорректные данные",
			})
			return
		}

		err := services.CreateUser(newUser, storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Создание не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusCreated, newUser)
			return
		}
	}
}

// @Summary Получение пользователя по логину
// @Desription Читает запись из таблицы
// @Tags Пользователь (users)
// @Success 200 {object} models.User
// @Failure 404
// @Router /read [get]
func ReadUser(login string, storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := services.ReadUser(login, storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Чтение не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
}
