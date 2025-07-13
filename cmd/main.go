package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/shivamks5/userserv/endpoint"
	"github.com/shivamks5/userserv/metrics"
	"github.com/shivamks5/userserv/middleware"
	"github.com/shivamks5/userserv/service"
	"github.com/shivamks5/userserv/transport"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	mt := metrics.NewPrometheusMetrics("user_service")

	svc := service.NewUserService()
	svc = middleware.NewMetricsMiddleware(svc, mt)
	svc = middleware.NewLoggingMiddleware(logger, svc)

	eps := endpoint.NewEndpoints(svc)
	handler := transport.MakeHTTPHandler(eps)

	fmt.Println("server port : 3000, http://localhost:3000")
	stdlog.Fatal(http.ListenAndServe(":3000", handler))
}
