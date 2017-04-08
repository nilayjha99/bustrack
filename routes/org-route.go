package routes

import (
	"bustrack/controller"

	"github.com/labstack/echo"
)

// OrgRoute multi route
func OrgRoute(e *echo.Echo) (err error) {
	e.GET("/organization/get/:id", controller.GetOrganization)
	e.PUT("/organization/update/:orgid", controller.UpdateOraganization)
	e.POST("/organization/add", controller.CreateOrganization)
	e.DELETE("/organization/delete/:id", controller.DeleteOraganization)
	return err
}
