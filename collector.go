package main // import "github.com/urlgrey/streammarker-collector"

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/urlgrey/streammarker-collector/handlers"
)

const (
	DEFAULT_QUEUE_NAME = "streammarker-collector-messages"
)

func main() {
	mainServer := negroni.New()

	// Token auth middleware
	tokenVerification := handlers.NewTokenVerificationMiddleware()
	tokenVerification.Initialize()
	mainServer.Use(negroni.HandlerFunc(tokenVerification.Run))

	// Create external service connections
	sqsService := createSQSConnection()

	// get queue name
	queueName := os.Getenv("STREAMMARKER_QUEUE_NAME")
	if queueName == "" {
		queueName = DEFAULT_QUEUE_NAME
	}

	// Initialize HTTP service handlers
	router := mux.NewRouter()
	queueURL := findQueueURL(sqsService, queueName)
	handlers.InitializeRouterForSensorReadingsHandlers(router, sqsService, queueURL)
	mainServer.UseHandler(router)
	go mainServer.Run(":3000")

	// Run healthcheck service
	healthCheckServer := negroni.New()
	healthCheckRouter := mux.NewRouter()
	handlers.InitializeRouterForHealthCheckHandler(healthCheckRouter, sqsService, queueName)
	healthCheckServer.UseHandler(healthCheckRouter)
	healthCheckServer.Run(":3100")
}

func createSQSConnection() *sqs.SQS {
	config := &aws.Config{}
	if endpoint := os.Getenv("STREAMMARKER_SQS_ENDPOINT"); endpoint != "" {
		config.Endpoint = &endpoint
	}

	return sqs.New(session.New(), config)
}

func findQueueURL(sqsService *sqs.SQS, queueName string) (queueURL string) {
	// check the environment variable first
	if queueURL = os.Getenv("STREAMMARKER_SQS_QUEUE_URL"); queueURL != "" {
		return
	}

	// otherwise, query SQS for the queue URL
	params := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}
	if resp, err := sqsService.GetQueueUrl(params); err == nil {
		queueURL = *resp.QueueUrl
	} else {
		log.Panicf("Unable to retrieve queue URL: %s", err.Error())
	}
	return
}
