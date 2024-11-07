package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"service-catalog/customerr"
)

// GetServices handles GET /services
func GetServices(w http.ResponseWriter, r *http.Request) {
	filter := Filter{}

	// Decode JSON body into filter struct
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		customerr.JSONErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	services, err := GetServicesCore(filter)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(services)
}

// GetServiceById handles GET /services/{serviceId}
func GetServiceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceId := vars["serviceId"]

	service, err := GetServiceByIdCore(serviceId)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(service)
}

// GetServiceVersions handles GET /services/{serviceId}/versions
func GetServiceVersions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceId := vars["serviceId"]

	versions, err := GetServiceVersionsCore(serviceId)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(versions)
}

// CreateService handles POST /services
func CreateService(w http.ResponseWriter, r *http.Request) {
	var service Service
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(service)
}

// UpdateService handles POST /services/{serviceId}
func UpdateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceId := vars["serviceId"]

	// Validate serviceId
	err := ValidateGetServiceById(serviceId)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var service Service
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = UpdateServiceCore(serviceId, &service)
	if err != nil {
		customerr.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(service)
}
