package data_test

import (
	"context"
	"deviceservice/data"
	"log"
	"os"
	"testing"

	"github.com/Smart-Pot/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	_testDeviceID = "test-1234"
)
func Connect() {
	const (
		dbURI = "mongodb://localhost:27017"
		collection = "devices"
		dbName = "smartpot"
	)
	
	if err := db.Connect(dbURI,dbName,collection) ; err != nil {
		log.Fatal(err)
		return
	}
}

func AddTestDevice() {
	d := data.Device {
		ID: _testDeviceID,
		Records: []data.Record{},
	}
	
	if _, err := db.Collection().InsertOne(context.TODO(),d);  err != nil {
		log.Fatal(err)
	}

}

func DeleteTestDevice() {
	r,err := db.Collection().DeleteOne(context.TODO(),bson.M{"id":_testDeviceID})

	if err != nil {
		log.Fatal(err)
	}

	if r.DeletedCount <= 0 {
		log.Fatal("Document Not deleted")
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {
	Connect()
	AddTestDevice()
	m.Run()
	DeleteTestDevice()
}	