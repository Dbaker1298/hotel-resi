package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/Dbaker1298/hotel-resi/db/fixtures"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "james", "bond", false)
	hotel := fixtures.AddHotel(db.Store, "The Grand Budapest Hotel", "Zubrowka", 5, nil)
	room := fixtures.AddRoom(db.Store, "large", true, 299.99, hotel.ID)
	from := time.Now()
	till := from.AddDate(0, 0, 5)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	fmt.Println("\nbooking -> ", booking.ID.Hex())
}
