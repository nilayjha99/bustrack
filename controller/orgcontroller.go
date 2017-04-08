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

// CreateOrganization create organization
func CreateOrganization(c echo.Context) (err error) {
	org := &models.Organization{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Address:  c.FormValue("address"),
		Password: c.FormValue("password"),
		Contact:  c.FormValue("contact"),
	}
	if err = c.Bind(org); err != nil {
		return err
	}
	// db := dbs.GetDB()
	result, err := models.InsertOrg(dbs.DB, org)
	tools.PanicIf(err)
	fmt.Println(result)
	// dbs.CloseDB(db)
	return c.JSON(http.StatusCreated, org)
}

//DeleteOraganization delete the organization by id
func DeleteOraganization(c echo.Context) (err error) {
	// db := dbs.GetDB()
	result, err := models.DeleteOrg(dbs.DB, stringtoInt(c.Param("id")))
	tools.PanicIf(err)
	fmt.Println(result)
	//dbs.CloseDB(db)
	return c.NoContent(http.StatusNoContent)
}

//GetOrganization get the organization by id
func GetOrganization(c echo.Context) (err error) {
	//	db := dbs.GetDB()
	result, err := models.GetOrg(dbs.DB, stringtoInt(c.Param("id")))
	tools.PanicIf(err)
	orgs := make([]models.Organization, 0.0)
	org := models.Organization{}
	for result.Next() {
		result.Scan(&org.Orgid, &org.Email, &org.Name, &org.Address, &org.Password, &org.Contact)
		fmt.Println(org.Orgid, org.Email, org.Name, org.Address, org.Password, org.Contact)
		orgs = append(orgs, org)
	}
	//	dbs.CloseDB(db)
	return c.JSON(http.StatusOK, orgs)
}

//UpdateOraganization update the organization by id
func UpdateOraganization(c echo.Context) (err error) {
	org := &models.Organization{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Address:  c.FormValue("address"),
		Password: c.FormValue("password"),
		Contact:  c.FormValue("contact"),
	}
	orm := map[string]string{
		"email":    org.Email,    //"nilayjha@gmail.com",
		"name":     org.Name,     //"nilayjha",
		"address":  org.Address,  //"mehsana",
		"password": org.Password, // "nilay",
		"contact":  org.Contact,  //"\"office\"=>\"7874969553\"",
	}

	query := tools.UpdateBuilder("organization", orm, "organization_id", c.Param("orgid"))
	//db := dbs.GetDB()
	result, err := models.UpdateOrg(dbs.DB, query)
	if err != nil {
		fmt.Println("Error is", err.Error())
	}
	//tools.PanicIf(err)
	//dbs.CloseDB(db)
	fmt.Println("result:", result)
	return c.JSON(http.StatusOK, "updated")
}
