package controller

import (
	"bustrack/dbs"
	"bustrack/models"
	"bustrack/tools"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq" //for postgres driver
)

// CreateDriver create driver
func CreateDriver(c echo.Context) (err error) {
	dr := &models.Driver{
		Email:    c.FormValue("email"),
		Vendorid: stringtoInt(c.FormValue("venid")),
		Name:     c.FormValue("name"),
		Address:  c.FormValue("address"),
		Password: c.FormValue("password"),
		Contact:  c.FormValue("contact"),
	}
	if err = c.Bind(dr); err != nil {
		return err
	}
	// db := dbs.GetDB()
	result, err := models.InsertDriver(dbs.DB, dr)
	tools.PanicIf(err)
	fmt.Println(result)
	// dbs.CloseDB(db)
	return c.JSON(http.StatusCreated, dr)
}

//DeleteDriver delete the driver by id
func DeleteDriver(c echo.Context) (err error) {
	// db := dbs.GetDB()
	result, err := models.DeleteDriver(dbs.DB, stringtoInt(c.Param("id")))
	tools.PanicIf(err)
	fmt.Println(result)
	//dbs.CloseDB(db)
	return c.NoContent(http.StatusNoContent)
}

//GetDriver get the driver by id
func GetDriver(c echo.Context) (err error) {
	//	db := dbs.GetDB()
	result, err := models.GetDriver(dbs.DB, stringtoInt(c.Param("id")))
	tools.PanicIf(err)
	drs := make([]models.Driver, 0.0)
	dr := models.Driver{}
	for result.Next() {
		result.Scan(&dr.Driverid, &dr.Email, &dr.Vendorid, &dr.Name, &dr.Address, &dr.Password, &dr.Contact)
		fmt.Println(dr.Driverid, dr.Email, dr.Vendorid, dr.Name, dr.Address, dr.Password, dr.Contact)
		drs = append(drs, dr)
	}
	//	dbs.CloseDB(db)
	return c.JSON(http.StatusOK, drs)
}

//GetDriverbyVendor get driver by vendor
func GetDriverbyVendor(c echo.Context) (err error) {
	fmt.Println("drid:", c.QueryParam("drid"))
	fmt.Println("venid:", c.QueryParam("venid"))
	result, err := models.GetDriverbyVen(dbs.DB, stringtoInt(c.QueryParam("drid")), stringtoInt(c.QueryParam("venid")))
	tools.PanicIf(err)
	fmt.Println("result:", result, "result end")
	dr := models.Driver{}
	for result.Next() {
		result.Scan(&dr.Driverid, &dr.Email, &dr.Vendorid, &dr.Name, &dr.Address, &dr.Password, &dr.Contact)
		fmt.Println(dr.Driverid, dr.Email, dr.Vendorid, dr.Name, dr.Address, dr.Password, dr.Contact)
	}
	//	dbs.CloseDB(db)
	return c.JSON(http.StatusOK, dr)
}

//UpdateDrivers update the vendor by id
func UpdateDrivers(c echo.Context) (err error) {
	fmt.Println(c.FormValue("email"))
	fmt.Println(c.FormValue("name"))
	fmt.Println(c.FormValue("address"))
	fmt.Println(c.FormValue("password"))
	fmt.Println(c.FormValue("contact"))
	fmt.Println(c.FormValue("venid"))
	orm := map[string]string{
		"email":     c.FormValue("email"),
		"name":      c.FormValue("name"),
		"address":   c.FormValue("address"),
		"password":  c.FormValue("password"),
		"contact":   c.FormValue("contact"),
		"vendor_id": c.FormValue("venid"),
	}
	query := tools.UpdateBuilder("driver", orm, "driver_id", c.FormValue("drid"))
	//db := dbs.GetDB()
	result, err := models.UpdateDriver(dbs.DB, query)
	if err != nil {
		fmt.Println("Error is", err)
	}
	//tools.PanicIf(err)
	//dbs.CloseDB(db)
	fmt.Println("result:", result)
	return c.JSON(http.StatusOK, "updated")
}
