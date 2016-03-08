package endpoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urlgrey/streammarker-collector/config"
	"github.com/urlgrey/streammarker-collector/mocks"

	"golang.org/x/net/context"
)

type SensorReadingsTestSuite struct {
	suite.Suite
	sqsAPI                 *mocks.SQSAPI
	sensorReadingsServicer SensorReadingsServicer
}

func (s *SensorReadingsTestSuite) SetupSuite() {
}

func (s *SensorReadingsTestSuite) TearDownSuite() {
}

func (s *SensorReadingsTestSuite) SetupTest() {
	s.sqsAPI = new(mocks.SQSAPI)
	c := &config.Configuration{
		SQSService: s.sqsAPI,
		QueueURL:   "queueURL",
	}
	s.sensorReadingsServicer = NewSensorReadingServicer(c)
}

func (s *SensorReadingsTestSuite) TearDownTest() {
}

func (s *SensorReadingsTestSuite) TestRecordReadings_SingleCompleteReading() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading.json"); err != nil {
		s.NoError(err)
	}
	message := new(MeasurementMessage)
	err = json.NewDecoder(bytes.NewReader(requestData)).Decode(message)
	if err != nil {
		s.Fail("Error while binding request to model", err.Error())
	}

	// WHEN
	response, err := s.sensorReadingsServicer.Run(context.Background(), message)

	// THEN
	readingResponse := response.(*ReadingResponse)
	s.Equal("OK", readingResponse.Status)
	s.NoError(err)
	s.sqsAPI.AssertExpectations(s.T())
}

func (s *SensorReadingsTestSuite) TestRecordReadings_NoTimestamp() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading_no_reading_timestamp.json"); err != nil {
		s.NoError(err)
	}
	message := new(MeasurementMessage)
	err = json.NewDecoder(bytes.NewReader(requestData)).Decode(message)
	if err != nil {
		s.Fail("Error while binding request to model", err.Error())
	}

	// WHEN
	response, err := s.sensorReadingsServicer.Run(context.Background(), message)

	// THEN
	readingResponse := response.(*ReadingResponse)
	s.Equal("OK", readingResponse.Status)
	s.NoError(err)
	s.sqsAPI.AssertExpectations(s.T())
}

func (s *SensorReadingsTestSuite) TestRecordReadings_MultipleReadings() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(new(sqs.SendMessageOutput), nil)
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_multiple_readings.json"); err != nil {
		s.NoError(err)
	}
	message := new(MeasurementMessage)
	err = json.NewDecoder(bytes.NewReader(requestData)).Decode(message)
	if err != nil {
		s.Fail("Error while binding request to model", err.Error())
	}

	// WHEN
	response, err := s.sensorReadingsServicer.Run(context.Background(), message)

	// THEN
	readingResponse := response.(*ReadingResponse)
	s.Equal("OK", readingResponse.Status)
	s.NoError(err)
	s.sqsAPI.AssertExpectations(s.T())
}

func (s *SensorReadingsTestSuite) TestRecordReadings_SQSError() {
	// GIVEN
	s.sqsAPI.On("SendMessage", mock.AnythingOfType("*sqs.SendMessageInput")).Return(nil, errors.New("Fake SQS error"))
	var requestData []byte
	var err error
	if requestData, err = ioutil.ReadFile("fixtures/success_single_reading.json"); err != nil {
		s.NoError(err)
	}
	message := new(MeasurementMessage)
	err = json.NewDecoder(bytes.NewReader(requestData)).Decode(message)
	if err != nil {
		s.Fail("Error while binding request to model", err.Error())
	}

	// WHEN
	response, err := s.sensorReadingsServicer.Run(context.Background(), message)

	// THEN
	s.Nil(response)
	s.Error(err)
	s.sqsAPI.AssertExpectations(s.T())

}

func TestSensorReadingsTestSuite(t *testing.T) {
	suite.Run(t, new(SensorReadingsTestSuite))
}
