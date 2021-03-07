package service

import (
	"context"
	"deviceservice/data"
	"time"

	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Service interface {
	AddRecord(ctx context.Context,deviceID string, record data.Record) error
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

func (s *service) AddRecord(ctx context.Context,deviceID string, record data.Record) error {
	var err error
	defer func(beginTime time.Time) {
		level.Info(s.logger).Log(
			"function", "AddRecord",
			"param:deviceID", deviceID,
			"param:record", record,
			"result:err", err,
			"took", time.Since(beginTime))
	}(time.Now())
	if err = data.AddRecord(context.TODO(),deviceID,record); err != nil {
		return err
	}
	return nil
}