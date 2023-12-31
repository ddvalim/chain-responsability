package routes

import (
	"chain-responsability/controller"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:    "/user",
		Method: http.MethodPost,
		Func:   controller.NewUserController().Create,
	},
	{
		URI:    "/user",
		Method: http.MethodDelete,
		Func:   controller.NewUserController().Delete,
	},
	{
		URI:    "/user",
		Method: http.MethodGet,
		Func:   controller.NewUserController().Get,
	},
}
