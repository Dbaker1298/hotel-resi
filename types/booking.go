package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `json:"id,omitempty"         bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId,omitempty"     bson:"userId,omitempty"`
	RoomID     primitive.ObjectID `json:"roomId,omitempty"     bson:"roomId,omitempty"`
	NumPersons int                `json:"numPersons,omitempty" bson:"numPersons,omitempty"`
	FromDate   time.Time          `json:"fromDate,omitempty"   bson:"fromDate,omitempty"`
	TillDate   time.Time          `json:"tillDate,omitempty"   bson:"tillDate,omitempty"`
	Canceled   bool               `json:"canceled"             bson:"canceled"`
}
