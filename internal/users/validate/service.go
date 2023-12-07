package validate

import "chain-responsability/internal/core/ports"

type Service struct {
	Next ports.Handler
}

func NewService() *Service {
	return &Service{}
}

func (s Service) StartProcess(process *ports.Process) error {
	// search for user in database

	if s.Next != nil {
		return s.Next.StartProcess(process)
	}

	return nil
}

func (s Service) SetNext(handler ports.Handler) ports.Handler {
	s.Next = handler
	return s.Next
}
