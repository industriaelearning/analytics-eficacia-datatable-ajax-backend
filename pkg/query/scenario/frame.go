package scenario

import (
	"eficacia-datatable-backend/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func NewDataFrame(query backend.DataQuery, settings models.PluginSettings) *data.Frame {
	switch query.QueryType {
	case GetEmpty:
		return getEmptyFrame(query, settings)
	}

	return nil
}
