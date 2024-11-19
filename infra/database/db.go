package database

import (
	"fmt"
	"luizalabs-chalenge/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	conn *sqlx.DB
}

var Conn *DatabaseConnection

func (db *DatabaseConnection) Close() error {
	if db == nil || db.conn == nil {
		return ErrDbNil
	}
	return db.conn.Close()
}

func (db *DatabaseConnection) Ping() error {
	if db == nil || db.conn == nil {
		return ErrDbNil
	}
	return db.conn.Ping()
}

func (db *DatabaseConnection) GetConn() (*sqlx.DB, error) {
	if db == nil || db.conn == nil {
		return nil, ErrDbNil
	}

	return db.conn, nil
}

func Connect() (*DatabaseConnection, error) {
	host := utils.GetEnv("DB_HOST", "0.0.0.0")
	port := utils.GetEnv("DB_PORT", "5432")
	dbName := utils.GetEnv("DB_NAME", "luizalabs")
	username := utils.GetEnv("DB_USERNAME", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "123")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	Conn = &DatabaseConnection{db}

	return Conn, nil
}
