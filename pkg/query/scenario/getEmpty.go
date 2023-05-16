package scenario

import (
	"eficacia-datatable-backend/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func getEmptyFrame(query backend.DataQuery, settings models.PluginSettings) *data.Frame {
	return data.NewFrame("empty")
}
