// Copyright 2022 Board of Trustees of the University of Illinois
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

package logs

import "net/http"

// HTTPRequestProperties is an entity which contains the properties of an HTTP request
type HTTPRequestProperties struct {
	Method     string
	Path       string
	RemoteAddr string
	UserAgent  string
}

// Match returns true if the provided http.Request matches the defined request properties
func (h HTTPRequestProperties) Match(r *http.Request) bool {
	if h.Method != "" && h.Method != r.Method {
		return false
	}

	if h.Path != "" && h.Path != r.URL.Path {
		return false
	}

	if h.RemoteAddr != "" && h.RemoteAddr != r.RemoteAddr {
		return false
	}

	if h.UserAgent != "" && h.UserAgent != r.UserAgent() {
		return false
	}

	return true
}

// NewAwsHealthCheckHTTPRequestProperties creates an HTTPRequestProperties object for a standard AWS ELB health checker
//
//	Path: The path that the health checks are performed on. If empty, "/version" is used as the default value.
func NewAwsHealthCheckHTTPRequestProperties(path string) HTTPRequestProperties {
	if path == "" {
		path = "/version"
	}
	return HTTPRequestProperties{Method: "GET", Path: path, UserAgent: "ELB-HealthChecker/2.0"}
}

// NewOpenShiftHealthCheckHTTPRequestProperties creates an HTTPRequestProperties object for a standard OpenShift health checker
//
//	Path: The path that the health checks are performed on. If empty, "/version" is used as the default value.
func NewOpenShiftHealthCheckHTTPRequestProperties(path string) HTTPRequestProperties {
	if path == "" {
		path = "/version"
	}
	return HTTPRequestProperties{Method: "GET", Path: path, UserAgent: "kube-probe/1.22+"}
}

// NewStandardHealthCheckHTTPRequestProperties creates a list of HTTPRequestProperties objects for known standard health checkers
//
//	Path: The path that the health checks are performed on. If empty, "/version" is used as the default value.
func NewStandardHealthCheckHTTPRequestProperties(path string) []HTTPRequestProperties {
	if path == "" {
		path = "/version"
	}
	return []HTTPRequestProperties{
		NewAwsHealthCheckHTTPRequestProperties(path),
		NewOpenShiftHealthCheckHTTPRequestProperties(path),
	}
}
