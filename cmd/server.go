package cmd

import (
	"deviceservice/endpoints"
	"deviceservice/service"
	"deviceservice/transport"
	golog "log"
	"net/http"
	"os"
	"time"

	"github.com/Smart-Pot/pkg"
	"github.com/Smart-Pot/pkg/adapter/amqp"
	"github.com/Smart-Pot/pkg/adapter/mqtt"
	"github.com/go-kit/kit/log"
)

func startServer() error {
	// Defining Logger
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)


	warningProducer, err := amqp.MakeProducer("WarningUser")

	if err != nil {
		return err
	}

	service := service.NewService(logger,warningProducer)
	endpoint := endpoints.MakeEndpoints(service)
	handler := transport.MakeHTTPHandlers(endpoint, logger)


	mqttClient,err := mqtt.Connect("mqtt://localhost:1883","","","device-service")
	if err != nil {
		return err
	}
	
	if err := runMQTTTasks(mqttClient,service,logger); err != nil {
		return err
	}

	l := golog.New(os.Stdout, "DEVICE-SERVICE", 0)
	// Set handler and listen given port
	s := http.Server{
		Addr:         pkg.Config.Server.Address, // configure the bind address
		Handler:      handler,                   // set the default handler
		ErrorLog:     l,                         // set the logger for the server
		ReadTimeout:  5 * time.Second,           // max time to read request from the client
		WriteTimeout: 10 * time.Second,          // max time to write response to the client
		IdleTimeout:  120 * time.Second,         // max time for connections using TCP Keep-Alive
	}
	return s.ListenAndServe()
}


func runMQTTTasks(client mqtt.Client,s service.Service,l log.Logger) error {
	c,err := endpoints.MakeDeviceRecordConsumer(client)
	if err != nil {
		return err
	}
	go transport.MakeDeviceRecordTask(l,c,s)()

	return nil
}