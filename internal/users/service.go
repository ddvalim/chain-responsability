package users

import "chain-responsability/internal/core/ports"

type Service interface {
	Init(p ports.Process) error
}

type ServiceImpl struct{}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (s ServiceImpl) Init(p ports.Process) error {
	chain := Default()

	err := chain.StartProcess(&p)
	if err != nil {
		return err
	}

	return nil
}
