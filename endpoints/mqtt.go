package endpoints

import "github.com/Smart-Pot/pkg/adapter/mqtt"

func MakeDeviceRecordConsumer(client mqtt.Client) (mqtt.Consumer, error) {
	return client.Subscribe("device/basic/+")
} 