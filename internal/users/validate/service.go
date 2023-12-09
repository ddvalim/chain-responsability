package validate

import (
	"chain-responsability/database"
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/platform/repository"
	"errors"
	"log"
)

type Service struct {
	Next ports.Handler
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartProcess(process *ports.Process) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	usersRepository := repository.NewUsersRepository(db)

	user, err := usersRepository.GetUser(process.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already exists")
	}

	if s.Next != nil {
		return s.Next.StartProcess(process)
	}

	return nil
}

func (s *Service) SetNext(handler ports.Handler) ports.Handler {
	s.Next = handler
	return s.Next
}
