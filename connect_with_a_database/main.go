package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// package that we will use database/sql

/*
	Things you will learn


	Set up a database.
	Import the database driver.
	Get a database handle and connect.
	Query for multiple rows.
	Query for a single row.
	Add data.


*/

func connectHandler() {

	env := os.Environ()

	for _, val := range env {
		fmt.Println(val)
	}
	dbUrl := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, os.Getenv("PG_USERNAME_V1"), os.Getenv("PG_PASS_V1"), "postgres",
	)

	os.Setenv("DATABASE_CONFIG", dbUrl)

	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_CONFIG"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer connection.Close(context.Background())

	var greeting string

	err = connection.QueryRow(context.Background(), "select name from dummy_table;").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query row %v", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

}

func main() {

	connectHandler()

	fmt.Println("Here is a go file")

}
