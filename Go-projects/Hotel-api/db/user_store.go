package db

import (
	"context"
	"fmt"
	"hotelReservationApi/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserStore interface lets you to control whatever the underlying implementation of the database is. That is, Handlers are separated from database.

type Dropper interface {
	Drop(ctx context.Context) error
}

type UserStore interface {
	Dropper //	Dropper will be common to all implementations, hotelStore etc

	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
}

// MongoDB implementation specfic

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongUserStore(client *mongo.Client, dbname string, coll string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(coll),
		// DBNAME is under db.go
	}
}

// If you use mongoDB compass, and check out the value under id, you would see some gibberish like Object<something> and thus the string which is the id, must be converted to object.
func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// Decode, then stores the information into the variable, make sure to send in the address!
	if err := s.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := s.coll.Find(ctx, bson.M{}) // empty filter for second argument
	if err != nil {
		return nil, err
	}
	// Decode all users into the slice
	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	//	InsertOne method returns the InsertOneResult object which can be invoked with InsertedID method to get the interface of the id created by mongoDB and then type casted to primitive.ID
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, userID string) error {
	objId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	// mongodb shenangians
	update := bson.D{
		{
			"$set", params.ToBson(),
		},
	}
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("--- Dropping ---")
	return s.coll.Drop(ctx)
}
