package app

// ProcessData is a representation of the information collected by the application
type ProcessData struct {
	CPUUsage float64 `json:"cpu_usage"`
	MemUsage float64 `json:"mem_usage"`
	CPU      float64 `json:"cpu"`
	Memory   float64 `json:"mem"`
}
