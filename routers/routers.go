package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// API versioning
	//api := r.PathPrefix("/api").Subrouter()
	//v1 := api.PathPrefix("/v1").Subrouter()
	//v2 := api.PathPrefix("/v2").Subrouter()

	// Health check
	r.HandleFunc("/health", HealthCheck).Methods("GET")

	// Setup route groups
	SetupHomeRoutes(r)
	//SetupUserRoutes(v1)
	//SetupTaskRoutes(v1)
	//SetupAssetRoutes(v1)
	//SetupCalendarRoutes(v1)
	//SetupReportRoutes(v1)

	// GraphQL route
	//setupGraphQLRoutes(r)

	return r
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
