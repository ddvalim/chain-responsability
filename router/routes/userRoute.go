package routes

import (
	"chain-responsability/controller"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:    "/user",
		Method: http.MethodPost,
		Func:   controller.NewUserController().CreateUser,
	},
}
