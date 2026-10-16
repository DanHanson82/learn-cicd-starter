package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expected      string
		expectedError string
	}{
		{
			name: "with valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-key"},
			},
			expected:      "test-key",
			expectedError: "",
		},
		{
			name: "with malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer test-token"},
			},
			expected:      "",
			expectedError: "malformed authorization header",
		},
		{
			name:          "missing authorization header",
			headers:       http.Header{},
			expected:      "",
			expectedError: "no authorization header included",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetAPIKey(tc.headers)
			if (tc.expectedError == "") && (err != nil || result != tc.expected) {
				t.Errorf("expected %q, got %q, %v", tc.expected, result, err)
			} else if tc.expectedError != "" && tc.expectedError != err.Error() {

				t.Errorf("expected %q, got %q, %v", tc.expected, result, err)
			}
		})
	}
}
