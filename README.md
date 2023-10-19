# Alert

## About

Enables users to send and read alerts using a REST api with JSON

## Running the service

1.  Tidy the project using `go mod tidy` commmand so that all necessary packages for this project are installed

2.  To run the service, go inside the project root folder and run: `go run main.go`

3.  Send all the requests you want to send

4.  Stop the service by pressing CTRL+C

## Writing a alert

Here is a cURL command you can run from your terminal to send a request to the service:

```shell
curl -X POST -H "Content-Type: application/json" -d '{
  "alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
  "service_id": "my_test_service_id",
  "service_name": "my_test_service",
  "model": "my_test_model",
  "alert_type": "anomaly",
  "alert_ts": "1695644160",
  "severity": "warning",
  "team_slack": "slack_ch"
}' http://localhost:8080
```

## Reading alerts

Here is a cURL command you can run from your terminal to get all allerts that belong to a service:
Query consists of 3 parameters:

- service_id: ID of the service you are trying to get alerts for
- start_ts: an int timestamp
- end_ts: an int timestamp

All retreived alerts are in between the start_ts and end_ts timestamps. If no alerts are found in between these times the alerts array will be empty.

```shell
curl "http://localhost:8080/alerts?service_id=dy_test_service_id&start_ts=1795643160&end_ts=1795644360"
```
