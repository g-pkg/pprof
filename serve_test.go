package pprof

import (
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
	go func() {
		t.Log(Serve("localhost:15001"))
	}()
	resp, err := http.Get("http://localhost:15001/debug/pprof")
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code: got %d; want %d", resp.StatusCode, http.StatusOK)
	}
}
