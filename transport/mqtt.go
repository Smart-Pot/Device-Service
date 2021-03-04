package transport

import (
	"deviceservice/data"
	"deviceservice/service"
	"encoding/json"
	"strings"

	"github.com/Smart-Pot/pkg/adapter/mqtt"
)


func MakeDeviceRecordTask(c mqtt.Consumer,s service.Service) func() {
	for {
		msg := c.Consume()
		var record data.Record
		if err := json.Unmarshal(msg.Payload(),&record); err != nil {
			continue	
		}
		topic := msg.Topic()
		parts := strings.Split(topic,"/")
		if len(parts) != 3 {
			continue
		}
		s.SaveRecord(parts[2],record)
	}
}