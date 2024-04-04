package models

import (
	"context"
	"fmt"
	"time"

	"github.com/my-crazy-lab/airline-group-services/flight-management-service/db"
	"github.com/my-crazy-lab/airline-group-services/flight-management-service/forms"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
ID: Unique identifier for the plane.
Model: Model name or type of the plane.
RegistrationNumber: Registration number assigned to the plane.
Capacity: Maximum capacity of passengers or cargo the plane can carry.
CurrentLocation: ID of the airport or port where the plane is currently located.
Status: Current operational status of the plane.
*/
/*
ID: "plane456"
Model: "Boeing 737"
RegistrationNumber: "N12345"
Capacity: 215 passengers
CurrentLocation: "airport123"
Status: "Available"
*/
type Plane struct {
	Id                 string    `bson:"_id,omitempty"`
	Model              string    `bson:"model"`
	RegistrationNumber string    `bson:"registration_number"`
	Capacity           int       `bson:"capacity"`
	CurrentLocation    string    `bson:"current_location"`
	Status             string    `bson:"status"`
	CreatedAt          time.Time `bson:"status"`
	UpdatedAt          time.Time `bson:"status"`
}

func (m Plane) Insert(form forms.InsertPlaneForm) (planeId string, err error) {
	coll := db.GetClient().Database("airline").Collection("planes")

	// convert objectId into string
	objectID := primitive.NewObjectID()
	objectIDString := objectID.Hex()

	doc := Plane{Id: objectIDString, Model: form.Model, CreatedAt: time.Now()}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println(err)

		return "", err
	}
	str, ok := result.InsertedID.(string)
	if ok {
		return str, nil
	} else {
		fmt.Println("Mongo Id not string")
		return "", nil
	}
}
