# OpenCensus and Go database/sql

I still remember that blissful time when we had our monolith. Behavior of our platform was easy to monitor and reason about. With the ongoing success of our platform and ever increasing amount of services offered to our customers we needed to expand and accelerate our development cycles. Our monolithic platform started to show the boundaries of what could be achieved with it. To allow for internal and external teams to work in parallel on features for us, we made the decision to transition our platform to microservices.

Behavior and performance characteristics changed when we moved to microservices. Unfortunately the amount of insight we had when dealing with our monolith was dramatically lowered by the move. At the same time our need for insight continued to increase and made it even more important then ever before. Faster iterations of features and services required more planning, more reasoning, better understanding of our platform as a whole; the opposite of what we initially achieved! While planning the migration we read up on microservices observability and decided on a combination of logs, metrics and distributed tracing. Together these have been vital to regaining the lost insight.

Most of our microservices are built with Go kit which by its flexibility has allowed us to gradually transition away from our legacy monolith and plays nicely with the various transports we use (REST, JSON-RPC, gRPC). While Prometheus metrics and Zipkin distributed tracing backends have given us back our insights for the largest part, we still had an issue with the changed behavior of our databases. Providing live results with custom searches for large scale participatory sports events is a challenge and making fast mutating data reliably accessible to large audiences with extreme request spikes during these races is really hard. Where our service to service interactions and handled API calls had become clearly visible, we noticed that we didn’t quite grasp yet how our database usage characteristics changed due to the microservices architecture. The distribution and timing of individual SQL queries altered and correlating this information with the rest of our service interactions showed us we hit a blind spot. In some of our earlier services we manually instrumented around our database calls to eliminate the issue, but this obviously is error prone and tedious.

Enter ocsql, a Go database/sql driver wrapper, which bridges the gap between service to service and service to database interactions. It seamlessly works together with other Go OpenCensus instrumentation like ochttp and ocgrpc and provides an “out of the way” experience for your microservice developers, allowing them to focus on business logic instead.

## context, context, context.
Before diving into ocsql and how to use it in your code I want to provide some notes on context.

When context is not coherent and you don’t provide and pass along the right context, you can’t attribute cause, which is the very thing you are trying to understand. You end up with useless fragmentation of instrumented data and senseless metrics. How do the pieces coming from different components and even services all fit back together? What did we signal and measure exactly?

>The three most important considerations in aggregation and correlation of observability data are context, context, context!

Unfortunately the word context itself requires some explanation to untangle some confusion; especially in Gopher land. In the distributed tracing world we often speak about context in which we refer to the set of details that identifies a span (a timed and annotated unit of work) and allows for establishing (parent-child) relationships with other spans. Together they make up a trace (which covers a request in its entirety across all services it touches). For traces to work we need to propagate these details inside our processes as well as between services when doing RPC’s. Let’s call these details `SpanContext` from now on.

When Gophers talk about context they often mean the context package and its Context type. I’ll refer to it as `Go-context` and it can carry deadlines, cancellation signals and other request scoped values across API boundaries and between processes. Go-context is the perfect vehicle for transporting our SpanContext where it needs to go and is therefore used extensively by OpenCensus.

Maybe because both are referred to as context and they are great companions, they often get conflated in discussions. As the code from opencensus-go shows below it is explicitly documented that SpanContext is not an implementation of context.Context.

```go
// SpanContext contains the state that must propagate across process boundaries.
//
// SpanContext is not an implementation of context.Context.
// TODO: add reference to external Census docs for SpanContext.
type SpanContext struct {
	TraceID      TraceID
	SpanID       SpanID
	TraceOptions TraceOptions
}
```

If unfamiliar with Go-context I advise you to read up on it as it is an important building block for keeping microservice architectures written in Go manageable and it is fundamental to instrumenting your code with OpenCensus. You can start with the package’s godoc and the Go blog article on Context.

## in-process propagation

Step one is passing along the Go-context to all functions on the call path between incoming and outgoing requests. This is vital for your microservices to enable handling request deadlines, cancellations and SpanContext propagation for distributed tracing. It is idiomatic to provide Go-context as first parameter to these functions. We should treat interactions between services and databases as being on the call path too. Executing a SQL query is an outgoing request just like a REST HTTP Post or gRPC method invocation is.

Let’s say you don’t pass context properly or forgot to pass it along somewhere in your call path, you will notice traces being incomplete and fragmented. When context is broken your traces can mislead you. It will reverse the intent of observability which ultimately is understanding your services!

OpenCensus explicitly uses Go-context to pass along the SpanContext on the call path. Because the authors of OpenCensus understand that SpanContext propagation is the most important piece in the chain, they have made the use of Go-context very explicit when creating spans. Instead of solely relying on helper functions moving spans into and out of Go-context they have added Go-context as input and output parameters to the span creation functions.

```go
// StartSpan starts a new child span of the current span in the context. If
// there is no span in the context, creates a new trace and span.
func StartSpan(
  ctx context.Context, name string, o ...StartOption,
) (context.Context, *Span) {
  ...
}

// StartSpanWithRemoteParent starts a new child span of the span from the given parent.
//
// If the incoming context contains a parent, it ignores. StartSpanWithRemoteParent is
// preferred for cases where the parent is propagated via an incoming request.
func StartSpanWithRemoteParent(
  ctx context.Context, name string, parent SpanContext, o ...StartOption,
) (context.Context, *Span) {
  ...
}
```
Now it’s up to you to make sure this Go-context is passed in all functions on the call path as this little example shows below:

```go
// GetDetails is a standard HTTP handler function
func (s *service) GetDetails(w http.ResponseWriter, r *http.Request) {
  // Each incoming HTTP Request holds a Go-context object.
  // Assume we are using ochttp to instrument our HTTP server,
  // so the Go-context provided in the http.Request holds a server
  // side endpoint span which this method is servicing.
  // We'll be passing it to all functions on the call path.
  ctx := r.Context()
  
  // Let's assume all is good and we extracted a userID to use.
  // When we call the repository method, we pass along the Go-context
  // as we found it in the HTTP request.
  details, err := s.repository.GetUserDetails(ctx, userID)
  
  // report retrieved data or error upstream
  ...
}

// GetUserDetails implements a Repository interface.
func (r *Repository) GetUserDetails(ctx context.Context, userID int) (*Details, error) {
  // We received the Go-context from the upstream function calling us. 
  // This allows us to adhere to deadlines, cancellation and be able to 
  // participate in distributed traces.
  
  // We need to use the __Context versions of database/sql methods and
  // provide them with the Go-context we've been passing along.
  // Since r.db was created using our ocsql driver wrapper it will use
  // the span found in Go-context as parent for the creation of a new
  // child span timing our SQL Query.
  res, err := r.db.QueryContext(
    ctx, "SELECT * FROM details WHERE user_id = $1", userID,
  )
  
  // handle the result and return either the retrieved details or error
  ...
}
```

Before the release of Go 1.8, the database/sql package was not Go-context aware and most people still use the non context aware methods. To take full advantage of ocsql you will need to convert them. This exercise is fairly easy if you’ve already been passing along Go-context in your functions on the call path. If not you have more work to do, but the rewards will be big. Next to being able to have ocsql participate in your distributed traces you can also stop (potentially expensive) queries if the request is cancelled upstream or apply a timeout for queries if you want to set a maximum time budget for being able to handle the call.

## ocsql
Since ocsql only deals with instrumentation of your database queries you will also need to have instrumented your service to service interactions. OpenCensus provides client and server middlewares for HTTP and gRPC in the opencensus-go project and Go kit provides native OpenCensus tracing middleware for its gRPC and HTTP transports, which I’ll highlight in a future post. For those of you using Twirp we have an example showing how to use Twirp in combination with OpenCensus.

So now that you have added your transport instrumentation, passed along Go-context to all functions on the call path and have used the context aware database/sql methods; it’s time to wire in ocsql.

With proper Go-context propagation you only need to wrap the database drivers you use in your services with ocsql and adjust the sql.Open calls to use the wrapped driver names instead of the original database driver name. And that’s it!

So now that you have added your transport instrumentation, passed along Go-context to all functions on the call path and have used the context aware database/sql methods; it’s time to wire in ocsql.

With proper Go-context propagation you only need to wrap the database drivers you use in your services with ocsql and adjust the sql.Open calls to use the wrapped driver names instead of the original database driver name. And that’s it!

```go
import (
    "github.com/basvanbeek/ocsql"
    _ "github.com/mattn/go-sqlite3"
)

var (
    driverName string
    err        error
    db         *sql.DB
)

func main() {
    // set-up our OpenCensus instrumentation and exporters
    ...
    
    // Register our ocsql wrapper for the provided SQLite3 driver.
    driverName, err = ocsql.Register("sqlite3", ocsql.WithAllTraceOptions())
    if err != nil {
        log.Fatalf("unable to register our ocsql driver: %v\n", err)
    }

    // Connect to a SQLite3 database using the ocsql driver wrapper.
    db, err = sql.Open(driverName, "resource.db")
    
    ...
}
```
Bonus points to those that noticed that ocsql.Register returns a dynamic driver name. This was done on purpose and allows you to wrap the same database driver with differently configured ocsql wrappers. This way you can have various levels of instrumentation granularity for the same database driver. You might want more granular details for one database connection pool over the other.

## example
I’ve published a small example service on github, highlighting the ocsql instrumentation in combination with Zipkin as the tracing backend. The images below show traces that originated from the project: https://github.com/basvanbeek/ocsql-example

conclusion
Adding ocsql to your microservices is fairly trivial in a greenfield deployment, you just need to make sure you pass along Go-context to all functions on the call path. In existing code you might have more work to get Go-context propagated but the exercise is well worth it. Next to enabling the valuable insights distributed tracing can provide, you also get a chance to set cancellation and deadline policies which ultimately improve your ecosystem’s expected behavior, performance and reliability.

links
* OpenCensus: https://opencensus.io/
* ocsql: https://github.com/basvanbeek/ocsql
* ocsql-example: https://github.com/basvanbeek/ocsql-example
* Prometheus: https://prometheus.io/
* Zipkin: https://zipkin.io/
* Go kit: http://gokit.io/

Content from https://medium.com/@bas.vanbeek/opencensus-and-go-database-sql-322a26be5cc5
