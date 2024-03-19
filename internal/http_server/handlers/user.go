package handlers

import (
	"ASCorpImportantDates/internal/http_server/services"
	"ASCorpImportantDates/internal/lib/jwt"
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
// @Router /users/create [post]
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
// @Router /users/read [get]
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

// @Summary Получение пользователей
// @Desription Читает записи из таблицы
// @Tags Пользователи (users)
// @Success 200 {object} []models.User
// @Failure 404
// @Router /users/all [get]
func ReadUsers(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := services.ReadUsers(storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Чтение не произошло",
			})
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, users)
			return
		}
	}
}

// SignIn @Summary Возвращает пользователя
// @Description Возвращает данные пользователя
// @Tags User (auth)
// @Param login body string true "Login"
// @Success 200 {object} models.User
// @Failure 404 {object}
// @Router /auth [post]
func SignIn(storage *sqlite.Storage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u models.UserAuth
		if err := ctx.BindJSON(&u); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err,
				"error":        "Некорректные данные",
			})
			return
		}
		user, err := services.ReadUser(u.Login, storage)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": "Пользователь с таким логином не найден",
				"error":        err,
			})
			return
		}
		token, err := jwt.GenerateJWT(user)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"errorMessage": "Не удалось создать jwt токен",
				"error":        err,
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	}
}
