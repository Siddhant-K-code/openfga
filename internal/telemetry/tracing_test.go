package telemetry

import (
	"testing"
)

func TestParseOTLPEndpoint(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedEndpoint string
		expectedInsecure bool
	}{
		{
			name:             "bare_host_port",
			input:            "collector.example.com:4317",
			expectedEndpoint: "collector.example.com:4317",
			expectedInsecure: false,
		},
		{
			name:             "http_scheme",
			input:            "http://collector.example.com:4317",
			expectedEndpoint: "collector.example.com:4317",
			expectedInsecure: true,
		},
		{
			name:             "https_scheme",
			input:            "https://collector.example.com:4317",
			expectedEndpoint: "collector.example.com:4317",
			expectedInsecure: false,
		},
		{
			name:             "http_k8s_service_dns",
			input:            "http://k8se-otel.k8se-apps.svc.cluster.local:4317",
			expectedEndpoint: "k8se-otel.k8se-apps.svc.cluster.local:4317",
			expectedInsecure: true,
		},
		{
			name:             "localhost",
			input:            "0.0.0.0:4317",
			expectedEndpoint: "0.0.0.0:4317",
			expectedInsecure: false,
		},
		{
			name:             "http_localhost",
			input:            "http://localhost:4317",
			expectedEndpoint: "localhost:4317",
			expectedInsecure: true,
		},
		{
			name:             "https_no_port",
			input:            "https://collector.example.com",
			expectedEndpoint: "collector.example.com",
			expectedInsecure: false,
		},
		{
			name:             "http_no_port",
			input:            "http://collector.example.com",
			expectedEndpoint: "collector.example.com",
			expectedInsecure: true,
		},
		{
			name:             "empty_string",
			input:            "",
			expectedEndpoint: "",
			expectedInsecure: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint, insecure := parseOTLPEndpoint(tt.input)
			if endpoint != tt.expectedEndpoint {
				t.Errorf("parseOTLPEndpoint(%q) endpoint = %q, want %q", tt.input, endpoint, tt.expectedEndpoint)
			}
			if insecure != tt.expectedInsecure {
				t.Errorf("parseOTLPEndpoint(%q) insecure = %v, want %v", tt.input, insecure, tt.expectedInsecure)
			}
		})
	}
}
