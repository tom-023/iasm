package monitor

import (
	"github.com/stretchr/testify/assert"
	"github.com/tom-023/iasm/config"
	"github.com/tom-023/iasm/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIsRespond(t *testing.T) {
	config.Init()
	logger.Init()
	defer logger.Sync()

	// Create a test server that responds with 200 OK
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	// Create a test server that responds with 500 Internal Server Error
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{"Responds with 200 OK", testServer.URL, true},
		{"Responds with 500 Error", errorServer.URL, false},
		{"Invalid URL", "http://invalid-url", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isRespond(tt.url, 2*time.Second)
			assert.Equal(t, tt.expected, result)
		})
	}
}
