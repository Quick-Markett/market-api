package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hermes_api_total_requests",
			Help: "Total nubmer of requests processed by the Hermes API",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hermes_api_total_error_requests",
			Help: "Total number of error requests processed by Hermes API",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()

		status := strconv.Itoa(c.Writer.Status())

		RequestCount.WithLabelValues(path, status).Inc()

		if c.Writer.Status() >= http.StatusBadRequest {
			ErrorCount.WithLabelValues(path, status).Inc()
		}
	}
}
