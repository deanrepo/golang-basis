package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Initiliazes corresponding files.
)

func main() {
	timeTest1()
	// timeTest2()
}

func timeTest1() {
	db, err := sql.Open("mysql", "dean:Dean#168168@tcp(192.168.1.12:3306)/test?collation=utf8_general_ci&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var myTime time.Time
	rows, err := db.Query("SELECT current_timestamp()")

	if rows.Next() {
		if err = rows.Scan(&myTime); err != nil {
			panic(err)
		}
	}

	fmt.Println(myTime)
}

// timeTest2 test mysql datetime and date scan to time.
func timeTest2() {
	db, err := sql.Open("mysql", "dean:Dean#168168@tcp(192.168.1.12:3306)/test?collation=utf8_general_ci&parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	var id int
	var regTime, logTime time.Time
	rows, err := db.Query("SELECT id, reg_time, login_time FROM datetime_test")

	if rows.Next() {
		if err = rows.Scan(&id, &regTime, &logTime); err != nil {
			panic(err)
		}
	}

	fmt.Println(id, regTime, logTime)

	y, m, d := regTime.Date()
	fmt.Println(y, m, d)

	y, m, d = logTime.Date()
	fmt.Println(y, m, d)
}
