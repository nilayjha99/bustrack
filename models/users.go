package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Users struct for crud
type Users struct {
	Userid         int
	Email          string
	Password       string
	Organizationid int
}

//to execute normal insert queries
func insertUser(db *sql.DB, user *Users) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into users(email, password,  organization_id) values(%s,%s,%d)",
			user.Email,
			user.Password,
			user.Organizationid,
		),
	)
}

//to delete selected entry
func deleteUser(db *sql.DB, userid int) (sql.Result, error) {
	if userid == -1 {
		return db.Exec(fmt.Sprintf("truncate table users"))
	}

	return db.Exec(fmt.Sprintf("delete from users where user_id=%d", userid))

}

//to get selected entry
func getUsers(db *sql.DB, userid int) (*sql.Rows, error) {
	if userid == -1 {
		return db.Query(fmt.Sprintf("select *from users"))
	}
	return db.Query(fmt.Sprintf("select *from users where user_id=%d", userid))
}

func getUserOrg(db *sql.DB, orgid, userid int) (*sql.Rows, error) {
	if orgid == -1 {
		return db.Query(fmt.Sprintf("select *from route where organization_id=%d", orgid))
	}
	return db.Query(fmt.Sprintf("select *from users where user_id=%d AND organization_id=%d", userid, orgid))
}

//to update selected entry
func updateUsers(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}
