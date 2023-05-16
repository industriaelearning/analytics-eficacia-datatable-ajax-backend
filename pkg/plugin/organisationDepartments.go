package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerOrganisationDepartments(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	db, err := connectDB(settings)
	if err != nil {
		CheckError(err)
	}
	defer db.Close()

	dataFromDatabaseInJson, _ := getOrganisationDepartmentsList(db, r)
	dataJson, _ := json.Marshal(dataFromDatabaseInJson)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(dataJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

/**
 * Get list of organisation departments from database filtered by tenantId if it is provided
 */
func getOrganisationDepartmentsList(db *sql.DB, r *http.Request) ([]json.RawMessage, error) {
	var dataFromDatabaseInJson []json.RawMessage
	tenantId := r.URL.Query().Get("tenantid")
	if tenantId == "" {
		tenantId = "1"
	}
	rows, err := db.Query(fmt.Sprintf("SELECT id, name FROM tool_organisation_department WHERE tenantid IN (%s) ORDER BY name", tenantId))
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

		dataFromDatabaseInJson = append(dataFromDatabaseInJson, json.RawMessage(data))
	}
	if dataFromDatabaseInJson == nil {
		dataFromDatabaseInJson = []json.RawMessage{}
	}

	return dataFromDatabaseInJson, nil
}
