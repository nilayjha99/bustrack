package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Trip struct for crud
type Trip struct {
	Tripid   int
	Routeid  int
	Driverid int
	Busid    int
	Start    string
	End      string
	Details  string
}

//to execute normal insert queries
func insertTrip(db *sql.DB, tr *Trip) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into trip(route_id, driver_id,  bus_id,start) values(%d,%d,%d,%s)",
			tr.Routeid,
			tr.Driverid,
			tr.Busid,
			tr.Start,
		),
	)
}

//to delete selected entry
func deleteTrip(db *sql.DB, tripid int) (sql.Result, error) {
	if tripid == -1 {
		return db.Exec(fmt.Sprintf("truncate table trip"))
	}

	return db.Exec(fmt.Sprintf("delete from trip where trip_id=%d", tripid))

}

//to get selected entry
func getTrip(db *sql.DB, tripid int) (*sql.Rows, error) {
	if tripid == -1 {
		return db.Query(fmt.Sprintf("select *from trip"))
	}
	return db.Query(fmt.Sprintf("select *from trip where trip_id=%d", tripid))
}

//to update selected entry
func updateTrip(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
