package users

import (
	"chain-responsability/database"
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/users/creator"
	"chain-responsability/internal/users/deleter"
	"chain-responsability/internal/users/getter"
	"chain-responsability/internal/users/validate"
	"log"
	"net/http"
)

func Create() ports.Handler {
	validateSvc := validate.NewService(http.MethodPost)

	createSvc := creator.NewService()

	validateSvc.SetNext(createSvc)

	return validateSvc
}

func Delete() ports.Handler {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	validateSvc := validate.NewService(http.MethodDelete)

	deleteSvc := deleter.NewService(db)

	validateSvc.SetNext(deleteSvc)

	return validateSvc
}

func Get() ports.Handler {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	getSvc := getter.NewService(db)

	return getSvc
}
