package data

import (
	"context"
	"errors"

	"github.com/Smart-Pot/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)


var (
	ErrNoRecordAdded = errors.New("no record added")
)

type Device struct {
	ID      string
	Records []Record
}

func AddRecord(ctx context.Context,deviceID string, record Record) error {
	filter := bson.M{"id" : deviceID}
    update := bson.M{"$push":bson.M{"records":record}}
	res,err := db.Collection().UpdateOne(ctx,filter,update)
	if res.ModifiedCount <= 0 {
		return ErrNoRecordAdded
	}
	return err
}

func GetDevice(ctx context.Context,deviceID string) (*Device,error) {
	filter := bson.M{"id" : deviceID}
	res := db.Collection().FindOne(ctx,filter)

	if res.Err() != nil {
		return nil,res.Err()
	}

	var dev Device

	if err := res.Decode(&dev); err != nil {
		return nil,err
	}
	return &dev,nil
}