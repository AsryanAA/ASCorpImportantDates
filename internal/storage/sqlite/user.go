package sqlite

import (
	"ASCorpImportantDates/internal/models"
	"ASCorpImportantDates/internal/storage"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
)

func (s *Storage) CreateUser(user models.User) (int64, error) {
	const op = "storage.sqlite.CreateUser"

	stmt, err := s.db.Prepare(`INSERT INTO users(login, surname, name, dob, reg_date) 
									 VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(user.Login, user.Surname, user.Name, user.DOB, user.RegDate)
	if err != nil {
		// TODO: refactoring this
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, err)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id %w", op, err)
	}

	return id, nil
}

func (s *Storage) ReadUser(login string) (models.User, error) {
	var user models.User
	const op = "storage.sqlite.GetWorkPlaces"

	stmt, err := s.db.Prepare(`SELECT login, surname, name, dob, reg_date FROM users WHERE login = ?`)
	if err != nil {
		return user, fmt.Errorf("%s: prepare statement %w", op, err)
	}

	rows := stmt.QueryRow(login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, storage.ErrUserNotFound
		}

		return user, fmt.Errorf("%s: execute statement %w", op, err)
	}

	err = rows.Scan(&user.Login, &user.Surname, &user.Name, &user.DOB, &user.RegDate)
	if err != nil {
		fmt.Printf("%s: error receive record %w", op, err)
	}

	return user, nil
}
