package slog

import "net/http"

type DebugPrintURLer struct {
	h http.Handler
}

func (l DebugPrintURLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	DebugPrint("Static Path requested: ", r.URL.String())
	l.h.ServeHTTP(w, r)
}

func DebugPrintURL(f http.Handler) DebugPrintURLer {
	return DebugPrintURLer{h: f}
}
