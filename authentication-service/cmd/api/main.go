package main

import (
	"authentication/data"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// db stuff
	conn := connectToDB()
	if conn == nil {
		log.Panic("cannot connect to db")
	}

	//setup config, create a new copy
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("DB_PORT")
	// Format the connection string
	// connStr := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=%s", user, password, host, dbname, sslmode)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)
	// connStr := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=%s password=%s", user, dbname, host, sslmode, password)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		connection, err := openDB(connStr)
		if err != nil {
			log.Println("Postgresql not ready")

		} else {
			log.Println("connected to posgresql")
			return connection
		}

		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			log.Println("time out after 5 seconds, stop database connection attempts")
			return nil
		}

	}
}
