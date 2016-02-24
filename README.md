# streammarker-collector [![Circle CI](https://circleci.com/gh/urlgrey/streammarker-collector.svg?style=svg)](https://circleci.com/gh/urlgrey/streammarker-collector) [![Go Report Card](https://goreportcard.com/badge/github.com/urlgrey/streammarker-collector)](https://goreportcard.com/report/github.com/urlgrey/streammarker-collector)
REST endpoint that collects data from data emitters and enqueues it for persistence

Posting a sample to the service:
```shell
curl -X POST -d @docs/sample_readings_put.json --header "Content-Type: application/json" -vvv http://api.streammarker.com/api/v1/sensor_readings
```

