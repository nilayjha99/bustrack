package models

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq" //for postgres driver
)

//Permit struct for crud
type Permit struct {
	Organizationid int
	Email          string
}

//to execute normal insert queries
func insertPerm(db *sql.DB, per *Permit) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into permittedusers(organization_id, email) values(%d,%s)",
			per.Organizationid,
			per.Email,
		),
	)
}

//to delete selected entry
func deletePerm(db *sql.DB, email string) (sql.Result, error) {
	if strings.Compare(email, "all") == 0 {
		return db.Exec(fmt.Sprintf("truncate table permittedusers"))
	}

	return db.Exec(fmt.Sprintf("delete from permittedusers where email=%s", email))

}

//to get selected entry
func getPerm(db *sql.DB, email string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select *from permittedusers"))
}

func getPermOrg(db *sql.DB, orgid int) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select *from permittedusers where organization_id=%d", orgid))
}

//to update selected entry
func updatePerm(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
