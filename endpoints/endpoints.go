package endpoints

import (
	"deviceservice/data"
)

type Endpoints struct {
	
}


type AddRecordRequest struct {
	Record data.Record
	DeviceID string
}

type DeviceResponse struct {
	Success int
	Message string
	DeviceID string
}