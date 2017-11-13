package main

//go:generate go-sqltpl  sample.sqlt main.go

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbUser = "test"
	dbPass = "test"
	dbName = "test"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	log.Println("Start")

	q := WithDB(db)

	// -- sqltpl: DropDb
	// drop table foo;
	// -- end
	q.DropDb()

	// -- sqltpl: InitDb
	// create table foo (bar int);
	// -- end
	err = q.InitDb()
	checkErr(err)

	// -- sqltpl: AddFoo
	// insert into foo values(?n@@int)
	// -- end

	x := []int{1, 2, 3, 4, 5, 6}

	for _, v := range x {
		checkErr(q.AddFoo(v))
	}

	// -- sqltpl: Content
	// select bar@@int from foo
	// -- end

	rows, err := q.Content(ContentQuery{})
	checkErr(err)
	log.Println("rows", rows)

}
