package routes

import (
	"github.com/Woonkly/feed/controllers"
	"github.com/labstack/echo/v4"
	
)

func Task(e *echo.Echo) {

	e.POST( "/add", controllers.Ctltask.Add )
	e.GET( "/show", controllers.Ctltask.Show )
	e.DELETE( "/del", controllers.Ctltask.Del )
	e.PUT( "/update", controllers.Ctltask.Update )

}

