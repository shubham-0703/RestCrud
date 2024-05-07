package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// API represents the RESTful API for managing employees
type API struct {
	store *EmployeeStore
}

// NewAPI creates a new instance of the API
func NewAPI(store *EmployeeStore) *API {
	return &API{store: store}
}

// ListEmployeesHandler handles requests to list employees with pagination
func (api *API) ListEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		http.Error(w, "Invalid pageSize", http.StatusBadRequest)
		return
	}

	api.store.Lock()
	defer api.store.Unlock()

	start := (page - 1) * pageSize
	end := start + pageSize

	employees := make([]Employee, 0)
	for _, employee := range api.store.employees {
		employees = append(employees, employee)
	}

	if start >= len(employees) {
		http.Error(w, "Page out of range", http.StatusBadRequest)
		return
	}

	if end > len(employees) {
		end = len(employees)
	}

	pageEmployees := employees[start:end]

	jsonBytes, err := json.Marshal(pageEmployees)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// InitializeAPI initializes the API endpoints
func InitializeAPI(store *EmployeeStore) {
	api := NewAPI(store)
	http.HandleFunc("/employees", api.ListEmployeesHandler)

	// Create a new HTTP server
	server := &http.Server{
		Addr:    ":8080", // Set your desired port
		Handler: nil,     // Use the default ServeMux
	}
	// Start the HTTP server
	log.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
