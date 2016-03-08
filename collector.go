package main // import "github.com/urlgrey/streammarker-collector"

import (
	"fmt"
	stdlog "log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	kitlog "github.com/go-kit/kit/log"
	"github.com/urlgrey/streammarker-collector/binding"
	"github.com/urlgrey/streammarker-collector/config"
	"golang.org/x/net/context"
)

const (
	defaultQueueName = "streammarker-collector-messages"
)

func main() {
	// `package log` domain
	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(os.Stderr)
	logger = kitlog.NewContext(logger).With("ts", kitlog.DefaultTimestampUTC)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger)) // redirect stdlib logging to us
	stdlog.SetFlags(0)                                // flags are handled in our logger

	// read configuration from environment
	c := loadConfiguration()

	// Mechanical stuff
	rand.Seed(time.Now().UnixNano())
	root := context.Background()
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	// HTTP REST Endpoint Listeners
	binding.StartApplicationHTTPListener(logger, root, errc, c)
	binding.StartHealthCheckHTTPListener(logger, root, errc, c)

	logger.Log("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}

func loadConfiguration() *config.Configuration {
	queueName := os.Getenv("STREAMMARKER_QUEUE_NAME")
	if queueName == "" {
		queueName = defaultQueueName
	}
	sqsService := createSQSConnection()
	queueURL := findQueueURL(sqsService, queueName)

	return &config.Configuration{
		QueueName:          queueName,
		QueueURL:           queueURL,
		SQSService:         sqsService,
		ApplicationAddress: ":3000",
		HealthCheckAddress: ":3100",
	}
}

// import (
// 	stdlog "log"
// 	"os"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/sqs"
// 	"github.com/codegangsta/negroni"
// 	kitlog "github.com/go-kit/kit/log"
// 	"github.com/gorilla/mux"
// 	"github.com/urlgrey/streammarker-collector/bindings"
// 	"github.com/urlgrey/streammarker-collector/config"
// 	"github.com/urlgrey/streammarker-collector/handlers"
// )

// const (
// 	defaultQueueName = "streammarker-collector-messages"
// )

// func main() {
// 	mainServer := negroni.New()

// 	// Token auth middleware
// 	tokenVerification := handlers.NewTokenVerificationMiddleware()
// 	tokenVerification.Initialize()
// 	mainServer.Use(negroni.HandlerFunc(tokenVerification.Run))

// 	// Create external service connections
// 	sqsService := createSQSConnection()

// 	// get queue name
// 	queueName := os.Getenv("STREAMMARKER_QUEUE_NAME")
// 	if queueName == "" {
// 		queueName = defaultQueueName
// 	}

// 	// Initialize HTTP service handlers
// 	router := mux.NewRouter()
// 	queueURL := findQueueURL(sqsService, queueName)
// 	handlers.InitializeRouterForSensorReadingsHandlers(router, sqsService, queueURL)
// 	mainServer.UseHandler(router)
// 	go mainServer.Run(":3000")

// 	// `package log` domain
// 	var logger kitlog.Logger
// 	logger = kitlog.NewLogfmtLogger(os.Stderr)
// 	logger = kitlog.NewContext(logger).With("ts", kitlog.DefaultTimestampUTC)
// 	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger)) // redirect stdlib logging to us
// 	stdlog.SetFlags(0)                                // flags are handled in our logger

// 	bindings.StartHealthCheckHTTPListener(logger)
// 	// Run healthcheck service
// 	healthCheckServer := negroni.New()
// 	healthCheckRouter := mux.NewRouter()
// 	handlers.InitializeRouterForHealthCheckHandler(healthCheckRouter, sqsService, queueName)
// 	healthCheckServer.UseHandler(healthCheckRouter)
// 	healthCheckServer.Run(":3100")
// }

func createSQSConnection() *sqs.SQS {
	config := &aws.Config{}
	if endpoint := os.Getenv("STREAMMARKER_SQS_ENDPOINT"); endpoint != "" {
		config.Endpoint = &endpoint
	}

	return sqs.New(session.New(), config)
}

func findQueueURL(sqsService *sqs.SQS, queueName string) string {
	// check the environment variable first
	var queueURL string
	if queueURL = os.Getenv("STREAMMARKER_SQS_QUEUE_URL"); queueURL != "" {
		return queueURL
	}

	// otherwise, query SQS for the queue URL
	params := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}
	if resp, err := sqsService.GetQueueUrl(params); err == nil {
		queueURL = *resp.QueueUrl
	} else {
		stdlog.Panicf("Unable to retrieve queue URL: %s", err.Error())
	}
	return queueURL
}
