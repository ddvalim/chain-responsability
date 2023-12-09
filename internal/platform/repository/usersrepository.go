package repository

import (
	"chain-responsability/model"
	"database/sql"
	"strconv"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur UsersRepository) Create(user model.User) (string, error) {
	stmt, err := ur.db.Prepare(
		"INSERT INTO users (name, lastname, genre, age, email) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.LastName, user.Genre, user.Age, user.Email)
	if err != nil {
		return "", err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(userID, 10), nil
}

func (ur UsersRepository) GetUser(email string) (*model.User, error) {
	users, err := ur.db.Query(
		"SELECT id, name, lastname, email, password, createdAt FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return nil, err
	}

	defer users.Close()

	var user model.User

	if users.Next() {
		if err := users.Scan(&user.ID, &user.Name, &user.LastName, &user.Genre, &user.Age, &user.Email,
			&user.CreatedAt); err != nil {
			return nil, err
		}
	}

	return &user, nil
}
