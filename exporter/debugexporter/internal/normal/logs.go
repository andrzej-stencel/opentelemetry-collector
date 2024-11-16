// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package normal // import "go.opentelemetry.io/collector/exporter/debugexporter/internal/normal"

import (
	"bytes"
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/pdata/plog"
)

type normalLogsMarshaler struct{}

// Ensure normalLogsMarshaller implements interface plog.Marshaler
var _ plog.Marshaler = normalLogsMarshaler{}

// NewNormalLogsMarshaler returns a plog.Marshaler for normal verbosity. It writes one line of text per log record
func NewNormalLogsMarshaler() plog.Marshaler {
	return normalLogsMarshaler{}
}

func (normalLogsMarshaler) MarshalLogs(ld plog.Logs) ([]byte, error) {
	var buffer bytes.Buffer
	for i := 0; i < ld.ResourceLogs().Len(); i++ {
		resourceLog := ld.ResourceLogs().At(i)

		resourceSchemaUrlString := writeSchemaUrlString(resourceLog.SchemaUrl())
		resourceAttributesString := writeAttributesString(resourceLog.Resource().Attributes())
		buffer.WriteString(fmt.Sprintf("ResourceLog #%d%s%s\n", i, resourceSchemaUrlString, resourceAttributesString))

		for j := 0; j < resourceLog.ScopeLogs().Len(); j++ {
			scopeLog := resourceLog.ScopeLogs().At(j)

			scopeSchemaUrlString := writeSchemaUrlString(scopeLog.SchemaUrl())
			scopeAttributesString := writeAttributesString(scopeLog.Scope().Attributes())
			buffer.WriteString(fmt.Sprintf("ScopeLog #%d%s%s\n", i, scopeSchemaUrlString, scopeAttributesString))

			for k := 0; k < scopeLog.LogRecords().Len(); k++ {
				logRecord := scopeLog.LogRecords().At(k)
				logAttributes := writeAttributes(logRecord.Attributes())

				logString := fmt.Sprintf("%s %s", logRecord.Body().AsString(), strings.Join(logAttributes, " "))
				buffer.WriteString(logString)
				buffer.WriteString("\n")
			}
		}
	}
	return buffer.Bytes(), nil
}
