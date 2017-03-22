package main

import (
	//	"bustrack/dbs"
	"bustrack/tools"
	//	"database/sql"
	"fmt"
	//	_ "github.com/lib/pq"
)

/*type organization struct {
	email    string
	name     string
	address  string
	password string
	contact  string
}

//to execute normal insert queries
func insert(db *sql.DB, org *organization) (sql.Result, error) {
	return db.Exec(
		fmt.Sprintf("insert into organization(email, name, address, password, contact) values(%s,%s,%s,%s,%s)",
			org.email,
			org.name,
			org.address,
			org.password,
			org.contact,
		),
	)
}

//to execute stored procedures
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
}

func delete(db *sql.DB, org_id int) (sql.Result, error) {
	return db.Exec(fmt.Sprintf("delete from organization where organization_id=%d", org_id))
}

func get(db *sql.DB, org_id int) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("select *from organization where organization_id=%d", org_id))
}
*/
/*func update(db *sql.DB, query string) (sql.Result, error) {
	return db.Exec(

		, ...)
}*/

func main() {
	//var err error
	/*db := dbs.GetDB()
	defer dbs.CloseDB(db)
	*/map1 := make(map[string]string)

	map1["a"] = "1"

	map1["b"] = "'er'"

	fmt.Println(tools.UpdateBuilder("test", map1, "e", "8"))
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
}
