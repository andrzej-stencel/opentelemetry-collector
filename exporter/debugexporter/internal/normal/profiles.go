// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package normal // import "go.opentelemetry.io/collector/exporter/debugexporter/internal/normal"

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"go.opentelemetry.io/collector/pdata/pprofile"
)

type normalProfilesMarshaler struct{}

// Ensure normalProfilesMarshaller implements interface pprofile.Marshaler
var _ pprofile.Marshaler = normalProfilesMarshaler{}

// NewNormalProfilesMarshaler returns a pprofile.Marshaler for normal verbosity. It writes one line of text per log record
func NewNormalProfilesMarshaler() pprofile.Marshaler {
	return normalProfilesMarshaler{}
}

func (normalProfilesMarshaler) MarshalProfiles(pd pprofile.Profiles) ([]byte, error) {
	var buffer bytes.Buffer
	for i := 0; i < pd.ResourceProfiles().Len(); i++ {
		resourceProfiles := pd.ResourceProfiles().At(i)

		resourceSchemaUrlString := writeSchemaUrlString(resourceProfiles.SchemaUrl())
		resourceAttributesString := writeAttributesString(resourceProfiles.Resource().Attributes())
		buffer.WriteString(fmt.Sprintf("ResourceProfiles #%d%s%s\n", i, resourceSchemaUrlString, resourceAttributesString))

		for j := 0; j < resourceProfiles.ScopeProfiles().Len(); j++ {
			scopeProfiles := resourceProfiles.ScopeProfiles().At(j)

			scopeSchemaUrlString := writeSchemaUrlString(scopeProfiles.SchemaUrl())
			scopeAttributesString := writeAttributesString(scopeProfiles.Scope().Attributes())
			buffer.WriteString(fmt.Sprintf("ScopeProfiles #%d%s%s\n", i, scopeSchemaUrlString, scopeAttributesString))

			for k := 0; k < scopeProfiles.Profiles().Len(); k++ {
				profile := scopeProfiles.Profiles().At(k)

				buffer.WriteString(profile.ProfileID().String())

				buffer.WriteString(" samples=")
				buffer.WriteString(strconv.Itoa(profile.Sample().Len()))

				if profile.AttributeIndices().Len() > 0 {
					attrs := []string{}
					for _, i := range profile.AttributeIndices().AsRaw() {
						a := profile.AttributeTable().At(int(i))
						attrs = append(attrs, fmt.Sprintf("%s=%s", a.Key(), a.Value().AsString()))
					}

					buffer.WriteString(" ")
					buffer.WriteString(strings.Join(attrs, " "))
				}

				buffer.WriteString("\n")
			}
		}
	}
	return buffer.Bytes(), nil
}
