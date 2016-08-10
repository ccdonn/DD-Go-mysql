package main

import (
	. "conf"
	"fmt"
	"log"
	. "repository"
)

func main() {
	fmt.Println("Hello")
	
	var (
	  id string
	  name string
	)
	
	db := InitDb()
	rows := FeedGet(db)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)

	}
}
