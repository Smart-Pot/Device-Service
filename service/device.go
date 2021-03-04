package service

import (
	"deviceservice/data"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/go-kit/kit/log"
)

type Service interface {
	SaveRecord(deviceID string, record data.Record) error
}

type service struct {
	logger       log.Logger
	warnProducer amqp.Producer
}

func NewService(l log.Logger, warnProducer amqp.Producer) Service {
	return &service{
		logger:       l,
		warnProducer: warnProducer,
	}
}

func (s *service) SaveRecord(deviceID string, record data.Record) error {
	if err := data.AddRecord(deviceID,record); err != nil {
		return err
	}
	return nil
}