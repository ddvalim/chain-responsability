package controller

import (
	"chain-responsability/controller/response"
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/users"
	"chain-responsability/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type UserController struct {
	UserService users.Service
}

func NewUserController() UserController {
	userService := users.NewService()

	return UserController{UserService: userService}
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var user model.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = user.ValidateUser(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	process := ports.Process{
		Name:     user.Name,
		LastName: user.LastName,
		Genre:    user.Genre,
		Email:    user.Email,
		Age:      user.Age,
	}

	err = u.UserService.Init(process, http.MethodPost)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.Write(w, http.StatusCreated, response.Response{
		Message:    "successfully created user",
		StatusCode: http.StatusCreated,
	})
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if len(email) == 0 {
		response.Error(w, http.StatusBadRequest, errors.New("invalid email query parameter"))
		return
	}

	process := ports.Process{
		Email: email,
	}

	err := u.UserService.Init(process, http.MethodDelete)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.Write(w, http.StatusNoContent, nil)
	return
}
