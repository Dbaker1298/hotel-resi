package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dbaker1298/hotel-resi/api"
	"github.com/Dbaker1298/hotel-resi/db"
	"github.com/Dbaker1298/hotel-resi/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client       *mongo.Client
	roomStore    db.RoomStore
	hotelStore   db.HotelStore
	userStore    db.UserStore
	bookingStore db.BookingStore
	ctx          = context.Background()
)

func seedUser(isAdmin bool, fname, lname, email, password string) *types.User {
	user, err := types.NewUserFromParams(&types.CreateUserParams{
		Email:     email,
		FirstName: fname,
		LastName:  lname,
		Password:  password,
	})
	if err != nil {
		log.Fatal(err)
	}

	user.IsAdmin = isAdmin

	insertedUser, err := userStore.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %s\n", user.Email, api.CreateTokenFromUser(user))
	return insertedUser
}

func seedRoom(size string, ss bool, price float64, hotelId primitive.ObjectID) *types.Room {
	room := types.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hotelId,
	}
	insertedRoom, err := roomStore.InsertRoom(context.Background(), &room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func seedBooking(userId, roomId primitive.ObjectID, from, till time.Time) {
	booking := &types.Booking{
		UserID:   userId,
		RoomID:   roomId,
		FromDate: from,
		TillDate: till,
	}
	resp, err := bookingStore.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("booking: ", resp.ID)

}

func seedHotel(name string, location string, rating int) *types.Hotel {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	return insertedHotel
}

func main() {
	john := seedUser(false, "John", "Doe", "john@doe.com", "supersecurepassword")
	seedUser(true, "admin", "admin", "admin@admin.com", "adminpassword123")
	seedHotel("The Grand Budapest Hotel", "Zubrowka", 5)
	seedHotel("The Overlook Hotel", "Estes Park, Colorado", 4)
	hotel := seedHotel("The Fawlty Towers", "Torquay, Devon", 1)
	seedRoom("small", true, 98.99, hotel.ID)
	seedRoom("medium", false, 199.99, hotel.ID)
	room := seedRoom("large", true, 299.99, hotel.ID)
	seedBooking(john.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))
}

func init() {
	var err error
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
	bookingStore = db.NewMongoBookingStore(client)
}
