package api

import (
	"errors"
	"hotelReservationApi/db"
	"hotelReservationApi/types"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handlers are specified here, Make sure their starting letters are uppercase, if not then they will be private

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(f *fiber.Ctx) error {
	id := f.Params("id") // retrieves the id from the URL
	user, err := h.userStore.GetUserByID(f.Context(), id)
	if err != nil {
		return err
	}
	return f.JSON(user)
}

func (h *UserHandler) HandleGetUsers(f *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(f.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { // if err is same as mongo.ErrNoDocument then execute this, this situation happens when you try to get user that was already deleted.
			f.JSON(map[string]string{"error": "Not found"})
		}
		return err
	}
	return f.JSON(users)
}

func (h *UserHandler) HandlePostUser(f *fiber.Ctx) error {
	var params types.CreateUserParams

	// parse the POST request and store it in params. JSON request looks like {"firstName": blah, "lastName": bond etc}, see the json tags for the struct type.

	if err := f.BodyParser(&params); err != nil {
		return err
	}

	// Validate the sent Params

	if err := params.Validate(); err != nil {
		return f.JSON(err) // normal err return wont handle a map
	}

	// Create user from CreateUserParams

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	// Insert the user, and return the user with the object id created by mongoDB

	insertedUser, err := h.userStore.CreateUser(f.Context(), user)
	if err != nil {
		return err
	}

	return f.JSON(insertedUser)
}

func (h *UserHandler) HandleDeleteUser(f *fiber.Ctx) error {
	userID := f.Params("id")
	err := h.userStore.DeleteUser(f.Context(), userID)
	if err != nil {
		return err
	}
	return f.JSON(map[string]string{"deleted": userID})
}

func (h *UserHandler) HandleUpdateUser(f *fiber.Ctx) error {
	var (
		userID = f.Params("id")
		params types.UpdateUserParams // Separate struct for updation validation
	)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	if err := f.BodyParser(&params); err != nil {
		return err
	}
	filter := bson.M{"_id": oid} // filter to get the user we want to update
	if err = h.userStore.UpdateUser(f.Context(), filter, params); err != nil {
		return err
	}
	return f.JSON(map[string]string{"updated": userID})
}
