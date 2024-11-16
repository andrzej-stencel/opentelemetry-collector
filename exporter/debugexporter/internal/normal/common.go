// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package normal // import "go.opentelemetry.io/collector/exporter/debugexporter/internal/normal"

import (
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

// writeAttributes returns a slice of strings in the form "attrKey=attrValue"
func writeAttributes(attributes pcommon.Map) (attributeStrings []string) {
	attributes.Range(func(k string, v pcommon.Value) bool {
		attribute := fmt.Sprintf("%s=%s", k, v.AsString())
		attributeStrings = append(attributeStrings, attribute)
		return true
	})
	return attributeStrings
}

// writeAttributesString returns a string in the form " attrKey=attrValue attr2=value2"
func writeAttributesString(attributesMap pcommon.Map) (attributesString string) {
	attributes := writeAttributes(attributesMap)
	if len(attributes) > 0 {
		attributesString = fmt.Sprintf(" %s", strings.Join(attributes, " "))
	}
	return attributesString
}

func writeSchemaUrlString(schemaUrl string) (schemaUrlString string) {
	if len(schemaUrl) > 0 {
		schemaUrlString = fmt.Sprintf(" SchemaUrl=%s", schemaUrl)
	}
	return schemaUrlString
}
