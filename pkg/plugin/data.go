package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"time"
)

var frameCache = map[string]json.RawMessage{}

const cacheDuration = time.Minute * 5

func handlerData(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	cacheKey := r.URL.RawQuery

	if frame, ok := frameCache[cacheKey]; ok {
		_, err := w.Write(frame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := connectDB(settings)
	if err != nil {
		CheckError(err)
	}
	defer db.Close()

	data := r.URL.Query()
	columns := getColumnsFromQuery(data)

	order := getOrderFromQuery(data)

	start := r.URL.Query().Get("start")
	if start == "" {
		start = "0"
	}

	length := r.URL.Query().Get("length")
	if length == "" {
		length = "5"
	}

	var orderBy string
	if len(columns) > 0 {
		for i := range order {
			columnIndex, _ := strconv.Atoi(order[i]["column"].(string))
			orderColumn := columns[columnIndex]["data"]
			orderDir := order[i]["dir"]
			orderBy = fmt.Sprintf("ORDER BY %s %s", orderColumn, orderDir)
		}
	}

	if orderBy == "" {
		orderBy = "ORDER BY 1 DESC"
	}

	whereClause, tenantClause := createWhereClausule(columns)
	limitClause := fmt.Sprintf("LIMIT %s OFFSET %s", length, start)

	dataFromDatabaseInJson, _ := getDataFromRow(db, whereClause, orderBy, limitClause)
	recordsTotal, _ := getRecordsTotal(db, tenantClause)
	recordsFiltered, _ := getRecordsFiltered(db, whereClause)
	dataJson, _ := json.Marshal(dataFromDatabaseInJson)

	draw := getDraw(r)
	frame := json.RawMessage(
		`{` +
			fmt.Sprintf(`"recordsTotal": %d`, recordsTotal) + `,` +
			fmt.Sprintf(`"recordsFiltered": %d`, recordsFiltered) + `,` +
			fmt.Sprintf(`"draw": %s`, draw) + `,` +
			fmt.Sprintf(`"data": %s`, dataJson) +
			`}`,
	)
	frameCache[cacheKey] = frame

	// Remove frame from cache after cacheDuration has passed
	time.AfterFunc(cacheDuration, func() {
		delete(frameCache, cacheKey)
	})

	_, err = w.Write(frame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getRecordsFiltered(db *sql.DB, whereClause string) (int, error) {
	var recordsFiltered int
	selectStm := "SELECT count(0) FROM public.c_reporte "
	queryStm := fmt.Sprintf("%s %s", selectStm, whereClause)

	err := db.QueryRow(queryStm).Scan(&recordsFiltered)
	CheckError(err)

	return recordsFiltered, nil
}

func getDataFromRow(db *sql.DB, whereClause string, orderBy string, limitClause string) ([]json.RawMessage, error) {
	var dataFromDatabaseInJson []json.RawMessage

	queryStm := fmt.Sprintf("%s %s %s %s", getSelectStatement(), whereClause, orderBy, limitClause)
	rows, err := db.Query(queryStm)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var cedula, nombre, telefono, email, cargo, regional, ciudad, area, cliente, jefeInmediato, correoJefeInmediato, telefonoJefe, canal, subcanal, catchupNombre, catchupEmail, timecreated, timesuspended, activo, lastaccess, numeroIngresos, ultimoIngresoCurso, organizacion, programa, departamento, departamentoid, idGrupo, grupoid, agrupacion, fechaHabilitacion, fechaDeshabilitacion, estadoGrupo, satisfaccionPromedioTotal, idCurso, nombreCapacitacion, fechaCompletado, fechaQuiz, transferenciaTotal, intentos, poblacion, pais, ot, nit, tenantid, programid, suspended, estadoManipulacion, fechaInicioAsignacionPrograma, fechaFinalAsignacionPrograma, campo1, campo2, campo3 string

		err := rows.Scan(&cedula, &nombre, &telefono, &email, &cargo, &regional, &ciudad, &area, &cliente, &jefeInmediato, &correoJefeInmediato, &telefonoJefe, &canal, &subcanal, &catchupNombre, &catchupEmail, &timecreated, &timesuspended, &activo, &lastaccess, &numeroIngresos, &ultimoIngresoCurso, &organizacion, &programa, &departamento, &departamentoid, &idGrupo, &grupoid, &agrupacion, &fechaHabilitacion, &fechaDeshabilitacion, &estadoGrupo, &satisfaccionPromedioTotal, &idCurso, &nombreCapacitacion, &fechaCompletado, &fechaQuiz, &transferenciaTotal, &intentos, &poblacion, &pais, &ot, &nit, &tenantid, &programid, &suspended, &estadoManipulacion, &fechaInicioAsignacionPrograma, &fechaFinalAsignacionPrograma, &campo1, &campo2, &campo3)
		CheckError(err)

		data, err := json.Marshal(map[string]interface{}{
			"cedula":                           cedula,
			"nombre":                           nombre,
			"telefono":                         telefono,
			"email":                            email,
			"cargo":                            cargo,
			"regional":                         regional,
			"ciudad":                           ciudad,
			"area":                             area,
			"cliente":                          cliente,
			"jefe_inmediato":                   jefeInmediato,
			"correo_jefe_inmediato":            correoJefeInmediato,
			"telefono_jefe":                    telefonoJefe,
			"canal":                            canal,
			"subcanal":                         subcanal,
			"catchup_nombre":                   catchupNombre,
			"catchup_email":                    catchupEmail,
			"timecreated":                      timecreated,
			"timesuspended":                    timesuspended,
			"activo":                           activo,
			"lastaccess":                       lastaccess,
			"numero_ingresos":                  numeroIngresos,
			"ultimo_ingreso_curso":             ultimoIngresoCurso,
			"organizacion":                     organizacion,
			"programa":                         programa,
			"departamento":                     departamento,
			"departamentoid":                   departamentoid,
			"id_grupo":                         idGrupo,
			"grupoid":                          grupoid,
			"agrupacion":                       agrupacion,
			"fecha_habilitacion":               fechaHabilitacion,
			"fecha_deshabilitacion":            fechaDeshabilitacion,
			"estado_grupo":                     estadoGrupo,
			"satisfaccion_promedio_total":      satisfaccionPromedioTotal,
			"id_curso":                         idCurso,
			"nombre_capacitacion":              nombreCapacitacion,
			"fecha_completado":                 fechaCompletado,
			"fecha_quiz":                       fechaQuiz,
			"transferencia_total":              transferenciaTotal,
			"intentos":                         intentos,
			"poblacion":                        poblacion,
			"pais":                             pais,
			"ot":                               ot,
			"nit":                              nit,
			"tenantid":                         tenantid,
			"programid":                        programid,
			"suspended":                        suspended,
			"estado_manipulacion":              estadoManipulacion,
			"fecha_inicio_asignacion_programa": fechaInicioAsignacionPrograma,
			"fecha_final_asignacion_programa":  fechaFinalAsignacionPrograma,
			"campo1":                           campo1,
			"campo2":                           campo2,
			"campo3":                           campo3,
		})
		CheckError(err)

		dataFromDatabaseInJson = append(dataFromDatabaseInJson, data)
	}

	if dataFromDatabaseInJson == nil {
		dataFromDatabaseInJson = make([]json.RawMessage, 0)
	}

	return dataFromDatabaseInJson, nil
}

func getRecordsTotal(db *sql.DB, whereClause string) (int64, error) {
	var recordsTotal int64
	selectStm := fmt.Sprintf("SELECT count(0) FROM public.c_reporte WHERE %s ", whereClause)
	row := db.QueryRow(selectStm)
	err := row.Scan(&recordsTotal)
	CheckError(err)
	return recordsTotal, nil
}
