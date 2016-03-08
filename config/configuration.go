package config

import (
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// Configuration holds application configuration details
type Configuration struct {
	ApplicationAddress string
	HealthCheckAddress string
	SQSService         sqsiface.SQSAPI
	QueueName          string
	QueueURL           string
}
