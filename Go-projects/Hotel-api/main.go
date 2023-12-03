package main

import (
	"context"
	"flag"
	"hotelReservationApi/api"
	"hotelReservationApi/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// I suppose this is the port the database system uses to talk with the app, this could go into db.go as well.
const db_uri = "mongodb://localhost:27017"

// Configuring our own custom error handling, copied from documentation
var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()}) // {"error": what_the_error_is}
	},
}

// Do not fret much, most of these are mongodb specific semantics, context.TODO and context.Background return the same empty context. URI is an endpoint that you are not clear with how to operate but URL shows you what protocol to use.
func main() {

	// A way to dynamically call for ports. (name, default_port, description). Call by building it and then run with "/api.exe --address :7000". If nothing is provided then default port of 5000 is used.
	listenAddress := flag.String("address", ":5000", "This is the port where you see the backend!")
	flag.Parse()

	app := fiber.New(config)

	// routes and initialization must come before the listen statement
	// UserHandler initialization
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}
	appU := app.Group("/U") // Groups the urls under one upper branch like localhost:5000/U/user
	userHandler := api.NewUserHandler(db.NewMongUserStore(client, db.DBNAME, db.UserColl))

	// routes
	appU.Get("/user/:id", userHandler.HandleGetUser) // :id --> key "id" for parsing
	appU.Post("/user", userHandler.HandlePostUser)
	appU.Get("/users", userHandler.HandleGetUsers)
	appU.Delete("/user/:id", userHandler.HandleDeleteUser) // separate DELETE method must be selected, so even if both are the same routes you have method control
	appU.Put("/user/:id", userHandler.HandleUpdateUser)    // Put method is for updating

	// Listen
	app.Listen(*listenAddress)
	// Or you can just say ":5000". And it will start the process there
}
