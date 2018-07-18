// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package propagation implement X-Cloud-Trace-Context header propagation used
// by Google Cloud products.
package propagation // import "go.opencensus.io/exporter/stackdriver/propagation"

import (
	"net/http"

	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

const httpHeader = `X-Cloud-Trace-Context`

var _ propagation.HTTPFormat = (*HTTPFormat)(nil)

// HTTPFormat implements propagation.HTTPFormat to propagate
// traces in HTTP headers for Google Cloud Platform and Stackdriver Trace.
type HTTPFormat struct{}

// SpanContextFromRequest extracts a Stackdriver Trace span context from incoming requests.
func (f *HTTPFormat) SpanContextFromRequest(req *http.Request) (sc trace.SpanContext, ok bool) {
	return propagation.FromHTTPHeader(req.Header.Get(httpHeader))
}

// SpanContextToRequest modifies the given request to include a Stackdriver Trace header.
func (f *HTTPFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	req.Header.Set(httpHeader, propagation.HTTPHeader(sc))
}
