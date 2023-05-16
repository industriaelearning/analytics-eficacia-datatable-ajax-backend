package plugin

import (
	"compress/gzip"
	"eficacia-datatable-backend/pkg/models"
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"net/http"
)

func handleQueryDownloads(w http.ResponseWriter, r *http.Request, settings *models.PluginSettings) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=reporte_eficacia.csv.gz")
	w.Header().Set("Content-Type", "application/gzip")

	// Create a new gzip writer to compress the CSV data
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	// Create a new CSV writer to write the data to the gzip writer
	csvWriter := csv.NewWriter(transform.NewWriter(gzipWriter, charmap.ISO8859_1.NewEncoder()))
	defer csvWriter.Flush()

	// Write the header row to the CSV file
	headerRow := getHeaderRow()
	csvWriter.Write(headerRow)

	db, err := connectDB(settings)
	if err != nil {
		CheckError(err)
	}
	defer db.Close()

	data := r.URL.Query()
	columns := getColumnsFromQuery(data)
	whereClause, _ := createWhereClausule(columns)

	queryStm := fmt.Sprintf("%s %s", getSelectStatement(), whereClause)
	rows, err := db.QueryContext(r.Context(), queryStm)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var cedula, nombre, telefono, email, cargo, regional, ciudad, area, cliente, jefeInmediato, correoJefeInmediato, telefonoJefe, canal, subcanal, catchupNombre, catchupEmail, timecreated, timesuspended, activo, lastaccess, numeroIngresos, ultimoIngresoCurso, organizacion, programa, departamento, departamentoid, idGrupo, grupoid, agrupacion, fechaHabilitacion, fechaDeshabilitacion, estadoGrupo, satisfaccionPromedioTotal, idCurso, nombreCapacitacion, fechaCompletado, fechaQuiz, transferenciaTotal, intentos, poblacion, pais, ot, nit, tenantid, programid, suspended, estadoManipulacion, fechaInicioAsignacionPrograma, fechaFinalAsignacionPrograma, campo1, campo2, campo3 string

		rows.Scan(&cedula, &nombre, &telefono, &email, &cargo, &regional, &ciudad, &area, &cliente, &jefeInmediato, &correoJefeInmediato, &telefonoJefe, &canal, &subcanal, &catchupNombre, &catchupEmail, &timecreated, &timesuspended, &activo, &lastaccess, &numeroIngresos, &ultimoIngresoCurso, &organizacion, &programa, &departamento, &departamentoid, &idGrupo, &grupoid, &agrupacion, &fechaHabilitacion, &fechaDeshabilitacion, &estadoGrupo, &satisfaccionPromedioTotal, &idCurso, &nombreCapacitacion, &fechaCompletado, &fechaQuiz, &transferenciaTotal, &intentos, &poblacion, &pais, &ot, &nit, &tenantid, &programid, &suspended, &estadoManipulacion, &fechaInicioAsignacionPrograma, &fechaFinalAsignacionPrograma, &campo1, &campo2, &campo3)

		dataRow := []string{
			"'" + cedula + "'",
			"'" + nombre + "'",
			"'" + telefono + "'",
			"'" + email + "'",
			"'" + cargo + "'",
			"'" + regional + "'",
			"'" + ciudad + "'",
			"'" + area + "'",
			"'" + cliente + "'",
			"'" + jefeInmediato + "'",
			"'" + correoJefeInmediato + "'",
			"'" + telefonoJefe + "'",
			"'" + canal + "'",
			"'" + subcanal + "'",
			"'" + catchupNombre + "'",
			"'" + catchupEmail + "'",
			"'" + timecreated + "'",
			"'" + timesuspended + "'",
			"'" + activo + "'",
			"'" + lastaccess + "'",
			"'" + numeroIngresos + "'",
			"'" + ultimoIngresoCurso + "'",
			"'" + organizacion + "'",
			"'" + programa + "'",
			"'" + departamento + "'",
			"'" + departamentoid + "'",
			"'" + idGrupo + "'",
			"'" + grupoid + "'",
			"'" + agrupacion + "'",
			"'" + fechaHabilitacion + "'",
			"'" + fechaDeshabilitacion + "'",
			"'" + estadoGrupo + "'",
			"'" + satisfaccionPromedioTotal + "'",
			"'" + idCurso + "'",
			"'" + nombreCapacitacion + "'",
			"'" + fechaCompletado + "'",
			"'" + fechaQuiz + "'",
			"'" + transferenciaTotal + "'",
			"'" + intentos + "'",
			"'" + poblacion + "'",
			"'" + pais + "'",
			"'" + ot + "'",
			"'" + nit + "'",
			"'" + tenantid + "'",
			"'" + programid + "'",
			"'" + suspended + "'",
			"'" + estadoManipulacion + "'",
			"'" + fechaInicioAsignacionPrograma + "'",
			"'" + fechaFinalAsignacionPrograma + "'",
			"'" + campo1 + "'",
			"'" + campo2 + "'",
			"'" + campo3 + "'",
		}

		csvWriter.Write(dataRow)

	}

	gzipWriter.Flush()

}

func getHeaderRow() []string {
	return []string{
		"'CEDULA'",
		"'NOMBRE'",
		"'TELEFONO'",
		"'CORREO ELECTRONICO DEL USUARIO'",
		"'CARGO PROPIO DEL CLIENTE/ PROCESO'",
		"'REGIONAL'",
		"'CIUDAD'",
		"'AREA'",
		"'CLIENTE'",
		"'NOMBRE DEL JEFE INMEDIATO'",
		"'CORREO ELECTRONICO DEL JEFE INMEDIATO'",
		"'TELEFONO JEFE'",
		"'CANAL'",
		"'SUBCANAL'",
		"'NOMBRE DEL CATCHUP'",
		"'CORREO ELECTRONICO CATCHUP'",
		"'FECHA DE CREACION DEL USUARIO'",
		"'FECHA ACTIVACION/FECHA_DESACTIVACION'",
		"'ESTADO PLATAFORMA'",
		"'ULTIMA FECHA DE INGRESO A LA PLATAFORMA'",
		"'NUMERO DE INGRESOS'",
		"'ULTIMO INGRESO AL CURSO'",
		"'ORGANIZACION'",
		"'NOMBRE PROGRAMA'",
		"'DEPARTAMENTO'",
		"'DEPARTAMENTOID'",
		"'ID GRUPO'",
		"'GRUPO ID'",
		"'AGRUPACION'",
		"'FECHA HABILITACION'",
		"'FECHA DESHABILITACION'",
		"'ESTADO GRUPO'",
		"'SATISFACCIÓN PROMEDIO TOTAL'",
		"'ID CURSO'",
		"'NOMBRE DE LA CAPACITACION'",
		"'FECHA COMPLETADO'",
		"'FECHA QUIZ'",
		"'TRANSFERENCIA PROMEDIO TOTAL'",
		"'INTENTOS'",
		"'POBLACION'",
		"'PAIS'",
		"'OT'",
		"'NIT'",
		"'TENANTID'",
		"'PROGRAMID'",
		"'SUSPENDIDO'",
		"'ESTADO MANIPULACION'",
		"'FECHA INICIO ASIGNACION PROGRAMA'",
		"'FECHA FINAL ASIGNACION PROGRAMA'",
		"'CAMPO 1'",
		"'CAMPO 2'",
		"'CAMPO 3'",
	}
}
