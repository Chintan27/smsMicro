package main

import (
	"fmt"
	"net/http"

	"github.com/Chintan27/smsMicro/controllers"
	"github.com/Chintan27/smsMicro/sqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	sqldb.ConnectDB()

	http.HandleFunc("/", controllers.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()
}
