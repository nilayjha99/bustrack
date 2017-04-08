package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Bus struct for crud
type Bus struct {
	Busid int
	Venid int
	//VehicleNo
	//Tripid
}

//to execute normal insert queries
func insertBus(db *sql.DB, bus *Bus) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into bus(vendor_id) values(%d)",
			bus.Venid,
		),
	)
}

//to delete selected entry
func deleteBus(db *sql.DB, busid int) (sql.Result, error) {
	if busid == -1 {
		return db.Exec(fmt.Sprintf("truncate table bus"))
	}

	return db.Exec(fmt.Sprintf("delete from bus where bus_id=%d", busid))

}

//to get selected entry
func getBus(db *sql.DB) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select *from bus"))
}

func getBusVen(db *sql.DB, vendorid int) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select *from bus where vendor_id=%d", vendorid))
}

//to update selected entry
func updateBus(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
