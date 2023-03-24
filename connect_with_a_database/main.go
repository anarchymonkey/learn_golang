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

func printAllRowsInTable(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "select * from dummy_table")

	if err != nil {
		fmt.Fprintf(os.Stderr, "There is a error while querying \n %v", err)
		os.Exit(1)
	}

	defer rows.Close()

	datas, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[DummyTable])

	for _, data := range datas {
		fmt.Println("data is", data)
	}
}

func printRowAfterDeleting(conn *pgx.Conn) {
	const QUERY = "DELETE FROM dummy_table WHERE id=$1"

	row, err := conn.Exec(context.Background(), QUERY, 15)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while deleting data into the table \n %v", err)
		os.Exit(1)
	}

	fmt.Println(row.RowsAffected())
	fmt.Println("Row after deleting")
	printRowInTable(conn, 15)
}

func printRowAfterInserting(conn *pgx.Conn) {
	const QUERY = "INSERT INTO dummy_table(id, name, age, email) values($1,$2,$3,$4);"

	row, err := conn.Exec(context.Background(), QUERY, 15, "Aniket", 27, "aniket6991@gmail.com")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while inserting data into the table \n %v", err)
		os.Exit(1)
	}

	fmt.Println(row.RowsAffected())
	fmt.Println("Row after inserting")
	printRowInTable(conn, 15)
}

func printRowAfterUpdating(conn *pgx.Conn) {
	const QUERY = "UPDATE dummy_table set name=$1 WHERE id=$2;"

	row, err := conn.Exec(context.Background(), QUERY, "Ankush", 15)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while update data on the table \n %v", err)
		os.Exit(1)
	}

	fmt.Println("Rows affected", row.RowsAffected())
	fmt.Println("Row after updating")
	printRowInTable(conn, 14)
}

func printRowInTable(conn *pgx.Conn, id int) {

	var dummy DummyTable
	err := conn.QueryRow(context.Background(), "SELECT * from dummy_table where id=$1", id).Scan(&dummy.Id, &dummy.Name, &dummy.Age, &dummy.Email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occoured while scanning row \n %v", err)
		os.Exit(1)
	}
	fmt.Println("Dummy table row is", dummy)
}

func getDBConnectionHandler() *pgx.Conn {

	dbUrl := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, os.Getenv("PG_USERNAME_V1"), os.Getenv("PG_PASS_V1"), "postgres",
	)

	os.Setenv("DATABASE_CONFIG", dbUrl)

	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_CONFIG"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return connection
}

func main() {

	conn := getDBConnectionHandler()
	defer conn.Close(context.Background())

	fmt.Println("Getting all rows in dummy table:")
	printAllRowsInTable(conn)

	fmt.Println("\nGetting one row in dummy table:")
	printRowInTable(conn, 8)

	fmt.Println(("\nInserting one row in dummy table"))
	printRowAfterInserting(conn)

	fmt.Println(("\nUpdating one row in dummy table"))
	printRowAfterUpdating(conn)

	fmt.Println(("\nDeleting one row in dummy table"))
	printRowAfterDeleting(conn)

}
