package models

import (
	"context"
	"time"

	"github.com/my-crazy-lab/airline-group-services/airport-management-service/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plane struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InsertForm struct {
	Name string
}

func (m Plane) Insert(userID int64, form InsertForm) (planeId string, err error) {
	coll := db.GetClient().Database("db").Collection("books")

	objectID := primitive.NewObjectID()
	objectIDString := objectID.Hex()

	doc := Plane{Id: objectIDString, Name: form.Name, CreatedAt: time.Now()}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return result.InsertedID.Hex(), nil
}
