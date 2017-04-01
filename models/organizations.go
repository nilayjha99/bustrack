package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //for postgres driver
)

//Organization struct for crud
type Organization struct {
	Orgid    int
	Email    string
	Name     string
	Address  string
	Password string
	Contact  string
}

//to execute normal insert queries
func insertOrg(db *sql.DB, org *Organization) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into organization(email, name, address, password, contact) values(%s,%s,%s,%s,%s)",
			org.Email,
			org.Name,
			org.Address,
			org.Password,
			org.Contact,
		),
	)
}

//to delete selected entry
func deleteOrg(db *sql.DB, orgid int) (sql.Result, error) {
	if orgid == -1 {
		return db.Exec(fmt.Sprintf("truncate table organization"))
	}

	return db.Exec(fmt.Sprintf("delete from organization where organization_id=%d", orgid))
}

//to get selected entry
func getOrg(db *sql.DB, orgid int) (*sql.Rows, error) {
	if orgid == -1 {
		return db.Query(fmt.Sprintf("select *from organization"))
	}
	return db.Query(fmt.Sprintf("select *from organization where organization_id=%d", orgid))
}

//to update selected entry
func updateOrg(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(query)
}

//---------------will be included in future updates---------------------//
/*//to execute stored procedures
func adder(db *sql.DB, org *organization) {

	fetch, err := db.Query(fmt.Sprintf("select add_organization(%s,%s,%s,%s,%s)",
		org.email,
		org.name,
		org.address,
		org.password,
		org.contact,
	),
	)

	defer fetch.Close()

	if err != nil {
		panic(err)
	}

	var s string
	for fetch.Next() {
		err = fetch.Scan(&s)
		if err != nil {
			panic(err)
		}
		fmt.Println(s) //if nothing is printed means the procedure was executed successfully
	}
}*/
/*
func main() {
	//var err error
	/*db := dbs.GetDB()
	defer dbs.CloseDB(db)
*/ //map1 := make(map[string]string)

//map1["name"] = "editedone"

//map1["b"] = "'er'"

//fmt.Println(tools.UpdateBuilder("test", map1, "organization_id", "2"))
//org := new(organization)
//to insert or update
/*org := new(organization)
org.email = "'three@local.com'"
org.name = "'three'"
org.address = "'qwd'"
org.password = "'12345678'"
org.contact = "'\"office\"=>\"147\"'"
adder(db, org)*/
//to select
/*row, err := get(db, 1)

if err != nil {
	fmt.Println(err)
}
var id int
for row.Next() {
	err = row.Scan(&id, &org.email, &org.name, &org.address, &org.password, &org.contact)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, org.email, org.name, org.address, org.password, org.contact) //if nothing is printed means the procedure was executed successfully
}*/

//to delete

/*result, err := delete(db, 5)
if err != nil {
	panic(err)
} else {
	fmt.Println(result)
}*/
//}
