package users

import (
	"chain-responsability/internal/core/ports"
	"chain-responsability/internal/users/create"
	"chain-responsability/internal/users/validate"
)

func Default() ports.Handler {
	validateSvc := validate.NewService()

	createSvc := create.NewService()

	validateSvc.SetNext(createSvc)

	return validateSvc
}
