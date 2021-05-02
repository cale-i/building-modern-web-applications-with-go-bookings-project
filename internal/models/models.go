package models

import "time"

// Reservation holds reservation data
// type Reservation struct {
// 	FirstName string
// 	LastName  string
// 	Email     string
// 	Phone     string
// }

// Uses is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room is the room model
type Room struct {
	ID        int
	RoonName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//  Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate string
	EndDate   string
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room // FK
}

//  RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     string
	EndDate       string
	RoomID        int
	ReservationID int
	RestrictionID int

	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room        Room        // FK
	Reservation Reservation // FK
	Restriction Restriction // FK

}
