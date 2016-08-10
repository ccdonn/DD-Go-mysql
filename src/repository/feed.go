package repository

import (
  "log"
//  "reflect"
  "database/sql"
)

func FeedGet(db *sql.DB) *sql.Rows {
  rows, err := db.Query("select ID id, Title name from FEED")
  
//  log.Println(reflect.TypeOf(rows))
	
	if err != nil {
		log.Fatal(err)
		return nil
	}
	
	return rows
}