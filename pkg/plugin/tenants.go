package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"net/http"
)

// handleGetTenants is the handler for the /tenants endpoint.
func handlerGetTenants(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	db, err := connectDB(settings)
	CheckError(err)
	defer db.Close()

	dataFromDatabaseInJson, _ := getTenantsList(db)
	dataJson, _ := json.Marshal(dataFromDatabaseInJson)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(dataJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTenantsList(db *sql.DB) ([]json.RawMessage, error) {
	var dataFromDatabaseInJson []json.RawMessage
	var queryStm = "SELECT id, name FROM tool_tenant ORDER BY name"
	rows, err := db.Query(queryStm)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		CheckError(err)

		data, err := json.Marshal(map[string]interface{}{
			"id":   id,
			"name": name,
		})
		CheckError(err)

		dataFromDatabaseInJson = append(dataFromDatabaseInJson, data)
	}

	if dataFromDatabaseInJson == nil {
		dataFromDatabaseInJson = []json.RawMessage{}
	}

	return dataFromDatabaseInJson, nil
}
