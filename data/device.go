package data

type Device struct {
	ID      string
	Records []Record
}

func AddRecord(deviceID string, record Record) error {
	return nil
}