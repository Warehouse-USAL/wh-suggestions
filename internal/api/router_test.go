package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	r := NewRouter()
	cases := []struct {
		method, path string
		want         int
	}{
		{http.MethodGet, "/healthz", http.StatusOK},
		{http.MethodPost, "/v1/recommendations/cart", http.StatusNotImplemented},
		{http.MethodGet, "/v1/recommendations/user/u1", http.StatusNotImplemented},
		{http.MethodGet, "/v1/recommendations/hot", http.StatusNotImplemented},
		{http.MethodGet, "/v1/recommendations/cart", http.StatusNotFound},
	}
	for _, tc := range cases {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		if w.Code != tc.want {
			t.Errorf("%s %s = %d, want %d", tc.method, tc.path, w.Code, tc.want)
		}
	}
}
