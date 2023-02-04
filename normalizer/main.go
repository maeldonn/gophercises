package main

import (
	"database/sql"
	"fmt"
	"strings"
	"unicode"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gophercise"
	password = "password"
	db_name  = "gophercise_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)

	must(resetDB(db))
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, db_name)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(createPhoneNumbersTable(db))
	_, err = insertPhoneNumbers(db, "1234567890")
	must(err)
	_, err = insertPhoneNumbers(db, "123 456 7891")
	must(err)
	_, err = insertPhoneNumbers(db, "(123) 456 7892")
	must(err)
	_, err = insertPhoneNumbers(db, "(123) 456-7893")
	must(err)
	_, err = insertPhoneNumbers(db, "123-456-7894")
	must(err)
	_, err = insertPhoneNumbers(db, "123-456-7890")
	must(err)
	_, err = insertPhoneNumbers(db, "1234567892")
	must(err)
	_, err = insertPhoneNumbers(db, "(123)456-7892")
	must(err)
}

func resetDB(db *sql.DB) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + db_name)
	if err != nil {
		return err
	}
	return createDB(db)
}

func createDB(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE " + db_name)
	return err
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
        CREATE TABLE IF NOT EXISTS phone_numbers (
          id SERIAL,
          value VARCHAR(255)
        )`
	_, err := db.Exec(statement)
	return err
}

func insertPhoneNumbers(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	row := db.QueryRow(statement, phone)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func normalize(phone string) string {
	var builder strings.Builder
	for _, ch := range phone {
		if unicode.IsNumber(ch) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
