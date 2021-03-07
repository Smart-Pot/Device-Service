package service_test

import (
	"context"
	"deviceservice/data"
	"deviceservice/service"
	"log"
	"os"
	"testing"

	"github.com/Smart-Pot/pkg/db"
	kitlog "github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)


var (
	_testDeviceID = "test-1234"
	_service service.Service
)

func TestMain(m *testing.M) {
	_service = service.NewService(kitlog.NewNopLogger(),nil)
	Connect()
	AddTestDevice()
	m.Run()
	DeleteTestDevice()
}


func TestAddRecord(t *testing.T) {
	r := data.Record{
		Humidity: "123",
		Temperature: "15",
		Light: "10p",
	}
	assert.Nil(t,_service.AddRecord(context.TODO(),_testDeviceID,r))

}

	

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