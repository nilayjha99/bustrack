package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Vendor struct for crud
type Vendor struct {
	//for crud in organization
	Vendorid int
	Email    string
	Name     string
	Password string
	Address  string
	Contact  string
	Orgid    int
}

//to execute normal insert queries
func insertVen(db *sql.DB, vendor *Vendor) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into vendor(email, name,  password,address, contact,organization_id) values(%s,%s,%s,%s,%s,%d)",
			vendor.Email,
			vendor.Name,
			vendor.Address,
			vendor.Password,
			vendor.Contact,
			vendor.Orgid,
		),
	)
}

//to delete selected entry
func deleteVen(db *sql.DB, venid int) (sql.Result, error) {
	if venid == -1 {
		return db.Exec(fmt.Sprintf("truncate table vendor"))
	}

	return db.Exec(fmt.Sprintf("delete from vendor where vendor_id=%d", venid))

}

//to get selected entry
func getVen(db *sql.DB, venid int) (*sql.Rows, error) {
	if venid == -1 {
		return db.Query(fmt.Sprintf("select *from vendor"))
	}
	return db.Query(fmt.Sprintf("select *from vendor where vendor_id=%d", venid))
}

func getVenOrg(db *sql.DB, orgid, venid int) (*sql.Rows, error) {
	if venid == -1 {
		return db.Query(fmt.Sprintf("select *from vendor organization_id=%d", orgid))
	}
	return db.Query(fmt.Sprintf("select *from vendor where vendor_id=%d AND organization_id=%d", venid, orgid))
}

//to update selected entry
func updateVen(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
