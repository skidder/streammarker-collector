package endpoint

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/mholt/binding"
	"github.com/urlgrey/streammarker-collector/config"

	"golang.org/x/net/context"
)

// SensorReadingsServicer provides functions for dealing with new sensor readings
type SensorReadingsServicer interface {
	Run(context.Context, interface{}) (interface{}, error)
}

type sensorReadingService struct {
	sqsService sqsiface.SQSAPI
	queueURL   string
	queueName  string
}

// NewSensorReadingServicer creates a new healthcheck
func NewSensorReadingServicer(c *config.Configuration) SensorReadingsServicer {
	return &sensorReadingService{c.SQSService, c.QueueURL, c.QueueName}
}

func (s *sensorReadingService) Run(ctx context.Context, i interface{}) (response interface{}, err error) {
	message, ok := i.(*MeasurementMessage)
	if !ok {
		return nil, kitendpoint.ErrBadCast
	}

	for _, sensor := range message.Sensors {
		for _, reading := range sensor.SensorReadings {
			// get timestamp for reading
			var readingTimestamp, reportingTimestamp int32
			if reading.Timestamp <= 0 {
				if message.Timestamp <= 0 {
					readingTimestamp = int32(time.Now().Unix())
					reportingTimestamp = readingTimestamp
				} else {
					readingTimestamp = message.Timestamp
				}
				reportingTimestamp = readingTimestamp
			} else {
				readingTimestamp = reading.Timestamp
				if reportingTimestamp <= 0 {
					reportingTimestamp = readingTimestamp
				} else {
					reportingTimestamp = message.Timestamp
				}
			}

			queueMessage := &SensorReadingQueueMessage{
				RelayID:            message.RelayID,
				SensorID:           sensor.ID,
				ReadingTimestamp:   readingTimestamp,
				ReportingTimestamp: reportingTimestamp,
				Measurements:       reading.Measurements,
			}
			queueMessageJSON, err := json.Marshal(queueMessage)
			if err != nil {
				return nil, fmt.Errorf("Error serializing parsed message for queueing: %s", err.Error())
			}

			params := &sqs.SendMessageInput{
				MessageBody: aws.String(string(queueMessageJSON)),
				QueueUrl:    aws.String(s.queueURL),
			}
			if _, err = s.sqsService.SendMessage(params); err != nil {
				return nil, fmt.Errorf("Error sending message to queue: %s", err.Error())
			}
		}
	}
	return &ReadingResponse{"OK"}, nil
}

// ReadingResponse has fields with operation status
type ReadingResponse struct {
	Status string `json:"status"`
}

// MeasurementMessage represents a message to be enqueued for later processing
type MeasurementMessage struct {
	Timestamp int32    `json:"timestamp"`
	RelayID   string   `json:"relay_id"`
	Status    string   `json:"status"`
	Sensors   []Sensor `json:"sensors"`
	APIToken  string
}

// FieldMap binds fields to their JSON labels
func (m *MeasurementMessage) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&m.Timestamp: "timestamp",
		&m.RelayID:   "relay_id",
		&m.Status:    "status",
		&m.Sensors:   "sensors",
	}
}

// GetToken returns the API token associated with the request
func (m *MeasurementMessage) GetToken() string {
	return m.APIToken
}

// Sensor represents a sensor with a collection of readings to be saved
type Sensor struct {
	ID             string          `json:"id"`
	SensorReadings []SensorReading `json:"readings"`
}

// FieldMap binds fields to their JSON labels
func (s *Sensor) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.ID:             "id",
		&s.SensorReadings: "readings",
	}
}

// SensorReading represents a reading from a point in time with one or more measurements
type SensorReading struct {
	Timestamp    int32         `json:"timestamp"`
	Measurements []Measurement `json:"measurements"`
}

// FieldMap binds fields to their JSON labels
func (s *SensorReading) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.Timestamp:    "timestamp",
		&s.Measurements: "measurements",
	}
}

// Measurement contains the measurement details
type Measurement struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// FieldMap binds fields to their JSON labels
func (s *Measurement) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.Name:  "name",
		&s.Value: "value",
	}
}

// SensorReadingQueueMessage holds one or more measurements to be written
type SensorReadingQueueMessage struct {
	RelayID            string        `json:"relay_id"`
	SensorID           string        `json:"sensor_id"`
	ReadingTimestamp   int32         `json:"reading_timestamp"`
	ReportingTimestamp int32         `json:"reporting_timestamp"`
	Measurements       []Measurement `json:"measurements"`
}
