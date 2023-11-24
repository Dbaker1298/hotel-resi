package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	RoomID     primitive.ObjectID `json:"roomId" bson:"roomId,omitempty"`
	NumPersons int                `json:"numPersons" bson:"numPersons,omitempty"`
	FromDate   time.Time          `json:"fromDate" bson:"fromDate,omitempty"`
	TillDate   time.Time          `json:"tillDate" bson:"tillDate,omitempty"`
}
