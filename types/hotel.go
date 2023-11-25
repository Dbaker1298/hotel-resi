package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   int                  `json:"rating" bson:"rating"`
}

type Room struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	// small, normal, kingsize
	Size      string             `json:"size" bson:"size"`
	Seaside   bool               `json:"seaSide" bson:"seaSide"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotelId" bson:"hotelId"`
}
