# Test of Prometheus client

## Run
> go run server.go

Use service:
> curl localhost:8080/hash -d "test"

Scrape metrics:
> curl localhost:8080/metrics


## Info
* Prometheus is a pull-based system,

Some best practices: https://prometheus.io/docs/practices/instrumentation/
* Use labels: rather than http_responses_500_total and http_responses_403_total,
  create a single metric called http_responses_total with a code label for the HTTP response code.
  Like:
    NewCounterVec creates a new CounterVec based on the provided CounterOpts and
    partitioned by the given label names.


PromQL:
https://prometheus.io/docs/prometheus/latest/querying/basics/

https://www.digitalocean.com/community/tutorials/how-to-query-prometheus-on-ubuntu-14-04-part-1
