// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)
	tb.LoadbalancerBackendLatency.Record(context.Background(), 1)
	tb.LoadbalancerBackendOutcome.Add(context.Background(), 1)
	tb.LoadbalancerNumBackendUpdates.Add(context.Background(), 1)
	tb.LoadbalancerNumBackends.Record(context.Background(), 1)
	tb.LoadbalancerNumResolutions.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_loadbalancer_backend_latency",
			Description: "Response latency in ms for the backends.",
			Unit:        "ms",
			Data: metricdata.Histogram[int64]{
				Temporality: metricdata.CumulativeTemporality,
				DataPoints: []metricdata.HistogramDataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_loadbalancer_backend_outcome",
			Description: "Number of successes and failures for each endpoint.",
			Unit:        "{outcomes}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_loadbalancer_num_backend_updates",
			Description: "Number of times the list of backends was updated.",
			Unit:        "{updates}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_loadbalancer_num_backends",
			Description: "Current number of backends in use.",
			Unit:        "{backends}",
			Data: metricdata.Gauge[int64]{
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_loadbalancer_num_resolutions",
			Description: "Number of times the resolver has triggered new resolutions.",
			Unit:        "{resolutions}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}