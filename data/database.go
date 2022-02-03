package data

import (
	"database/sql"
	"desafio/config"
	"desafio/models"
	"fmt"
	"log"
	"strconv"
	"strings"

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

func InsertSuper(supers []models.SuperInsert) error {
	for _, super := range supers {
		log.Printf("Inserting %s\n", super.Name)
		stmt := `INSERT INTO supers (uuid, hero_name, full_name, alignment, intelligence, power, occupation, image, group_connections, relatives)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
		result, err := DB.Exec(stmt, super.UUID, super.Name, super.Biography.FullName, super.Biography.Alignment, super.Powerstats.Intelligence, super.Powerstats.Power,
			super.Work.Occupation, super.Image.URL, super.Connections.GroupAffiliations, super.Connections.Relatives)
		if err != nil {
			return fmt.Errorf("error inserting super '%s': %s", super.Name, err)
		}
		log.Printf("Inserted: %+v", result)
	}
	return nil
}

func GetSupers(args []string, filters []interface{}) []models.SuperResponse {
	var supers []models.SuperResponse
	stmt := `SELECT uuid, hero_name, full_name, alignment, intelligence, power, occupation, image, group_connections, relatives FROM supers`

	if len(filters) > 0 {
		for index := range filters {
			if index == 0 {
				stmt += " WHERE "
			} else {
				stmt += "AND "
			}
			stmt += "LOWER(" + args[index] + ")" + " = LOWER($" + strconv.Itoa(index+1) + ")"
		}
	}
	stmt += ";"

	rows, err := DB.Query(stmt, filters...)
	if err != nil {
		log.Printf("Error querying database: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			uuid             string
			heroName         string
			fullName         string
			alignment        string
			intelligence     string
			power            string
			occupation       string
			image            string
			groupConnections string
			relatives        string
		)
		if err := rows.Scan(&uuid, &heroName, &fullName, &alignment, &intelligence, &power, &occupation, &image, &groupConnections, &relatives); err != nil {
			log.Printf("Error scanning database results: %s", err)
		}
		if err := rows.Err(); err != nil {
			log.Printf("Error querying database: %s", err)
		}

		splitConnections := strings.Split(groupConnections, ", ")
		splitRelatives := strings.SplitAfter(relatives, "),")

		super := models.SuperResponse{
			UUID: uuid,
			Name: heroName,
			Powerstats: models.Powerstats{
				Intelligence: intelligence,
				Power:        power,
			},
			Biography: models.Biography{
				FullName:  fullName,
				Alignment: alignment,
			},
			Occupation:        occupation,
			RelativesCount:    len(splitRelatives),
			GroupAffiliations: splitConnections,
			Image:             image,
		}
		fmt.Printf("super: %+v", super)
		supers = append(supers, super)
	}
	return supers
}

func DeleteSupers(uuids []string) (int, error) {
	stmt := "DELETE FROM supers WHERE uuid="
	index := 0
	for _, uuid := range uuids {
		if index > 0 {
			stmt += " OR uuid="
		}
		stmt += "'" + uuid + "'"
		index++
	}
	stmt += ";"

	result, err := DB.Exec(stmt)
	if err != nil {
		return 0, fmt.Errorf("error deleting supers with UUIDs '%s': %s", strings.Join(uuids, ", "), err)
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		return int(deletedCount), fmt.Errorf("error counting affected rows: %s", err)
	}

	return int(deletedCount), nil

}
