package routes

import (
	"bustrack/controller"

	"github.com/labstack/echo"
)

// VendorRoute multi route
func VendorRoute(e *echo.Echo) (err error) {
	e.GET("/vendor/get/:id", controller.GetVendor)
	e.GET("/vendor/get", controller.GetVendorOrg)
	e.POST("/vendor/add", controller.CreateVender)
	e.PUT("/vendor/update", controller.UpdateVendor)
	e.DELETE("/vendor/delete/:id", controller.DeleteVendor)
	return err
}
