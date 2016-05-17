# streammarker-collector

[![wercker status](https://app.wercker.com/status/028d96b71103579daa1a9188eca1e03a/m "wercker status")](https://app.wercker.com/project/bykey/028d96b71103579daa1a9188eca1e03a) [![Go Report Card](https://goreportcard.com/badge/github.com/skidder/streammarker-collector)](https://goreportcard.com/report/github.com/skidder/streammarker-collector)

REST endpoint that collects data from data emitters and enqueues it for persistence

Posting a sample to the service:
```shell
curl -X POST -d @docs/sample_readings_put.json --header "Content-Type: application/json" -vvv http://api.streammarker.com/api/v1/sensor_readings
```

