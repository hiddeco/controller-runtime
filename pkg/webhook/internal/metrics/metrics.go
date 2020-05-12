/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	// RequestLatency is a prometheus metric which is a histogram of the latency
	// of processing admission requests.
	RequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "controller_runtime_webhook_latency_seconds",
			Help: "Histogram of the latency of processing admission requests",
		},
		[]string{"webhook", "code"},
	)

	// RequestTotal is a prometheus metric which is a counter of the total processed admission requests.
	RequestTotal = func() *prometheus.CounterVec {
		return prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "controller_runtime_webhook_requests_total",
				Help: "Total number of admission requests by HTTP status code.",
			},
			[]string{"webhook", "code"},
		)
	}()

	// RequestInFlight is a prometheus metric which is a gauge of the in-flight admission requests.
	RequestInFlight = func() *prometheus.GaugeVec {
		return prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "controller_runtime_webhook_requests_in_flight",
				Help: "Current number of admission requests being served.",
			},
			[]string{"webhook"},
		)
	}()
)

func init() {
	metrics.Registry.MustRegister(RequestLatency, RequestTotal, RequestInFlight)
}
