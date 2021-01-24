package test

import (
	"fmt"
	p "github.com/eriktisme/passport/cmd/passport"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {
	ts := httptest.NewServer(p.SetupServer())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestLoginRoute(t *testing.T) {
	r := p.SetupServer()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginPayload()))

	httpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusOK
	})
}

func loginPayload() string {
	payload := url.Values{}
	payload.Add("email", "john@example.org")
	payload.Add("password", "secret")

	return payload.Encode()
}
