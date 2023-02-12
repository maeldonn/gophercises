package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Phone struct {
	ID     int
	Number string
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, number := range data {
		if _, err := insertPhoneNumber(db.db, number); err != nil {
			return err
		}
	}
	return nil
}

func insertPhoneNumber(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(number) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createPhoneNumbersTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func (db *DB) GetAllPhones() ([]Phone, error) {
	rows, err := db.db.Query("SELECT * FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var phones []Phone
	for rows.Next() {
		var p Phone
		if err = rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		phones = append(phones, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return phones, nil
}

func (db *DB) GetPhone(id int) (Phone, error) {
	var p Phone
	row := db.db.QueryRow("SELECT * FROM phone_numbers WHERE id=$1", id)
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		return Phone{}, err
	}
	return p, nil
}

func (db *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	row := db.db.QueryRow("SELECT * FROM phone_numbers WHERE number=$1", number)
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (db *DB) UpdatePhone(p Phone) error {
	statement := `UPDATE phone_numbers SET number=$2 WHERE id=$1`
	_, err := db.db.Exec(statement, p.ID, p.Number)
	return err
}

func (db *DB) DeletePhone(id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(statement, id)
	return err
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
        CREATE TABLE IF NOT EXISTS phone_numbers (
          id SERIAL,
          number VARCHAR(255)
        )`
	_, err := db.Exec(statement)
	return err
}

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if err != nil {
		return err
	}
	return createDB(db, dbName)
}

func createDB(db *sql.DB, dbName string) error {
	_, err := db.Exec("CREATE DATABASE " + dbName)
	return err
}
