package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerGetPrograms(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	db, err := connectDB(settings)
	if err != nil {
		CheckError(err)
	}
	defer db.Close()

	dataFromDatabaseInJson, _ := getProgramsList(db, r)
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
 * Get list of programs from database filtered by tenantId if it is provided
 */
func getProgramsList(db *sql.DB, r *http.Request) ([]json.RawMessage, error) {
	var dataFromDatabaseInJson []json.RawMessage
	tenantId := r.URL.Query().Get("tenantid")
	if tenantId == "" {
		tenantId = "1"
	}
	rows, err := db.Query(fmt.Sprintf("SELECT id, fullname FROM tool_program WHERE tenantid IN (%s) ORDER BY fullname", tenantId))
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var id, fullname string

		err := rows.Scan(&id, &fullname)
		CheckError(err)

		data, err := json.Marshal(map[string]interface{}{
			"id":   id,
			"name": fullname,
		})
		CheckError(err)

		dataFromDatabaseInJson = append(dataFromDatabaseInJson, data)
	}

	if dataFromDatabaseInJson == nil {
		dataFromDatabaseInJson = []json.RawMessage{}
	}

	return dataFromDatabaseInJson, nil
}
