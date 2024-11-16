// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package normal // import "go.opentelemetry.io/collector/exporter/debugexporter/internal/normal"

import (
	"bytes"
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

type normalTracesMarshaler struct{}

// Ensure normalTracesMarshaller implements interface ptrace.Marshaler
var _ ptrace.Marshaler = normalTracesMarshaler{}

// NewNormalTracesMarshaler returns a ptrace.Marshaler for normal verbosity. It writes one line of text per log record
func NewNormalTracesMarshaler() ptrace.Marshaler {
	return normalTracesMarshaler{}
}

func (normalTracesMarshaler) MarshalTraces(md ptrace.Traces) ([]byte, error) {
	var buffer bytes.Buffer
	for i := 0; i < md.ResourceSpans().Len(); i++ {
		resourceTraces := md.ResourceSpans().At(i)

		resourceSchemaUrlString := writeSchemaUrlString(resourceTraces.SchemaUrl())
		resourceAttributesString := writeAttributesString(resourceTraces.Resource().Attributes())
		buffer.WriteString(fmt.Sprintf("ResourceTraces #%d%s%s\n", i, resourceSchemaUrlString, resourceAttributesString))

		for j := 0; j < resourceTraces.ScopeSpans().Len(); j++ {
			scopeTraces := resourceTraces.ScopeSpans().At(j)

			scopeSchemaUrlString := writeSchemaUrlString(scopeTraces.SchemaUrl())
			scopeAttributesString := writeAttributesString(scopeTraces.Scope().Attributes())
			buffer.WriteString(fmt.Sprintf("ScopeTraces #%d%s%s\n", i, scopeSchemaUrlString, scopeAttributesString))

			for k := 0; k < scopeTraces.Spans().Len(); k++ {
				span := scopeTraces.Spans().At(k)

				buffer.WriteString(span.Name())

				buffer.WriteString(" ")
				buffer.WriteString(span.TraceID().String())

				buffer.WriteString(" ")
				buffer.WriteString(span.SpanID().String())

				if span.Attributes().Len() > 0 {
					spanAttributes := writeAttributes(span.Attributes())
					buffer.WriteString(" ")
					buffer.WriteString(strings.Join(spanAttributes, " "))
				}

				buffer.WriteString("\n")
			}
		}
	}
	return buffer.Bytes(), nil
}
