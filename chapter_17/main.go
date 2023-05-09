package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// map this type to the record in the table

type Course struct {
	ID      string
	Details string
}

func GetRecords(db *sql.DB) {
	fmt.Println("In Getting Records")
	results, err := db.Query("Select * FROM Course;")
	fmt.Println("Query has been made.")
	if err != nil {
		fmt.Println("Error in the Query itself, perhaps wrong credentials")
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var course Course
		err = results.Scan(&course.ID,
			&course.Details)
		if err != nil {
			fmt.Println("Errors from the database query results")
			panic(err.Error())
		}

		fmt.Println(course.ID, "-", course.Details)
	}
}

func InsertRecord(db *sql.DB, ID string, Details string) {
	// use parameterized SQL statements
	result, err := db.Exec(
		"insert into Course values (?, ?)", ID, Details)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func EditRecord(db *sql.DB, ID string, Details string) {
	result, err := db.Exec(
		"update Course set Details=? where ID=?",
		Details, ID)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if count, err := result.RowsAffected(); err == nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func DeleteRecord(db *sql.DB, ID string) {
	result, err := db.Exec(
		"delete from Course where ID=?", ID)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		if count, err := result.RowsAffected(); err != nil {
			fmt.Println(count, "row(s) affected")
		}
	}
}

func main() {
	// Use mysql as driverName and a valid DNS
	db, err := sql.Open("mysql",
		"root:root@/CoursesDB")

	// handle error
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database object created")
		// InsertRecord(db, "IOS101", "iOS Programming")
		// EditRecord(db, "IOS101", "SwiftUi Programming")
		DeleteRecord(db, "IOS101")
		GetRecords(db)
	}
	// defer the close till after the main function has
	// finished executing
	defer db.Close()
}
