package users

import (
	"chain-responsability/internal/core/ports"
	"net/http"
)

type Service interface {
	Init(p ports.Process, method string) error
}

type ServiceImpl struct{}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func getChain(method string) ports.Handler {
	switch method {
	case http.MethodPost:
		return Create()
	case http.MethodDelete:
		return Delete()
	}

	// TODO: Fix me!
	return Create()
}

func (s ServiceImpl) Init(p ports.Process, method string) error {
	chain := getChain(method)

	err := chain.StartProcess(&p)
	if err != nil {
		return err
	}

	return nil
}
