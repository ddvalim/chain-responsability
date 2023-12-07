package model

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"strconv"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Genre     string    `json:"genre,omitempty"`
	Age       string    `json:"age,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u User) ValidateUser() error {
	fields := []string{u.Name, u.LastName, u.Genre, u.Age, u.Email}

	for _, f := range fields {
		if len(f) == 0 {
			return errors.New(fmt.Sprintf("field %s can not be empty", f))
		}
	}

	if err := u.ValidateEmail(); err != nil {
		return errors.New(fmt.Sprintf("invalid email: %s", u.Email))
	}

	if err := u.ValidateAge(); err != nil {
		return err
	}

	return nil
}

func (u User) ValidateEmail() error {
	return checkmail.ValidateFormat(u.Email)
}

func (u User) ValidateAge() error {
	intAge, err := strconv.ParseInt(u.Age, 10, 64)
	if err != nil {
		return errors.New("invalid age entry")
	}

	if intAge < 18 {
		return errors.New("user age is lower than 18")
	}

	return nil
}
