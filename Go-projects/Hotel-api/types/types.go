// Model-View-Controller, here lies the models, aka struts
package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Constants for validation and password cost
const (
	bcryptCost      = 12
	MinFirstNameLen = 3
	MinLastNameLen  = 3
	MinPasswordLen  = 7
)

// bson, json are tags for the database fields and json respectively. omitempty, will omit the value in the return statement if there is nothing under the field, while "-" returns nothing irrespective of whether the value is there or not. Make sure there is no space!. bson omitempty, will create an id for it automatically if nothing is specified.
type User struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p *UpdateUserParams) ToBson() bson.M {
	m := bson.M{} // Bson.M is just same as map[string]any
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"Password"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	// Encrpyting the password using bcrypt package, type casting it to slice of bytes and then second argument is the cost.
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw), // type casting it back to string from bytes
	}, nil
}

func (p *CreateUserParams) Validate() map[string]string {

	errors := map[string]string{}
	// Making it easier for frontend to respond to various errors that may happen at the same time

	if len(p.FirstName) < MinFirstNameLen {
		errors["firstName"] = fmt.Sprintf("first name length should be atleast %d", MinFirstNameLen)
	}
	if len(p.LastName) < MinLastNameLen {
		errors["lastName"] = fmt.Sprintf("last name length should be atleast %d", MinLastNameLen)
	}
	if len(p.Password) < MinPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be atleast %d", MinPasswordLen)

	}
	if !isEmailValid(p.Email) {
		errors["email"] = "Invalid email"
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
