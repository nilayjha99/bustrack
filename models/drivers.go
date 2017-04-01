package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Driver struct for crud
type Driver struct {
	Driverid int
	Email    string
	Vendorid int
	Name     string
	Address  string
	Password string
	Contact  string
}

//to execute normal insert queries
func insertDriver(db *sql.DB, dr *Driver) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into driver(email, vendor_id,name ,address,password, contact) values(%s,%d,%s,%s,%s,%s)",
			dr.Email,
			dr.Vendorid,
			dr.Name,
			dr.Address,
			dr.Password,
			dr.Contact,
		),
	)
}

//to delete selected entry
func deleteDriver(db *sql.DB, drid int) (sql.Result, error) {
	if drid == -1 {
		return db.Exec(fmt.Sprintf("truncate table driver"))
	}

	return db.Exec(fmt.Sprintf("delete from driver where driver_id=%d", drid))
}

//to get selected entry
func getDriver(db *sql.DB, drid int) (*sql.Rows, error) {
	if drid == -1 {
		return db.Query(fmt.Sprintf("select *from driver"))
	}
	return db.Query(fmt.Sprintf("select *from driver where driver_id=%d", drid))
}

//get driver by vendor_id
func getDriverbyVen(db *sql.DB, venid, drid int) (*sql.Rows, error) {
	if venid == -1 {
		return db.Query(fmt.Sprintf("select *from driver where vendor_id=%d", venid))
	}
	return db.Query(fmt.Sprintf("select *from driver where driver_id=%d AND vendor_id=%d", drid, venid))
}

//to update selected entry
func updateDriver(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
