package main

import (
	"fmt"
	"strings"
	"unicode"

	phonedb "github.com/maeldonn/gophercises/normalizer/db"
)

const (
	driver_name = "postgres"
	host        = "localhost"
	port        = 5432
	user        = "gophercise"
	password    = "password"
	db_name     = "gophercise_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(phonedb.Reset(driver_name, psqlInfo, db_name))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, db_name)
	must(phonedb.Migrate(driver_name, psqlInfo))

	db, err := phonedb.Open(driver_name, psqlInfo)
	must(err)
	defer db.Close()

	must(db.Seed())

	phones, err := db.GetAllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				must(db.UpdatePhone(p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
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
