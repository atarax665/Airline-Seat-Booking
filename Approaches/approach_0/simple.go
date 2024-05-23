package main

import (
	airlines "airline-checkin-system/airline_checkin"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "host=localhost port=5432 user=admin dbname=airlineDb password=dbpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	airlines.InitializeDBRecords(db)

	fmt.Printf("enter the user id: ")
	var userID int
	_, err = fmt.Scanln(&userID)
	if err != nil {
		log.Fatalf("Invalid input for user ID: %v", err)
	}

	transaction, err := db.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
	}

	user, err := airlines.GetUser(transaction, userID)
	if err != nil {
		log.Fatalf("Invalid input for user ID: %v", err)
	} else {
		fmt.Printf("Welcome %s to SP Airlines\n", user.Name)
	}

	fmt.Printf("enter the seat id: ")
	var seatID int
	_, err = fmt.Scanln(&seatID)
	if err != nil {
		log.Fatalf("Invalid input for seat ID: %v", err)
	}

	seat, err := airlines.GetSeatByID(transaction, seatID)
	if err != nil {
		log.Fatalf("Invalid input for seat ID: %v", err)
	}

	tripID := 1

	sqlStatement := `UPDATE seats SET user_id = $1, trip_id = $2 WHERE id = $3;`
	_, err = transaction.Exec(sqlStatement, user.ID, tripID, seat.ID)

	if err != nil {
		transaction.Rollback()
		log.Fatalf("Failed to execute insert: %v", err)
	}

	if err := transaction.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Printf("User %s was added to seat %s \n", user.Name, seat.Name)
	airlines.PrettyPrintAllSeats(db)

}
