package binding

import (
	"encoding/json"
	"fmt"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	levlog "github.com/go-kit/kit/log/levels"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"github.com/skidder/streammarker-collector/config"
	"github.com/skidder/streammarker-collector/endpoint"

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

		router := createApplicationRouter(logger, ctx, c, endpoint.NewSensorReadingServicer(c))
		errc <- http.ListenAndServe(c.ApplicationAddress, router)
	}()
}

func createApplicationRouter(logger kitlog.Logger, ctx context.Context, c *config.Configuration, sensorReadingsServicer endpoint.SensorReadingsServicer) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/api/v1/sensor_readings",
		kithttp.NewServer(
			ctx,
			endpoint.VerifyAPIKey(c.APITokens)(sensorReadingsServicer.HandleMeasurementMessage),
			decodeSensorReadingsHTTPRequest,
			encodeSensorReadingsHTTPResponse,
			kithttp.ServerErrorEncoder(errorEncoder),
			kithttp.ServerErrorLogger(logger),
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

		router := createHealthCheckRouter(logger, ctx, endpoint.NewHealthCheck(c))
		errc <- http.ListenAndServe(c.HealthCheckAddress, router)
	}()
}

func createHealthCheckRouter(logger kitlog.Logger, ctx context.Context, healthCheckEndpoint endpoint.HealthCheckServicer) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/healthcheck",
		kithttp.NewServer(
			ctx,
			healthCheckEndpoint.Run,
			func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
			encodeHealthCheckHTTPResponse,
			kithttp.ServerErrorEncoder(errorEncoder),
			kithttp.ServerErrorLogger(logger),
		)).Methods(getHTTPMethod)
	return router
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	e, ok := err.(kithttp.Error)
	if ok {
		switch e.Err {
		case endpoint.ErrTokenVerificationFailure:
			w.WriteHeader(http.StatusUnauthorized)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeHealthCheckHTTPResponse(ctx context.Context, w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(i.(*endpoint.HealthCheckResponse))
}

func decodeSensorReadingsHTTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	message := new(endpoint.MeasurementMessage)
	errs := binding.Bind(r, message)
	if errs != nil {
		return nil, fmt.Errorf("Error while binding request to model: %s", errs.Error())
	}
	message.APIToken = r.Header.Get("X-API-KEY")
	return message, nil
}

func encodeSensorReadingsHTTPResponse(ctx context.Context, w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(i.(*endpoint.ReadingResponse))
}
