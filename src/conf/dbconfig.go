package conf

import (
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
//	"reflect"
)

func InitDb() *sql.DB {
  db, err := sql.Open("mysql", "root:45645655@tcp(127.0.0.1:3306)/zipdb")
  checkErr(err, "Connection Fail")
//  log.Println(reflect.TypeOf(db))
  return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
