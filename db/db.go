package db

const (
	DBNAME     = "hotel-resi"
	TestDBNAME = "hotel-resi-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
