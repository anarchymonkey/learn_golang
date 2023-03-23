package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DummyTable struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func connectHandler() {

	// env := os.Environ()

	// for _, val := range env {
	// 	fmt.Println(val)
	// }
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

	// rows, err := connection.Query(context.Background(), "select * from dummy_table;")

	// err = pgxscan.Select(context.Background(), connection, &rows, "select * from dummy_table;")

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "The error is \n %v", err)
	// 	os.Exit(1)
	// }

	rows, err := connection.Query(context.Background(), "select * from dummy_table")

	if err != nil {
		fmt.Fprintf(os.Stderr, "There is a error while querying \n %v", err)
		os.Exit(1)
	}

	defer rows.Close()

	datas, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[DummyTable])

	for _, data := range datas {
		fmt.Println("data is", data.Name)
	}
}

func main() {

	connectHandler()

	fmt.Println("Here is a go file")

}
