// Package play exposes the placeholder playback surface for tv-stream.
//
// Live MPEG-TS multicast -> HLS bridge using a stream-copy remux.
package play

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/zaentrum/tv-stream/internal/config"
)

func Routes(_ config.Config) http.Handler {
	r := chi.NewRouter()
	r.Get("/*", placeholder)
	return r
}

func placeholder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte("tv-stream playback not yet implemented; scaffold only\n"))
}
