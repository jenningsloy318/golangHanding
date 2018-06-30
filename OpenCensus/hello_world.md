# Hello world with opencensus

In Go, this is ordinarily what “Hello, world!” with an HTTP server would look like
```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
```

If you then run it by go run main.go you can visit it at URL http://localhost:8080/hello or say with cURL

```sh
curl -i http://localhost:8080/hello
HTTP/1.1 200 OK
Date: Wed, 18 Apr 2018 10:08:41 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello world!
```


Given the above web server, we can do a whole lot more and sprinkle the appropriate tools to add insights to our microservices with distributed tracing and monitoring, beginning with this simple example below:
```go
package main

import (
	"log"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloWorld)
	h := &ochttp.Handler{Handler: mux}
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatal("Failed to register ochttp.DefaultServerViews")
	}
	log.Fatal(http.ListenAndServe(":8080", h))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

```


The above is sufficient for your microservice to propagate traces and metrics through requests that it receives.

However, to succinctly consume the fruits of tracing and monitoring, a picture usually being worth a thousand words — these visuals we can get from a few exporters that OpenCensus provides out of the box, such as:

By modifying the vanilla-opencensus.go to the diff below

```sh
--- vanilla.go	2018-04-18 03:12:50.000000000 -0700
+++ opencensus-enabled.go	2018-04-18 04:30:26.000000000 -0700
@@ -3,14 +3,51 @@
 import (
 	"log"
 	"net/http"
+	"os"
 
+	xray "github.com/census-instrumentation/opencensus-go-exporter-aws"
+	"go.opencensus.io/exporter/prometheus"
+	"go.opencensus.io/exporter/stackdriver"
 	"go.opencensus.io/plugin/ochttp"
 	"go.opencensus.io/stats/view"
+	"go.opencensus.io/trace"
 )
 
 func main() {
+	// 1. Create the exporters
+	se, err := stackdriver.NewExporter(stackdriver.Options{
+		ProjectID: os.Getenv("STACKDRIVER_PROJECT_ID"),
+	})
+	if err != nil {
+		log.Fatalf("Failed to create Stackdriver exporter: %v", err)
+	}
+	trace.RegisterExporter(se)
+	view.RegisterExporter(se)
+
+	pe, err := prometheus.NewExporter(prometheus.Options{
+		Namespace: "hellohttp",
+	})
+	if err != nil {
+		log.Fatalf("Failed to create Prometheus exporter: %v", err)
+	}
+	view.RegisterExporter(pe)
+
+	xe, err := xray.NewExporter(xray.WithVersion("latest"))
+	if err != nil {
+		log.Fatalf("Failed to create X-Ray exporter: %v", err)
+	}
+	trace.RegisterExporter(xe)
+
+	// 2. Set the tracer
+	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
+
+	// 3. Create the handlers
 	mux := http.NewServeMux()
 	mux.HandleFunc("/hello", helloWorld)
+
+	// Ensure that the Prometheus endpoint is exposed for scraping
+	mux.Handle("/metrics", pe)
+
 	h := &ochttp.Handler{Handler: mux}
 	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
 		log.Fatal("Failed to register ochttp.DefaultServerViews")
```

Which then becomes this file

```go
package main

import (
	"log"
	"net/http"
	"os"

	xray "github.com/census-instrumentation/opencensus-go-exporter-aws"
	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func main() {
	// 1. Create the exporters
	se, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("STACKDRIVER_PROJECT_ID"),
	})
	if err != nil {
		log.Fatalf("Failed to create Stackdriver exporter: %v", err)
	}
	trace.RegisterExporter(se)
	view.RegisterExporter(se)

	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "hellohttp",
	})
	if err != nil {
		log.Fatalf("Failed to create Prometheus exporter: %v", err)
	}
	view.RegisterExporter(pe)

	xe, err := xray.NewExporter(xray.WithVersion("latest"))
	if err != nil {
		log.Fatalf("Failed to create X-Ray exporter: %v", err)
	}
	trace.RegisterExporter(xe)

	// 2. Set the tracer
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// 3. Create the handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloWorld)

	// Ensure that the Prometheus endpoint is exposed for scraping
	mux.Handle("/metrics", pe)

	h := &ochttp.Handler{Handler: mux}
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatal("Failed to register ochttp.DefaultServerViews")
	}
	log.Fatal(http.ListenAndServe(":8080", h))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world HTTP!"))
}
```


In order to run the example above, you’ll need:

* AWS credentials and in particular in your environment, please set the following environment variables AWS_SECRET_ACCESS_KEY AWS_ACCESS_KEY_ID AWS_REGION
* Google Cloud Platform credentials, if you don’t have those, you can create a free project at https://cloud.google.com/go/home but the environment variable in question is GOOGLE_APPLICATION_CREDENTIALS
* A Prometheus installed and able to access your demo server. 
