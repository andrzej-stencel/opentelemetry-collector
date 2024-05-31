// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/receiver/scraperhelper")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/receiver/scraperhelper")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	ScraperErroredMetricPoints metric.Int64Counter
	ScraperScrapedMetricPoints metric.Int64Counter
	level                      configtelemetry.Level
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// WithLevel sets the current telemetry level for the component.
func WithLevel(lvl configtelemetry.Level) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.level = lvl
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{level: configtelemetry.LevelBasic}
	for _, op := range options {
		op(&builder)
	}
	var (
		err, errs error
		meter     metric.Meter
	)
	if builder.level >= configtelemetry.LevelBasic {
		meter = Meter(settings)
	} else {
		meter = noop.Meter{}
	}
	builder.ScraperErroredMetricPoints, err = meter.Int64Counter(
		"scraper_errored_metric_points",
		metric.WithDescription("Number of metric points that were unable to be scraped."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	builder.ScraperScrapedMetricPoints, err = meter.Int64Counter(
		"scraper_scraped_metric_points",
		metric.WithDescription("Number of metric points successfully scraped."),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
