package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
)

type SensorReadingsHandler struct {
	sqsService sqsiface.SQSAPI
	queueURL   string
}

func NewSensorReadingsHandler(sqsService sqsiface.SQSAPI, queueURL string) *SensorReadingsHandler {
	return &SensorReadingsHandler{sqsService: sqsService, queueURL: queueURL}
}

func InitializeRouterForSensorReadingsHandlers(r *mux.Router, sqsService sqsiface.SQSAPI, queueURL string) {
	m := NewSensorReadingsHandler(sqsService, queueURL)
	r.HandleFunc("/api/v1/sensor_readings", m.SubmitMessage).Methods("POST")
}

func (m *SensorReadingsHandler) SubmitMessage(resp http.ResponseWriter, req *http.Request) {
	message := new(MeasurementMessage)
	errs := binding.Bind(req, message)
	if errs.Handle(resp) {
		log.Printf("Error while binding request to model: %s", errs.Error())
		return
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

			if queueMessageJson, err := json.Marshal(queueMessage); err != nil {
				log.Printf("Error serializing parsed message for queueing: %s", err.Error())
				http.Error(resp,
					"Error serializing APRS message for queueing",
					http.StatusInternalServerError)
				return
			} else {
				params := &sqs.SendMessageInput{
					MessageBody: aws.String(string(queueMessageJson)),
					QueueUrl:    aws.String(m.queueURL),
				}
				if _, err = m.sqsService.SendMessage(params); err != nil {
					log.Printf("Error sending message to queue: %s", err.Error())
					http.Error(resp,
						"Error sending message to queue",
						http.StatusInternalServerError)
					return
				}
			}
		}
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	responseEncoder := json.NewEncoder(resp)
	responseEncoder.Encode("{}")
}

type MeasurementMessage struct {
	Timestamp int32    `json:"timestamp"`
	RelayID   string   `json:"relay_id"`
	Status    string   `json:"status"`
	Sensors   []Sensor `json:"sensors"`
}

func (m *MeasurementMessage) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&m.Timestamp: "timestamp",
		&m.RelayID:   "relay_id",
		&m.Status:    "status",
		&m.Sensors:   "sensors",
	}
}

type Sensor struct {
	ID             string          `json:"id"`
	SensorReadings []SensorReading `json:"readings"`
}

func (s *Sensor) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.ID:             "id",
		&s.SensorReadings: "readings",
	}
}

type SensorReading struct {
	Timestamp    int32         `json:"timestamp"`
	Measurements []Measurement `json:"measurements"`
}

func (s *SensorReading) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.Timestamp:    "timestamp",
		&s.Measurements: "measurements",
	}
}

type Measurement struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

func (s *Measurement) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.Name:  "name",
		&s.Value: "value",
	}
}

type SensorReadingQueueMessage struct {
	RelayID            string        `json:"relay_id"`
	SensorID           string        `json:"sensor_id"`
	ReadingTimestamp   int32         `json:"reading_timestamp"`
	ReportingTimestamp int32         `json:"reporting_timestamp"`
	Measurements       []Measurement `json:"measurements"`
}
