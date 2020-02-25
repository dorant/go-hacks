package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"time"
	"os"
	"strconv"

	"contrib.go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"

	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)

var (
	Commit = "unset"
	BuildTime = "unset"
	Release = "unset"
)

func main() {
	log.Printf("Commit: %s, build time: %s, release: %s",
		Commit, BuildTime, Release)

	endpoint := os.Getenv("ZIPKIN_EP")
	if endpoint == "" {
		endpoint = "localhost:9411"
	}
	log.Printf("Using Zipkin endpoint: %s", endpoint)

	localEndpoint, err := openzipkin.NewEndpoint("go-zipkin-client", ":0")
	if err != nil {
		log.Fatalf("Failed to create the local Zipkin endpoint: %v", err)
	}
	reporter := zipkinHTTP.NewReporter("http://" + endpoint + "/api/v2/spans")
	ze := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(ze)

	// Configure 100% sample rate
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// Create a span with the background context, making this the parent span.
	{
		ctx, span := trace.StartSpan(context.Background(), "main")
		defer span.End()

		for i := 0; i < 10; i++ {
			doWork(i, ctx)
		}
	}
	time.Sleep(2000 * time.Millisecond)
	log.Printf("Done")
}

func doWork(id int, ctx context.Context) {
	// Start a child span
	_, span := trace.StartSpan(ctx, "doWork"+ strconv.Itoa(id))
	defer span.End()

	fmt.Println("doing busy work")
	time.Sleep(80 * time.Millisecond)
	buf := bytes.NewBuffer([]byte{0xFF, 0x00, 0x00, 0x00})
	num, err := binary.ReadVarint(buf)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeUnknown,
			Message: err.Error(),
		})
	}

	span.Annotate([]trace.Attribute{
		trace.Int64Attribute("bytes to int", num),
	}, "Invoking doWork")
	time.Sleep(20 * time.Millisecond)
}
