package deleter

import (
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/platform/repository"
	"database/sql"
)

type Service struct {
	Next ports.Handler
	DB   *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) StartProcess(process *ports.Process) error {
	defer s.DB.Close()

	usersRepository := repository.NewUsersRepository(s.DB)

	err := usersRepository.Delete(process.Email)
	if err != nil {
		return err
	}

	if s.Next != nil {
		return s.Next.StartProcess(process)
	}

	return nil
}

func (s *Service) SetNext(handler ports.Handler) ports.Handler {
	s.Next = handler
	return handler
}
