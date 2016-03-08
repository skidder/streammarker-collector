package binding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	levlog "github.com/go-kit/kit/log/levels"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/urlgrey/streammarker-collector/config"
	"github.com/urlgrey/streammarker-collector/endpoint"

	"golang.org/x/net/context"
)

const (
	optionsHTTPMethod = "OPTIONS"
	getHTTPMethod     = "GET"
	postHTTPMethod    = "POST"
)

// StartApplicationHTTPListener creates a Go-routine that has an HTTP listener sensor readings
func StartApplicationHTTPListener(logger kitlog.Logger, root context.Context, errc chan error, c *config.Configuration) {
	go func() {
		ctx, cancel := context.WithCancel(root)
		defer cancel()

		l := levlog.New(logger)
		l.Info().Log("ApplicationAddress", c.ApplicationAddress, "transport", "HTTP/JSON")

		router := createApplicationRouter(ctx, endpoint.NewSensorReadingServicer(c))
		errc <- http.ListenAndServe(c.ApplicationAddress, router)
	}()
}

func createApplicationRouter(ctx context.Context, sensorReadingsServicer endpoint.SensorReadingsServicer) *mux.Router {
	apiTokens := strings.Split(os.Getenv("STREAMMARKER_COLLECTOR_API_TOKENS"), ",")
	router := mux.NewRouter()
	router.Handle("/api/v1/sensor_readings",
		kithttp.NewServer(
			ctx,
			endpoint.VerifyAPIKey(apiTokens)(sensorReadingsServicer.Run),
			decodeSensorReadingsHTTPRequest,
			encodeSensorReadingsHTTPResponse,
			kithttp.ServerErrorEncoder(errorEncoder),
		)).Methods(postHTTPMethod)
	return router
}

// StartHealthCheckHTTPListener creates a Go-routine that has an HTTP listener for the healthcheck endpoint
func StartHealthCheckHTTPListener(logger kitlog.Logger, root context.Context, errc chan error, c *config.Configuration) {
	go func() {
		ctx, cancel := context.WithCancel(root)
		defer cancel()

		l := levlog.New(logger)
		l.Info().Log("HealthCheckAddress", c.HealthCheckAddress, "transport", "HTTP/JSON")

		router := createHealthCheckRouter(ctx, endpoint.NewHealthCheck(c))
		errc <- http.ListenAndServe(c.HealthCheckAddress, router)
	}()
}

func createHealthCheckRouter(ctx context.Context, healthCheckEndpoint endpoint.HealthCheckServicer) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/healthcheck",
		kithttp.NewServer(
			ctx,
			healthCheckEndpoint.Run,
			func(*http.Request) (interface{}, error) { return struct{}{}, nil },
			encodeHealthCheckHTTPResponse,
		)).Methods(getHTTPMethod)
	return router
}

func errorEncoder(w http.ResponseWriter, err error) {
	switch err {
	case endpoint.ErrTokenVerificationFailure:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func encodeHealthCheckHTTPResponse(w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(i.(*endpoint.HealthCheckResponse))
}

func decodeSensorReadingsHTTPRequest(r *http.Request) (interface{}, error) {
	message := new(endpoint.MeasurementMessage)
	errs := binding.Bind(r, message)
	if errs != nil {
		return nil, fmt.Errorf("Error while binding request to model: %s", errs.Error())
	}
	message.APIToken = r.Header.Get("X-API-KEY")
	return message, nil
}

func encodeSensorReadingsHTTPResponse(w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(i.(*endpoint.ReadingResponse))
}
