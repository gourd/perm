package perm

import (
	"net/http"
	"testing"
)

func TestDefaultMux(t *testing.T) {
	// test if default mux implements mux
	var m Mux = NewMux()
	_ = m
}

func TestMuxFoundFunc(t *testing.T) {
	m := NewMux()
	m.HandleFunc("access something", func(r *http.Request, perm string, info ...interface{}) error {
		return nil
	})
	if err := m.Allow(nil, "access something"); err != nil {
		t.Errorf("Unexpected error. Failed to obtain handler for permission")
	}
}

func TestMuxFoundMux(t *testing.T) {

	// child mux
	m1 := NewMux()
	m1.HandleFunc("access something", func(r *http.Request, perm string, info ...interface{}) error {
		return nil
	})

	// parent mux
	m2 := NewMux()
	m2.Handle("access something", m1)

	// test parent mux
	if err := m2.Allow(nil, "access something"); err != nil {
		t.Errorf("Unexpected error. Failed to obtain handler for permission")
	}
}

func TestMuxNotFound(t *testing.T) {
	m := NewMux()
	err := m.Allow(nil, "access something")
	if err != HandlerNotFound {
		t.Errorf("Error is not of expected type. Expecting perm.HandlerNotFound by get %#v", err)
	}
}

func TestMuxHttpHandler(t *testing.T) {
	// make sure mux implements http.Handler
	// can be used as a middleware
	var h http.Handler = NewMux()
	_ = h
}
