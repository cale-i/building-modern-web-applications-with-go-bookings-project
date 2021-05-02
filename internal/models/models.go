package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the user model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is the room model
type Rooms struct {
	ID        int
	RoonName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is the restriction model
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

//  Reservations is the reservations model
type Reservations struct {
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
	Room      Rooms // FK
}

//  RoomRestrictions is the room restriction model
type RoomRestrictions struct {
	ID            int
	StartDate     string
	EndDate       string
	RoomID        int
	ReservationID int
	RestrictionID int

	CreatedAt   time.Time
	UpdatedAt   time.Time
	Room        Rooms        // FK
	Reservation Reservations // FK
	Restriction Restrictions // FK

}
