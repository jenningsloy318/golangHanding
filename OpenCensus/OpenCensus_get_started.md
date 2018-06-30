# OpenCensus with Prometheus  
OpenCensus is a Google-initiated project “that automatically collects traces and metrics from your app, displays them locally, and sends them to any analysis tool”.

I’m going to show you this working two ways using Golang with Prometheus and with Stackdriver… It’s a small cheat because the difference is effectively a single line change but this demonstrates the utility in this tool.

## Setup
```sh
PROJECT=[[YOUR-PROJECT]] # Only needed for Stackdriver
DIR="/tmp/OpenCensus"
mkdir -p ${DIR}/go
export GOPATH=${DIR}/go
export PATH=${PATH}:${GOPATH}/bin
go get -u go.opencensus.io/...
```


## OpenCensus
There’s a bug in the Golang SDK that arises when multiple measures are created. For the time being, here’s the Prometheus sample reduced to using a single measure; this is a trivial tweak to the sample so that it works:

```go
package main

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opencensus.io/exporter/stats/prometheus"
	"go.opencensus.io/stats"
)

const (
	html = `<!doctype html><html><body><a href="/metrics">metrics</a></body></html>`
)

func main() {
	ctx := context.Background()

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}
	stats.RegisterExporter(exporter)

	videoSize, err := stats.NewMeasureInt64("my.org/measures/video_size_cum", "size of processed video", "MBy")
	if err != nil {
		log.Fatalf("Video size measure not created: %v", err)
	}

	viewSize, err := stats.NewView(
		"video_cum",
		"processed video size over time",
		nil,
		videoSize,
		stats.DistributionAggregation([]float64{0, 1 << 16, 1 << 32}),
		stats.Cumulative{},
	)
	if err != nil {
		log.Fatalf("Cannot create view: %v", err)
	}

	if err := viewSize.Subscribe(); err != nil {
		log.Fatalf("Cannot subscribe to the view: %v", err)
	}

	stats.SetReportingPeriod(1 * time.Second)

	go func() {
		for {
			stats.Record(ctx, videoSize.M(rand.Int63()))
			<-time.After(time.Millisecond * time.Duration(1+rand.Intn(400)))
		}
	}()

	addr := ":9999"
	log.Printf("Serving at %s", addr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("foo").Parse(html)
		if err != nil {
			log.Fatalf("Cannot parse template: %v", err)
		}
		t.Execute(w, "")
	})
	http.Handle("/metrics", exporter)
	log.Fatal(http.ListenAndServe(addr, nil))

```


You can create this file anywhere you wish under `${DIR}/go/src`. Convention is that you create it under `github.com/[YOUR-GITHUB]` if you have one. But, if you wish, you may simply create a file called `prometheus.go` under `${DIR}/go/src`.

From within that directory:
```sh
go run [YOUR-FILENAME].go
``

Should report:

```sh
2018/01/21 16:02:37 Serving at :9999
```

and then you may curl or browse to the endpoint:

```sh
curl localhost:9999/metrics
# HELP opencensus_video_cum processed video size over time
# TYPE opencensus_video_cum histogram
opencensus_video_cum_bucket{le="0"} 0
opencensus_video_cum_bucket{le="65536"} 0
opencensus_video_cum_bucket{le="4.294967296e+09"} 0
opencensus_video_cum_bucket{le="+Inf"} 11
opencensus_video_cum_sum 3.753271169191e+19
opencensus_video_cum_count 11
```


This (metrics) data is in Prometheus’ exporter format and this means that we can point a Prometheus server to this endpoint. So, let’s create a Prometheus configuration and run a Prometheus server. Please leave your code running.
 

- Content  from: https://medium.com/google-cloud/opencensus-and-prometheus-66812a7503f



# Return to OpenCensus

This post updates [OpenCensus with Prometheus](./OpenCensus_with_Prometheus.md). One of the software engineers on the OpenCensus team pinged me with a couple of updates on the Golang library, yay! Thanks JBD!

The OpenCensus Golang API has been refined (“stats” is split into 2 packages), the bug is resolved and merged (thanks again JBD) and the Golang SDK works with multiple measures.

I won’t duplicate the content from the previous post, so here are the headlines. My samples had minor changes to the published samples because, at that time, the published samples didn’t work with multiple measures.

## Prometheus
The sample published in the GitHub repo now works. For consistency, here’s my copy of it:

```go
package main

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

const (
	html = `<!doctype html>
<html>
<body>
	<a href="/metrics">metrics</a>
</body>
</html>`
)

func main() {
	ctx := context.Background()

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}
	view.RegisterExporter(exporter)

	videoCount, err := stats.Int64("my.org/measures/video_count", "number of processed videos", "")
	if err != nil {
		log.Fatalf("Video count measure not created: %v", err)
	}

	viewCount, err := view.New(
		"video_count",
		"number of videos processed over time",
		nil,
		videoCount,
		view.CountAggregation{},
	)
	if err != nil {
		log.Fatalf("Cannot create view: %v", err)
	}

	if err := viewCount.Subscribe(); err != nil {
		log.Fatalf("Cannot subscribe to view: %v", err)
	}

	videoSize, err := stats.Int64("my.org/measures/video_size_cum", "size of processed video", "MBy")
	if err != nil {
		log.Fatalf("Video size measure not created: %v", err)
	}

	viewSize, err := view.New(
		"video_cum",
		"processed video size over time",
		nil,
		videoSize,
		view.DistributionAggregation([]float64{0, 1 << 16, 1 << 32}),
	)
	if err != nil {
		log.Fatalf("Cannot create view: %v", err)
	}

	if err := viewSize.Subscribe(); err != nil {
		log.Fatalf("Cannot subscribe to the view: %v", err)
	}

	view.SetReportingPeriod(1 * time.Second)

	go func() {
		for {
			stats.Record(ctx, videoCount.M(1))
			stats.Record(ctx, videoSize.M(rand.Int63()))
			<-time.After(time.Millisecond * time.Duration(1+rand.Intn(400)))
		}
	}()

	addr := ":9999"
	log.Printf("Serving at %s", addr)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("foo").Parse(html)
		if err != nil {
			log.Fatalf("Cannot parse template: %v", err)
		}
		t.Execute(w, "")
	})
	http.Handle("/metrics", exporter)
	log.Fatal(http.ListenAndServe(addr, nil))
}
```


- Content from https://medium.com/@DazWilkin/return-to-opencensus-42623f1b55b8