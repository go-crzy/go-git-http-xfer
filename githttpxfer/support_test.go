package githttpxfer

import (
	"net/http"
	"net/http/httptest"
	"os/exec"
	"runtime"
	"testing"
)

func Test_GetServiceType(t *testing.T) {

	tests := []struct {
		description string
		method      string
		url         string
		expected    string
	}{
		{
			description: "it should return upload-pack",
			method:      http.MethodPost,
			url:         "http://example.com/base/foo/git-upload-pack?service=git-upload-pack",
			expected:    "upload-pack",
		},
		{
			description: "it should return receive-pack",
			method:      http.MethodPost,
			url:         "http://example.com/base/foo/git-upload-pack?service=git-receive-pack",
			expected:    "receive-pack",
		},
		{
			description: "it should return empty",
			method:      http.MethodPost,
			url:         "http://example.com/base/foo/git-upload-pack?service=foo-receive-pack",
			expected:    "",
		},
	}

	for _, tc := range tests {
		r := httptest.NewRequest(tc.method, tc.url, nil)

		if serviceType := getServiceType(r); tc.expected != serviceType {
			t.Errorf("service type is not %s . result: %s", tc.expected, serviceType)
		}
	}
}

func Test_cleanUpProcessGroup(t *testing.T) {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("tail", "-f", "/dev/null")
		cmd.Start()
		cleanUpProcessGroup(cmd)
		if err := cmd.Wait(); err != nil && err.Error() != "signal: killed" {
			t.Errorf("tail should have been killed, instead %v", err)
		}
	}
}
