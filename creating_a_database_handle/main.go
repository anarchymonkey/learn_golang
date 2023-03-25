package main

import (
	"context"
	"fmt"
	"os"

	// stdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Data struct {
	Id    int
	Name  string
	Age   int
	Email string
}

func selectDataFromDummyTable(conn *pgxpool.Pool) {
	rows, err := conn.Query(context.Background(), "SELECT * from dummy_table")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while querying dummy table \n%v", err)
		os.Exit(1)
	}

	var datas []Data = []Data{}

	for rows.Next() {
		var data Data = Data{}

		err := rows.Scan(&data.Id, &data.Name, &data.Age, &data.Email)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning data %v", err)
			os.Exit(1)
		}
		datas = append(datas, data)
	}

	fmt.Fprintf(os.Stdout, "The data list is: \n%v", datas)
}

func main() {
	fmt.Println("This is a demo to create connection pools using a databse handle")

	configUrl := fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres", os.Getenv("PG_USERNAME_V1"), os.Getenv("PG_PASS_V1"))

	dbConfig, err := pgxpool.ParseConfig(configUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while getting the config \n%v", err)
		os.Exit(1)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), dbConfig)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while getting the connection pool from pgx pool \n%v", err)
		os.Exit(1)
	}

	defer pool.Close()

	selectDataFromDummyTable(pool)

}
