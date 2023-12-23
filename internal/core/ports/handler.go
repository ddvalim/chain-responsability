package ports

import "chain-responsability/model"

type Process struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	LastName string      `json:"last_name"`
	Genre    string      `json:"genre"`
	Email    string      `json:"email"`
	Age      string      `json:"age"`
	User     *model.User `json:"user"`
}

type Handler interface {
	StartProcess(process *Process) error
	SetNext(handler Handler) Handler
}

type Response struct {
	UserID string
}
