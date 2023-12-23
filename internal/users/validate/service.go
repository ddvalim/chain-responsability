package validate

import (
	"chain-responsability/database"
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/platform/repository"
	"errors"
	"log"
	"net/http"
)

type Service struct {
	Next   ports.Handler
	method string
}

func NewService(method string) *Service {
	return &Service{
		method: method,
	}
}

func (s *Service) StartProcess(process *ports.Process) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	usersRepository := repository.NewUsersRepository(db)

	user, err := usersRepository.Get(process.Email)
	if err != nil {
		return err
	}

	if user != nil {
		if s.method == http.MethodPost {
			return errors.New("user already exists")
		}
	} else {
		if s.method == http.MethodDelete {
			return errors.New("user not found")
		}
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
