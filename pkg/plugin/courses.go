package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"net/http"
)

func handlerCourses(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	db, err := connectDB(settings)
	if err != nil {
		CheckError(err)
	}
	defer db.Close()

	dataFromDatabaseInJson, _ := getCoursesList(db)
	dataJson, _ := json.Marshal(dataFromDatabaseInJson)

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(dataJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getCoursesList(db *sql.DB) ([]json.RawMessage, error) {
	var dataFromDatabaseInJson []json.RawMessage
	var queryStm = "SELECT id, fullname FROM course ORDER BY fullname"
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

		dataFromDatabaseInJson = append(dataFromDatabaseInJson, json.RawMessage(data))
	}

	if dataFromDatabaseInJson == nil {
		dataFromDatabaseInJson = []json.RawMessage{}
	}

	return dataFromDatabaseInJson, err
}
