package sqldb

import "database/sql"

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() {
	db, err := sql.Open("mysql", "chintan:chintan@123@tcp(192.168.64.2)/crmMaster")
	if err != nil {
		panic(err.Error())
	}
	DB = db
}
