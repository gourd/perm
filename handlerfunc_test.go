package perm

import (
	"net/http"
	"testing"
)

func TestHandlerFunc1(t *testing.T) {

	// fix interface parameter
	var f1 HandlerFunc = func(r *http.Request, perm string, info ...interface{}) error {
		return nil
	}

	// no parameter
	f1(nil, "some perm")

	// permission with some info
	f1(nil, "some perm with info", 1, 2)

}

func TestHandlerFunc2(t *testing.T) {

	// test if HandlerFunc implements Handler interface
	var f2 HandlerFunc = func(r *http.Request, perm string, info ...interface{}) error {
		return nil
	}

	var h Handler = f2

	// no parameter
	h.Allow(nil, "some perm")

	// permission with some info
	h.Allow(nil, "some perm with info", 1, 2)

}
