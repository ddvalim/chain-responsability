package getter

import (
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/platform/repository"
	"database/sql"
	"errors"
)

type Service struct {
	next ports.Handler
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

	res, err := usersRepository.Get(process.Email)
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("no user found")
	}

	process.User = res

	if s.next != nil {
		return s.next.StartProcess(process)
	}

	return nil
}

func (s *Service) SetNext(handler ports.Handler) ports.Handler {
	s.next = handler
	return handler
}
