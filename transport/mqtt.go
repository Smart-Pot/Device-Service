package transport

import (
	"context"
	"deviceservice/data"
	"deviceservice/service"
	"encoding/json"
	"errors"
	"strings"

	"github.com/Smart-Pot/pkg/adapter/mqtt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)


func MakeDeviceRecordTask(logger log.Logger,c mqtt.Consumer,s service.Service) func() {
	er := getErFunction(logger,"DeviceRecordTask")
	return func() {
		for {
			msg := c.Consume()
			var record data.Record
			if err := json.Unmarshal(msg.Payload(),&record); err != nil {
				er(err)
				continue	
			}
			topic := msg.Topic()
			parts := strings.Split(topic,"/")
			if len(parts) != 3 {
				er(errors.New("Invalid topic "+topic))
				continue
			}
			s.AddRecord(context.Background(),parts[2],record)
		}
	}
}


func getErFunction(logger log.Logger,taskName string) func(error) {
	return func (err error)  {
		level.Error(logger).Log("type","MQTT","name",taskName,"err",err.Error())
	}
}