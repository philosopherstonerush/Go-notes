package db

import (
	"context"
	"hotelReservationApi/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Making it consistent with User

type RoomStore interface {
	CreateRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, HotelStore *MongoHotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(dbname).Collection("Rooms"),
		HotelStore: HotelStore,
	}
}

func (s *MongoRoomStore) CreateRoom(ctx context.Context, Room *types.Room) (*types.Room, error) {
	resp, err := s.coll.InsertOne(ctx, Room)
	if err != nil {
		return nil, err
	}

	// Converting ID string to primitive Object ID type

	Room.ID = resp.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": Room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": Room.ID}}
	if err := s.HotelStore.UpdateHotel(ctx, filter, update); err != nil {
		return nil, err
	}
	return Room, nil
}
