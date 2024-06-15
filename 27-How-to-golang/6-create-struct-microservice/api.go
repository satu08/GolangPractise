package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJson(w, http.StatusOK, fact)

}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
