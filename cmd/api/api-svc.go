package main

import (
	"encoding/json"
	"fmt"
	"github.com/digital-technology-agency/api-middleware/pkg/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	authData = auth{
		User:     "user",
		Password: "user123"}
)

type auth struct {
	User     string
	Password string
}

type Response struct {
	Status  int
	Service string
}

func main() {
	port := utils.GetEnv("PORT", "9999")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", health)
	mux.HandleFunc("/api/v1/protected", basicAuth(protected))
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 15,
	}
	log.Printf("Starting server on %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Err %v\n", err.Error())
	}
}

func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Content-Type", "application/json")
		user, password, ok := r.BasicAuth()
		if !ok {
			(w).WriteHeader(http.StatusUnauthorized)
			return
		}
		if authData.User != user || authData.Password != strings.ReplaceAll(password, "\n", "") {
			(w).WriteHeader(http.StatusUnauthorized)
			return
		}
		(w).WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(http.StatusOK)
	dataBytes, _ := json.Marshal(Response{Status: http.StatusOK, Service: "health"})
	(w).Write(dataBytes)
}

func protected(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")
	(w).WriteHeader(http.StatusOK)
	dataBytes, _ := json.Marshal(Response{Status: http.StatusOK, Service: "protected"})
	(w).Write(dataBytes)
}
