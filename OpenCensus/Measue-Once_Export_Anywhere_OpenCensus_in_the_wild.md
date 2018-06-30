# Measure Once — Export Anywhere: OpenCensus in the wild

A few months ago, Google has announced OpenCensus, a vendor-neutral open source library for telemetry and tracing collection. OpenCensus makes it easy to collect metrics from your app and to trace the progression or requests.

In most of the cases, there is a clear distinction between collecting the data and exporting it to tracing and monitoring systems. Once you have started to use OpenCensus in your project and realize what metrics are particularly interesting to you to trace and monitor, you can switch between different monitoring systems or even use more than one without making any changes to your metric collection logic.

We, at DoiT International, are in the process of developing a reference architecture and actual implementation of event analytics pipeline (still work in progress, but you can have a sneak peek here).

We are writing our API in Go and we wanted to give the users some visibility into the system and to some custom metrics we’d like to expose. In the “old days” you had to choose your monitoring system, and if you wanted to replace it with something else, you would have to invest some time (and money). But what if you are building an open source system and you don’t know upfront where and how it’s going to be deployed and how it would be monitored?

With OpenCensus, you don’t need to decide almost anything upfront. The selection of the monitoring system is not tightly coupled and you can even use more than one monitoring or metric collection system.

So, in order to expose the metrics to you monitoring system of choice, all you’ll need to do, is to create a new exporter and register it:

```go

import (
  	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"

)
Exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		logger.Error("Error creating prometheus exporter  ", zap.Error(err))
	}
	// Export to Prometheus Monitoring.
	view.RegisterExporter(pExporter)
	sExporter, err := stackdriver.NewExporter(stackdriver.Options{ProjectID: config.ProjectID})
	if err != nil {
		logger.Error("Error creating stackdriver exporter  ", zap.Error(err))
	}
	// Export to Stackdriver Monitoring.
	view.RegisterExporter(sExporter)
```

Creating metrics is also very easy, you’ll just need to decide on their type and about how you’d like them to be displayed:


```go
import (
  "go.opencensus.io/stats"
  "go.opencensus.io/tag"
  "go.opencensus.io/stats/view"
)

var (
  requestCounter             *stats.Float64Measure
  requestlatency             *stats.Float64Measure
  codeKey                    tag.Key
  DefaultLatencyDistribution = view.DistributionAggregation{0, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000}
)
	codeKey, _ = tag.NewKey("banias/keys/code")
	requestCounter, _ = stats.Float64("banias/measures/request_count", "Count of HTTP requests processed", stats.UnitNone)
	requestlatency, _ = stats.Float64("banias/measures/request_latency", "Latency distribution of HTTP requests", stats.UnitMilliseconds)
	view.Subscribe(
		&view.View{
			Name:        "request_count",
			Description: "Count of HTTP requests processed",
			TagKeys:     []tag.Key{codeKey},
			Measure:     requestCounter,
			Aggregation: view.CountAggregation{},
		})
	view.Subscribe(
		&view.View{
			Name:        "request_latency",
			Description: "Latency distribution of HTTP requests",
			TagKeys:     []tag.Key{codeKey},
			Measure:     requestlatency,
			Aggregation: DefaultLatencyDistribution,
		})

	view.SetReportingPeriod(1 * time.Second)
```

Now you are ready to collect the metrics, just call the Record method and provide the required data.

```go
func (c *Collector) Collect(ctx *fasthttp.RequestCtx) {
  defer func(begin time.Time) {
      responseTime := float64(time.Since(begin).Nanoseconds() / 1000)
      occtx, _ := tag.New(context.Background(), tag.Insert(codeKey, strconv.Itoa(ctx.Response.StatusCode())), )
      stats.Record(occtx, requestCounter.M(1))
      stats.Record(occtx, requestlatency.M(responseTime))
    }(time.Now())
/*do some stuff */
}
```



- Content from https://blog.doit-intl.com/measure-once-export-anywhere-opencensus-in-the-wild-61724f44eb00