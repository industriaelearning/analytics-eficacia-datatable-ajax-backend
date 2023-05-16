package plugin

import (
	"database/sql"
	"eficacia-datatable-backend/pkg/models"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// connectDB connects to the database and returns a *sql.DB instance.
func connectDB(settings *models.PluginSettings) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", settings.PostgresHost, settings.PostgresPort, settings.PostgresUsername, settings.Secrets.PostgresPassword, settings.PostgresDatabase)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getColumnsFromQuery(data url.Values) []map[string]interface{} {
	var columns []map[string]interface{}
	for i := 0; ; i++ {
		index := fmt.Sprintf("columns[%d][", i)
		dataKey := index + "data]"
		nameKey := index + "name]"
		searchableKey := index + "searchable]"
		orderableKey := index + "orderable]"
		valueKey := index + "search][value]"
		regexKey := index + "search][regex]"

		if _, ok := data[dataKey]; !ok {
			break
		}

		column := map[string]interface{}{
			"data":       data.Get(dataKey),
			"name":       data.Get(nameKey),
			"searchable": data.Get(searchableKey),
			"orderable":  data.Get(orderableKey),
			"search": map[string]string{
				"value": data.Get(valueKey),
				"regex": data.Get(regexKey),
			},
		}

		columns = append(columns, column)
	}

	if columns == nil {
		columns = []map[string]interface{}{}
	}

	return columns
}

func getOrderFromQuery(data url.Values) []map[string]interface{} {
	var order []map[string]interface{}
	for i := 0; ; i++ {
		index := fmt.Sprintf("order[%d][", i)
		columnKey := index + "column]"
		dirKey := index + "dir]"

		if _, ok := data[columnKey]; !ok {
			break
		}

		orderItem := map[string]interface{}{
			"column": data.Get(columnKey),
			"dir":    data.Get(dirKey),
		}

		order = append(order, orderItem)
	}

	return order
}

func getSelectStatement() string {
	return "SELECT " +
		"CASE WHEN cedula IS NULL THEN '' ELSE cedula::text END AS cedula, " +
		"CASE WHEN nombre IS NULL THEN '' ELSE nombre::text END AS nombre, " +
		"CASE WHEN telefono IS NULL THEN '' ELSE telefono::text END AS telefono, " +
		"CASE WHEN email IS NULL THEN '' ELSE email::text END AS email, " +
		"CASE WHEN cargo IS NULL THEN '' ELSE cargo::text END AS cargo, " +
		"CASE WHEN regional IS NULL THEN '' ELSE regional::text END AS regional, " +
		"CASE WHEN ciudad IS NULL THEN '' ELSE ciudad::text END AS ciudad, " +
		"CASE WHEN area IS NULL THEN '' ELSE area::text END AS area, " +
		"CASE WHEN cliente IS NULL THEN '' ELSE cliente::text END AS cliente, " +
		"CASE WHEN jefe_inmediato IS NULL THEN '' ELSE jefe_inmediato::text END AS jefe_inmediato, " + /*10*/
		"CASE WHEN correo_jefe_inmediato IS NULL THEN '' ELSE correo_jefe_inmediato::text END AS correo_jefe_inmediato, " +
		"CASE WHEN telefono_jefe IS NULL THEN '' ELSE telefono_jefe::text END AS telefono_jefe, " +
		"CASE WHEN canal IS NULL THEN '' ELSE canal::text END AS canal, " +
		"CASE WHEN subcanal IS NULL THEN '' ELSE subcanal::text END AS subcanal, " +
		"CASE WHEN catchup_nombre IS NULL THEN '' ELSE catchup_nombre::text END AS catchup_nombre, " +
		"CASE WHEN catchup_email IS NULL THEN '' ELSE catchup_email::text END AS catchup_email, " +
		"CASE WHEN timecreated IS NULL THEN '' ELSE timecreated::text END AS timecreated, " +
		"CASE WHEN timesuspended IS NULL THEN '' ELSE timesuspended::text END AS timesuspended, " +
		"CASE WHEN activo IS NULL THEN '' ELSE activo::text END AS activo, " +
		"CASE WHEN lastaccess IS NULL THEN '' ELSE lastaccess::text END AS lastaccess, " + /*20*/
		"CASE WHEN numero_ingresos IS NULL THEN '' ELSE numero_ingresos::text END AS numero_ingresos, " +
		"CASE WHEN ultimo_ingreso_curso IS NULL THEN '' ELSE ultimo_ingreso_curso::text END AS ultimo_ingreso_curso, " +
		"CASE WHEN organizacion IS NULL THEN '' ELSE organizacion::text END AS organizacion, " +
		"CASE WHEN programa IS NULL THEN '' ELSE programa::text END AS programa, " +
		"CASE WHEN departamento IS NULL THEN '' ELSE departamento::text END AS departamento, " +
		"CASE WHEN departamentoid IS NULL THEN '' ELSE departamentoid::text END AS departamentoid, " +
		"CASE WHEN id_grupo IS NULL THEN '' ELSE id_grupo::text END AS id_grupo, " +
		"CASE WHEN grupoid IS NULL THEN '' ELSE grupoid::text END AS grupoid, " +
		"CASE WHEN agrupacion IS NULL THEN '' ELSE agrupacion::text END AS agrupacion, " +
		"CASE WHEN fecha_habilitacion IS NULL THEN '' ELSE fecha_habilitacion::text END AS fecha_habilitacion, " + /*30*/
		"CASE WHEN fecha_deshabilitacion IS NULL THEN '' ELSE fecha_deshabilitacion::text END AS fecha_deshabilitacion, " +
		"CASE WHEN estado_grupo IS NULL THEN '' ELSE estado_grupo::text END AS estado_grupo, " +
		"CASE WHEN satisfaccion_promedio_total IS NULL THEN '' ELSE satisfaccion_promedio_total::text END AS satisfaccion_promedio_total, " +
		"CASE WHEN id_curso IS NULL THEN '' ELSE id_curso::text END AS id_curso, " +
		"CASE WHEN nombre_capacitacion IS NULL THEN '' ELSE nombre_capacitacion::text END AS nombre_capacitacion, " +
		"CASE WHEN fecha_completado IS NULL THEN '' ELSE fecha_completado::text END AS fecha_completado, " +
		"CASE WHEN fecha_quiz IS NULL THEN '' ELSE fecha_quiz::text END AS fecha_quiz, " +
		"CASE WHEN transferencia_total IS NULL THEN '' ELSE transferencia_total::text END AS transferencia_total, " +
		"CASE WHEN intentos IS NULL THEN '' ELSE intentos::text END AS intentos, " +
		"CASE WHEN poblacion IS NULL THEN '' ELSE poblacion::text END AS poblacion, " + /*40*/
		"CASE WHEN pais IS NULL THEN '' ELSE pais::text END AS pais, " +
		"CASE WHEN ot IS NULL THEN '' ELSE ot::text END AS ot, " +
		"CASE WHEN nit IS NULL THEN '' ELSE nit::text END AS nit, " +
		"CASE WHEN tenantid IS NULL THEN '' ELSE tenantid::text END AS tenantid, " +
		"CASE WHEN programid IS NULL THEN '' ELSE programid::text END AS programid, " +
		"CASE WHEN suspended IS NULL THEN '' ELSE suspended::text END AS suspended, " +
		"CASE WHEN estado_manipulacion IS NULL THEN '' ELSE estado_manipulacion::text END AS estado_manipulacion, " +
		"CASE WHEN fecha_inicio_asignacion_programa IS NULL THEN '' ELSE fecha_inicio_asignacion_programa::text END AS fecha_inicio_asignacion_programa, " +
		"CASE WHEN fecha_final_asignacion_programa IS NULL THEN '' ELSE fecha_final_asignacion_programa::text END AS fecha_final_asignacion_programa, " +
		"CASE WHEN campo_1 IS NULL THEN '' ELSE campo_1::text END AS campo_1, " + /*50*/
		"CASE WHEN campo_2 IS NULL THEN '' ELSE campo_2::text END AS campo_2, " +
		"CASE WHEN campo_3 IS NULL THEN '' ELSE campo_3::text END AS campo_3 " +
		"FROM c_reporte "
}

func getDraw(r *http.Request) string {
	draw := r.URL.Query().Get("draw")
	if draw != "" {
		return draw
	}
	return "0"
}

func createWhereClausule(columns []map[string]interface{}) (string, string) {
	var whereClause string
	var tenantClause string
	for i := range columns {
		if searchVal := columns[i]["search"].(map[string]string)["value"]; searchVal != "" {
			if whereClause == "" {
				whereClause = "WHERE "
			} else {
				whereClause += " AND "
			}
			if columns[i]["data"].(string) == "tenantid" {
				tenantClause = fmt.Sprintf(" %s IN (%s) ", columns[i]["data"], searchVal)
				whereClause += fmt.Sprintf(" %s IN (%s) ", columns[i]["data"], searchVal)
			} else if inArray(columns[i]["data"].(string), []string{"programid", "departamentoid", "id_curso", "grupoid", "suspended"}) {
				whereClause += fmt.Sprintf(" %s = %s ", columns[i]["data"], searchVal)
			} else if columns[i]["data"] == "fecha_quiz" {
				dateRange := strings.Split(searchVal, "-delimiter-")
				if len(dateRange) == 2 {
					start, err := strconv.ParseInt(dateRange[0], 10, 64)
					CheckError(err)
					end, err := strconv.ParseInt(dateRange[1], 10, 64)
					CheckError(err)
					whereClause += fmt.Sprintf(" fecha_quiz >= %d AND fecha_quiz <= %d ", start, end)
				}
			} else {
				whereClause += fmt.Sprintf(" %s ILIKE '%%%s%%' ", columns[i]["data"], searchVal)
			}
		}
	}

	if !strings.Contains(whereClause, "tenantid") {
		if whereClause == "" {
			whereClause = "WHERE "
		} else {
			whereClause += " AND "
		}
		tenantClause = fmt.Sprintf(" tenantid = %d ", 1)
		whereClause += tenantClause
	}

	return whereClause, tenantClause
}

func inArray(str string, arr []string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
