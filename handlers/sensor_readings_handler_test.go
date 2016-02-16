package handlers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urlgrey/streammarker-collector/mocks"
)

type SensorReadingsHandlerSuite struct {
	suite.Suite
	handler        *SensorReadingsHandler
	sqsAPI         *mocks.SQSAPI
	request        *http.Request
	responseWriter *http.ResponseWriter
}

func (s *SensorReadingsHandlerSuite) SetupTest() {
	s.request = new(http.Request)
	s.responseWriter = new(http.ResponseWriter)
	s.sqsAPI = new(mocks.SQSAPI)
	s.handler = NewSensorReadingsHandler(s.sqsAPI, "queueURL")
}

func (s *SensorReadingsHandlerSuite) TestAllGoodSingleReading() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusCreated, rec.Code)
}

func (s *SensorReadingsHandlerSuite) TestSingleReadingFailSendingToSQS() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(nil, errors.New("Fake Error"))
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusInternalServerError, rec.Code)
}

func (s *SensorReadingsHandlerSuite) TestNoTimestampSingleReading() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading_no_timestamp.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusCreated, rec.Code)
}

func (s *SensorReadingsHandlerSuite) TestNoReadingTimestampSingleReading() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading_no_reading_timestamp.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusCreated, rec.Code)
}

func (s *SensorReadingsHandlerSuite) TestAllGoodMultipleReadings() {
	// GIVEN
	s.NotNil(s.sqsAPI)
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_multiple_readings.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusCreated, rec.Code)
}

func (s *SensorReadingsHandlerSuite) TestNoContentType() {
	// GIVEN
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading.json"); err != nil {
		s.NoError(err)
	}
	r, _ := http.NewRequest("POST", "/api/v1/sensor_readings", bytes.NewReader(requestData))
	r.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	rw := negroni.NewResponseWriter(rec)
	rw.Header().Set("Content-Type", "application/json")

	// WHEN
	s.handler.SubmitMessage(rw, r)

	// THEN
	s.Equal(http.StatusUnsupportedMediaType, rec.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSensorReadingsHandlerSuite(t *testing.T) {
	suite.Run(t, new(SensorReadingsHandlerSuite))
}
