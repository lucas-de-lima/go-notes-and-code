package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandle(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	res := httptest.NewRecorder()

	PingHandler(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("esperado status 200, obtido %d", res.Code)
	}

	if res.Body.String() != "OK" {
		t.Errorf("esperado corpo 'OK', obtido '%s'", res.Body.String())
	}

}
