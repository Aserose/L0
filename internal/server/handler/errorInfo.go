package handler

import (
	"L0/internal/service"
	"L0/pkg/customLogger"
	"net/http"
)

type errorInfo struct {
	service service.Service
	log     customLogger.Logger
}

func newErrInfo(log customLogger.Logger, s service.Service) errorInfo {
	return errorInfo{
		service: s,
		log:     log,
	}
}

func (e errorInfo) setup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		e.getLast(w, r)

	}
}

func (e errorInfo) getLast(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(e.service.GetLast()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (e errorInfo) getLastFormat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`last error: ` + e.service.GetLast()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
