package duck_feed

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	DuckFeedReportAPIRequestsDuration = kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Name:       "ge_api_gateway_duck_feed_api_requests_duration_seconds",
		Help:       "DuckFeed Report api requests duration in seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01},
	}, []string{"endpoint"})

	DuckFeedReportAPIRequestsTotal = kitprometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Name: "ge_api_gateway_duck_feed_api_requests_total",
		Help: "DuckFeed Report api requests count",
	}, []string{"endpoint"})
)
