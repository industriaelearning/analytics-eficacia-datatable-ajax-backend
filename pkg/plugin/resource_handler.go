package plugin

import (
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"net/http"

	"eficacia-datatable-backend/pkg/query/scenario"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/resource/httpadapter"
)

func newResourceHandler(settings *models.PluginSettings) backend.CallResourceHandler {
	mux := http.NewServeMux()
	mux.HandleFunc("/query-types", handleQueryTypes)
	mux.HandleFunc("/tenants", func(w http.ResponseWriter, r *http.Request) {
		handlerGetTenants(w, r, settings)
	})
	mux.HandleFunc("/programs", func(w http.ResponseWriter, r *http.Request) {
		handlerGetPrograms(w, r, settings)
	})
	mux.HandleFunc("/organisationpositions", func(w http.ResponseWriter, r *http.Request) {
		handlerOrganisationPositions(w, r, settings)
	})
	mux.HandleFunc("/organisationdepartments", func(w http.ResponseWriter, r *http.Request) {
		handlerOrganisationDepartments(w, r, settings)
	})
	mux.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		handlerCourses(w, r, settings)
	})
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		handlerData(w, r, settings)
	})
	mux.HandleFunc("/download-handler", func(w http.ResponseWriter, r *http.Request) {
		handleQueryDownloads(w, r, settings)
	})

	return httpadapter.New(mux)
}

type queryTypesResponse struct {
	QueryTypes []string `json:"queryTypes"`
}

func handleQueryTypes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	queryTypes := &queryTypesResponse{
		QueryTypes: []string{
			scenario.GetEmpty,
		},
	}

	j, err := json.Marshal(queryTypes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
