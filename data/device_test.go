package data_test

import (
	"context"
	"deviceservice/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevice(t *testing.T) {
	d,err := data.GetDevice(context.TODO(),_testDeviceID)
	assert.Nil(t,err)
	assert.NotNil(t,d)
}

func TestAddRecord(t *testing.T) {
	records :=[]data.Record {
		{
			Humidity: "1024",
			Temperature: "14C",
			Light: "14P",
		},
		{
			Humidity: "1224",
			Temperature: "34C",
			Light: "14t",
		},
	}

	for _,r := range records {
		assert.Nil(t,data.AddRecord(context.TODO(),_testDeviceID,r))
	}

}