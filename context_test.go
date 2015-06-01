package perm

import (
	"net/http"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	// test the middleware operation
	m1 := NewMux()
	r := &http.Request{}
	m1.ServeHTTP(nil, r)
	m2 := GetMux(r)

	if m1 != m2 {
		t.Errorf("Middleware got is not Mux")
	}
}
