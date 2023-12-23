package creator

import (
	"chain-responsability/database"
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/platform/repository"
	"chain-responsability/model"
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
		return err
	}

	defer db.Close()

	usersRepository := repository.NewUsersRepository(db)

	user := model.User{
		Name:     process.Name,
		LastName: process.LastName,
		Genre:    process.Genre,
		Age:      process.Age,
		Email:    process.Email,
	}

	id, err := usersRepository.Create(user)
	if err != nil {
		return err
	}

	process.ID = id

	if s.Next != nil {
		return s.Next.StartProcess(process)
	}

	return nil
}

func (s *Service) SetNext(handler ports.Handler) ports.Handler {
	s.Next = handler
	return s.Next
}
