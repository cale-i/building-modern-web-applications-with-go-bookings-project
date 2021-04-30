package main

import (
	"fmt"
	"testing"

	"github.com/cale-i/building-modern-web-applications-with-go-bookings-project/internal/config"
	"github.com/go-chi/chi"
)

func Test_routes(t *testing.T) {
	var app *config.AppConfig

	mux := routes(app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing: test passed
	default:
		t.Error(fmt.Sprintf("type is not *chiMux, type is %T", v))
	}
}
