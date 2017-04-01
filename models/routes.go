package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Route struct for crud
type Route struct {
	Routeid        int
	Organizationid int
	Source         string
	Destination    string
	Coords         string
}

//to execute normal insert queries
func insertRoute(db *sql.DB, route *Route) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into route(organization_id, source,  destination,coords) values(%d,%s,%s,%s)",
			route.Organizationid,
			route.Source,
			route.Destination,
			route.Coords,
		),
	)
}

//to delete selected entry
func deleteRoute(db *sql.DB, routeid int) (sql.Result, error) {
	if routeid == -1 {
		return db.Exec(fmt.Sprintf("truncate table route"))
	}

	return db.Exec(fmt.Sprintf("delete from route where route_id=%d", routeid))

}

//to get selected entry
func getRoute(db *sql.DB, routeid int) (*sql.Rows, error) {
	if routeid == -1 {
		return db.Query(fmt.Sprintf("select *from route"))
	}
	return db.Query(fmt.Sprintf("select *from route where route_id=%d", routeid))
}

func getRouteOrg(db *sql.DB, orgid, routeid int) (*sql.Rows, error) {
	if routeid == -1 {
		return db.Query(fmt.Sprintf("select *from route where organization_id=%d", orgid))
	}
	return db.Query(fmt.Sprintf("select *from route where route_id=%d AND organization_id=%d", routeid, orgid))
}

//to update selected entry
func updateRoute(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
