package api

import (
	"bytes"
	"context"
	"encoding/json"
	"hotelReservationApi/db"
	"hotelReservationApi/types"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "hotel-reservation-test"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	// Removes the database after the test
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongUserStore(client, dbname, db.UserColl),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandle := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandle.HandlePostUser) // We are testing the post user function

	params := types.CreateUserParams{
		Email:     "someidiot@gmail.com",
		FirstName: "james",
		LastName:  "bond",
		Password:  "asndasdsd",
	}

	// Converting it to a json type
	b, _ := json.Marshal(params)

	// Making a test request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))

	// specifying the header so it doesnt panic
	req.Header.Add("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	// Decode json into the specific type
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if user.FirstName != params.FirstName {
		t.Errorf("Different first name")
	}
	if user.LastName != params.LastName {
		t.Error("Different last name")
	}
	if user.Email != params.Email {
		t.Error("Different email")
	}
}
