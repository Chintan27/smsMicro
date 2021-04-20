package controllers

import (
	"fmt"
	"net/http"

	"github.com/Chintan27/smsMicro/sqldb"
)

// HelloWorld returns Hello, World
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if err := sqldb.DB.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Write([]byte("Hello, World"))
}
