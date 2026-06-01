package app

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct{}

func (s Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write([]byte("OK")); err != nil {
		panic(err)
	}
}

func loadRoutes(server *Server) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.HandleHealth)
	mux.HandleFunc("/", server.HandleRoot)
	return mux
}

func StartAPI(conf Config) error {
	s := &Server{}
	slog.Info("Initializing Server", "environment", conf.Environment, "port", conf.Port, "host", conf.Host)
	mux := loadRoutes(s)
	slog.Info("Listening to Requests", "environment", conf.Environment, "port", conf.Port, "host", conf.Host)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Host, conf.Port), mux); err != nil {
		return fmt.Errorf("StartAPI ListenAndServe failed: %w", err)
	}
	slog.Info("Shutting down server", "environment", conf.Environment)
	return nil
}
