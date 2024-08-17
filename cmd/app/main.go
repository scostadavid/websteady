package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	// monitorRepo "github.com/scostadavid/tiger/internal/app/monitor/repository"
	monitorModel "github.com/scostadavid/tiger/internal/app/monitor/entity"
	// monitorServ "github.com/scostadavid/tiger/internal/app/monitor/service"
)

func main() {
	db, err := sql.Open("sqlite3", "./meubanco.db")
	if err != nil {
		log.Fatal("err", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS monitorables (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100),
			url VARCHAR(100)
		)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO monitorables (name, url) VALUES
		('google', 'https://google.com')
	`)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name, uri FROM monitorables`)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("rows selected")

	for rows.Next() {
		var monitorable monitorModel.Monitorable
		err := rows.Scan(&monitorable.ID, &monitorable.Name, &monitorable.URI)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", monitorable.ID, monitorable.Name, monitorable.URI)
	}

	// Initialize repositories and services
	// monitorableRepository := monitorRepo.NewMonitorableRepository() // db
	// monitorableService := monitorServ.NewMonitorableService(monitorableRepository)
	fmt.Println("tiger")
	// init api handler
	// setup routes
}
