package services

import (
	"ASCorpImportantDates/internal/models"
	"ASCorpImportantDates/internal/storage/sqlite"
	"fmt"
)

func CreateUser(user models.User, storage *sqlite.Storage) error {
	userId, err := storage.CreateUser(user)
	if err != nil {
		return err
	}
	fmt.Println(userId)
	return nil
}

func ReadUser(login string, storage *sqlite.Storage) (models.User, error) {
	user, err := storage.ReadUser(login)
	if err != nil {
		return user, err
	}
	return user, nil
}
