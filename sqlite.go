package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func sqlite(u Worker) {
	database, err := sql.Open("sqlite3", "./workerData.db")
	if err != nil {
		fmt.Println("\n ошибка создателя", err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS УчётРаботяг (Имя TEXT, ВремяРаботы TEXT, Опоздания TEXT, ВремяОпоздания TEXT, Дата TEXT)")
	if err != nil {
		fmt.Println("\n ошибка statement", err)
	}
	statement.Exec()
	statement2, err := database.Prepare("INSERT INTO УчётРаботяг (Имя, ВремяРаботы, Опоздания, ВремяОпоздания, Дата) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("\n ошибка statement2", err)
	}
	statement2.Exec(u.name(), u.time(), u.penns(), u.lates(), u.wTime().Format("Mon Jan 2 15:04"))
	fmt.Println("sqlite done")
}

type element struct {
	userName string
	workTime string
	pens     string
	lateTime string
	timeNow  string
}

func sqliteImport() []element {
	db, err := sql.Open("sqlite3", "./workerData.db")
	if err != nil {
		fmt.Println("\n ошибка открытия", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from УчётРаботяг")
	if err != nil {
		fmt.Println("\n ошибка импорта Работяг", err)
	}
	defer rows.Close()
	elements := []element{}

	for rows.Next() {
		p := element{}
		err := rows.Scan(&p.userName, &p.workTime, &p.pens, &p.lateTime, &p.timeNow)
		if err != nil {
			fmt.Println("\n ошибка импорта", err)
			continue
		}
		elements = append(elements, p)
	}
	return elements
}
