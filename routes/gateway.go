package routes

import (
	"github.com/labstack/echo/v4"
)

func ApiGateway(api *echo.Echo) {

	Routes := api.Group("/v1")
	RoutesManagment := Routes.Group("/managment")
	{
		//	RoutesManagment.GET("/tickets", controllers.TicketManagmentController)
	}

}
