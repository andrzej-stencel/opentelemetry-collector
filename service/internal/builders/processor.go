// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package builders // import "go.opentelemetry.io/collector/service/internal/builders"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumerprofiles"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorprofiles"
	"go.opentelemetry.io/collector/processor/processortest"
)

// ProcessorBuilder processor is a helper struct that given a set of Configs
// and Factories helps with creating processors.
type ProcessorBuilder struct {
	cfgs      map[component.ID]component.Config
	factories map[component.Type]processor.Factory
}

// NewProcessor creates a new ProcessorBuilder to help with creating components form a set of configs and factories.
func NewProcessor(cfgs map[component.ID]component.Config, factories map[component.Type]processor.Factory) *ProcessorBuilder {
	return &ProcessorBuilder{cfgs: cfgs, factories: factories}
}

// CreateTraces creates a Traces processor based on the settings and config.
func (b *ProcessorBuilder) CreateTraces(ctx context.Context, set processor.Settings, next consumer.Traces) (processor.Traces, error) {
	if next == nil {
		return nil, errNilNextConsumer
	}
	cfg, existsCfg := b.cfgs[set.ID]
	if !existsCfg {
		return nil, fmt.Errorf("processor %q is not configured", set.ID)
	}

	f, existsFactory := b.factories[set.ID.Type()]
	if !existsFactory {
		return nil, fmt.Errorf("processor factory not available for: %q", set.ID)
	}

	logStabilityLevel(set.Logger, f.TracesProcessorStability())
	return f.CreateTracesProcessor(ctx, set, cfg, next)
}

// CreateMetrics creates a Metrics processor based on the settings and config.
func (b *ProcessorBuilder) CreateMetrics(ctx context.Context, set processor.Settings, next consumer.Metrics) (processor.Metrics, error) {
	if next == nil {
		return nil, errNilNextConsumer
	}
	cfg, existsCfg := b.cfgs[set.ID]
	if !existsCfg {
		return nil, fmt.Errorf("processor %q is not configured", set.ID)
	}

	f, existsFactory := b.factories[set.ID.Type()]
	if !existsFactory {
		return nil, fmt.Errorf("processor factory not available for: %q", set.ID)
	}

	logStabilityLevel(set.Logger, f.MetricsProcessorStability())
	return f.CreateMetricsProcessor(ctx, set, cfg, next)
}

// CreateLogs creates a Logs processor based on the settings and config.
func (b *ProcessorBuilder) CreateLogs(ctx context.Context, set processor.Settings, next consumer.Logs) (processor.Logs, error) {
	if next == nil {
		return nil, errNilNextConsumer
	}
	cfg, existsCfg := b.cfgs[set.ID]
	if !existsCfg {
		return nil, fmt.Errorf("processor %q is not configured", set.ID)
	}

	f, existsFactory := b.factories[set.ID.Type()]
	if !existsFactory {
		return nil, fmt.Errorf("processor factory not available for: %q", set.ID)
	}

	logStabilityLevel(set.Logger, f.LogsProcessorStability())
	return f.CreateLogsProcessor(ctx, set, cfg, next)
}

// CreateProfiles creates a Profiles processor based on the settings and config.
func (b *ProcessorBuilder) CreateProfiles(ctx context.Context, set processor.Settings, next consumerprofiles.Profiles) (processorprofiles.Profiles, error) {
	if next == nil {
		return nil, errNilNextConsumer
	}
	cfg, existsCfg := b.cfgs[set.ID]
	if !existsCfg {
		return nil, fmt.Errorf("processor %q is not configured", set.ID)
	}

	f, existsFactory := b.factories[set.ID.Type()]
	if !existsFactory {
		return nil, fmt.Errorf("processor factory not available for: %q", set.ID)
	}

	logStabilityLevel(set.Logger, f.ProfilesProcessorStability())
	return f.CreateProfilesProcessor(ctx, set, cfg, next)
}

func (b *ProcessorBuilder) Factory(componentType component.Type) component.Factory {
	return b.factories[componentType]
}

// NewNopProcessorConfigsAndFactories returns a configuration and factories that allows building a new nop processor.
func NewNopProcessorConfigsAndFactories() (map[component.ID]component.Config, map[component.Type]processor.Factory) {
	nopFactory := processortest.NewNopFactory()
	configs := map[component.ID]component.Config{
		component.NewID(nopType): nopFactory.CreateDefaultConfig(),
	}
	factories := map[component.Type]processor.Factory{
		nopType: nopFactory,
	}

	return configs, factories
}
