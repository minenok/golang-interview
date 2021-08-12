package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

type DB struct {
	conn *sql.DB
}

func NewDB() *DB {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return &DB{db}
}

func (db *DB) Products() (*sql.Rows, error) {
	return db.conn.Query("select * from products")
}

func (db *DB) SaveProduct(name, description string) error {
	res, err := db.conn.Exec(fmt.Sprintf("insert into products(name,description) values ('%s', '%s')", name, description))
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	if id == 0 {
		return errors.New("can't insert new row")
	}
	return nil
}
