package data

import (
	"database/sql"
	"desafio/config"
	"desafio/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDBConnection() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database! Terminating: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to database! Terminating: %s", err)
	}

	log.Println("Successfully connected to database!")

	DB = db
}

func InsertSuper(supers []models.Super) error {
	for _, super := range supers {
		log.Printf("Inserting %s\n", super.Name)
		stmt := `INSERT INTO supers (uuid, hero_name, full_name, alignment, intelligence, power, occupation, image, group_connections, relatives)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
		log.Println(stmt)
		// args := []string{super.UUID, super.Name, super.Biography.FullName, super.Biography.Alignment, super.Powerstats.Intelligence, super.Powerstats.Power,
		// 	super.Work.Occupation, super.Image.URL, super.Connections.GroupAffiliations, super.Connections.Relatives}
		// log.Println(args)
		result, err := DB.Exec(stmt, super.UUID, super.Name, super.Biography.FullName, super.Biography.Alignment, super.Powerstats.Intelligence, super.Powerstats.Power,
			super.Work.Occupation, super.Image.URL, super.Connections.GroupAffiliations, super.Connections.Relatives)
		if err != nil {
			return fmt.Errorf("Error inserting super '%s': %s", super.Name, err)
		}
		log.Printf("Inserted: %+v", result)
	}
	return nil
}
