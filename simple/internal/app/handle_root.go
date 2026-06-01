package app

import (
	"encoding/json"
	"net/http"
)

func loadProcessData() (*ProcessData, error) {
	p := &ProcessData{
		CPUUsage: 10.4,
		MemUsage: 33.6,
		CPU:      12,
		Memory:   16 * 1024 * 1024 * 1024,
	}
	return p, nil
}

type HandleRootResponse struct {
	ProcessData
}

func (s Server) HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	data, err := loadProcessData()
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
