package airline_checkin

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
)

type Airline struct {
	ID   uint
	Name string
}

type Flight struct {
	ID        uint
	AirlineID uint
	Name      string
}

type Trip struct {
	ID         uint
	FlightID   uint
	FlightTime time.Time
}

type User struct {
	ID   uint
	Name string
}

type Seat struct {
	ID     uint
	Name   string
	TripID uint
	UserID sql.NullInt64
}

func ClearAllContents(db *sql.DB) {
	tables := []string{"seats", "users", "trips", "flights", "airlines"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table))
		if err != nil {
			log.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}
	fmt.Println("All tables truncated successfully.")
}

func InitializeDBRecords(db *sql.DB) {

	ClearAllContents(db)

	_, err := db.Exec(`INSERT INTO airlines (name) VALUES ('Air India') RETURNING id;`)
	if err != nil {
		log.Fatalf("Failed to insert airline: %v", err)
	}

	var airlineID int
	err = db.QueryRow(`SELECT id FROM airlines WHERE name = 'Air India';`).Scan(&airlineID)
	if err != nil {
		log.Fatalf("Failed to query airline ID: %v", err)
	}

	_, err = db.Exec(`INSERT INTO flights (name, airline_id) VALUES ('AIR_01', $1);`, airlineID)
	if err != nil {
		log.Fatalf("Failed to insert flight: %v", err)
	}

	var flightID int
	err = db.QueryRow(`SELECT id FROM flights WHERE name = 'AIR_01';`).Scan(&flightID)
	if err != nil {
		log.Fatalf("Failed to query flight ID: %v", err)
	}

	specificTime := time.Date(2024, time.April, 19, 21, 0, 0, 0, time.UTC)
	_, err = db.Exec(`INSERT INTO trips (flight_id, flight_time) VALUES ($1, $2);`, flightID, specificTime)
	if err != nil {
		log.Fatalf("Failed to insert trip: %v", err)
	}

	tripID := 1
	for i := 0; i < 144; i++ {

		userName := faker.Name()
		_, err = db.Exec(`INSERT INTO users (name) VALUES ($1) RETURNING id;`, userName)
		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}

		seatRow := i/6 + 1
		seatLetter := rune('A' + (i % 6))
		seatName := fmt.Sprintf("%d-%c", seatRow, seatLetter)
		_, err := db.Exec(`INSERT INTO seats (name, trip_id) VALUES ($1, $2);`, seatName, tripID)
		if err != nil {
			log.Fatalf("Failed to insert seat: %v", err)
		}
	}

	fmt.Println("Data insertion complete")
}

func ClearAllUsersFromSeats(db *sql.DB, tripID int) {
	_, err := db.Exec("TRUNCATE TABLE seats RESTART IDENTITY CASCADE;")
	if err != nil {
		log.Fatalf("Failed to truncate seat: %v", err)
	}
	for i := 0; i < 120; i++ {
		seatRow := i/5 + 1
		seatLetter := rune('A' + (i % 5))
		seatName := fmt.Sprintf("%d-%c", seatRow, seatLetter)
		_, err = db.Exec(`INSERT INTO seats (name, trip_id) VALUES ($1, $2);`, seatName, tripID)
		if err != nil {
			log.Fatalf("Failed to insert seat: %v", err)
		}
	}
}

func PrettyPrintAllSeats(db *sql.DB) {
	seats, _ := GetAllSeats(db)
	for i := 0; i < 6; i++ {
		var start = i
		for j := 0; j < 20; j++ {
			index := start + j*6
			if index >= len(seats) {
				fmt.Println("Index out of range, not enough seats to display.")
				return
			}
			seat := seats[index]

			seatCode := "."
			if seat.UserID.Valid {
				seatCode = "X"
			}
			fmt.Printf("%s ", seatCode)

		}

		if i == 2 {
			fmt.Println()
		}
		fmt.Println()
	}

	fmt.Println()

}

func GetAvailableSeatWithUpdateSkipLockedRows(db *sql.Tx, tripID int) (Seat, error) {
	var seat Seat
	sqlStatement := `SELECT id,name,trip_id, user_id FROM seats where trip_id= $1
											and user_id is null order by id limit 1 for update skip locked;`
	err := db.QueryRow(sqlStatement, tripID).Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)
	if err != nil {
		log.Fatalf("Could not get new seat: %v", err)
		return Seat{}, err
	}
	return seat, nil
}

func GetAvailableSeatWithUpdate(db *sql.Tx, tripID int) (Seat, error) {
	var seat Seat
	sqlStatement := `SELECT id,name,trip_id, user_id FROM seats where trip_id= $1
											and user_id is null order by id limit 1 for update;`
	err := db.QueryRow(sqlStatement, tripID).Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)
	if err != nil {
		log.Fatalf("Could not get new seat: %v", err)
		return Seat{}, err
	}
	return seat, nil
}

func GetAvailableSeat(transaction *sql.Tx, tripID int) (Seat, error) {
	var seat Seat
	sqlStatement := `SELECT id,name,trip_id, user_id FROM seats where trip_id= $1
											and user_id is null order by id limit 1;`
	err := transaction.QueryRow(sqlStatement, tripID).Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)
	if err != nil {
		log.Fatalf("Could not get new seat: %v", err)
		return Seat{}, err
	}
	return seat, nil
}

func GetSeatByID(transaction *sql.Tx, seatID int) (Seat, error) {
	var seat Seat
	sqlStatement := `SELECT * FROM seats where id = $1;`
	err := transaction.QueryRow(sqlStatement, seatID).Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)
	if err != nil {
		log.Fatalf("Could not get new seat: %v", err)
		return Seat{}, err
	}
	return seat, nil
}

func PrintUserSeats(db *sql.DB) {
	fmt.Println("-----")
	sqlStatement := `
        SELECT s.id, s.name, s.trip_id, s.user_id, u.id, u.name 
        FROM seats s 
        LEFT JOIN users u ON s.user_id = u.id 
        WHERE s.user_id IS NOT NULL
        ORDER BY s.id;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var seat Seat
		var user User
		var userID sql.NullInt64
		var userName sql.NullString

		err := rows.Scan(&seat.ID, &seat.Name, &seat.TripID, &userID, &user.ID, &userName)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
		}

		if userID.Valid {
			user.ID = uint(userID.Int64)
			user.Name = userName.String
			fmt.Printf("UserID: %d Name: %s Seat: %s\n", user.ID, user.Name, seat.Name)
		}
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
	}
	fmt.Println("-----")

}

func GetAllSeats(db *sql.DB) ([]Seat, error) {
	seats := make([]Seat, 0)
	sqlStatement := `SELECT * FROM seats ORDER by id;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return nil, err
	}
	for rows.Next() {
		var seat Seat
		rows.Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)
		seats = append(seats, seat)
	}
	return seats, nil
}

func GetUser(transaction *sql.Tx, userID int) (User, error) {
	var user User
	sqlStatement := `SELECT name,id FROM users where id = $1;`
	err := transaction.QueryRow(sqlStatement, userID).Scan(&user.Name, &user.ID)
	if err != nil {
		log.Fatalf("Could not get new user: %v", err)
		return User{}, err
	}
	return user, nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	users := make([]User, 0)
	sqlStatement := `SELECT * FROM users ORDER by id;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return nil, err
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	return users, nil
}

func ShowAllSeats(db *sql.DB) {
	sqlStatement := `SELECT id, name, trip_id, user_id FROM seats;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	fmt.Println("ID | Name | Trip ID | User ID")
	for rows.Next() {
		var id int
		var name string
		var tripID int
		var userID int
		err = rows.Scan(&id, &name, &tripID, &userID)
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}
		fmt.Printf("%d | %s | %d | %d\n", id, name, tripID, userID)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}
}
