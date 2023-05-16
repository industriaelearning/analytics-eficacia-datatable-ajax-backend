package models

import (
	"encoding/json"
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type PluginSettings struct {
	PostgresHost     string                `json:"postgresHost"`
	PostgresPort     string                `json:"postgresPort"`
	PostgresUsername string                `json:"postgresUsername"`
	PostgresDatabase string                `json:"postgresDatabase"`
	Secrets          *SecretPluginSettings `json:"-"`
}

// SecretPluginSettings contains the plugin settings that are encrypted in the database.
type SecretPluginSettings struct {
	PostgresPassword string
}

func LoadPluginSettings(source backend.DataSourceInstanceSettings) (*PluginSettings, error) {
	if source.JSONData == nil || len(source.JSONData) < 1 {
		// If no settings have been saved return default values
		return &PluginSettings{
			PostgresHost:     "",
			PostgresPort:     "",
			PostgresUsername: "",
			PostgresDatabase: "",
			Secrets:          LoadSecretPluginSettings(source.DecryptedSecureJSONData),
		}, nil
	}

	settings := PluginSettings{
		PostgresHost:     "postgresHost",
		PostgresPort:     "postgresPort",
		PostgresUsername: "postgresUsername",
		PostgresDatabase: "postgresDatabase",
	}

	err := json.Unmarshal(source.JSONData, &settings)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal PluginSettings json: %w", err)
	}

	settings.Secrets = LoadSecretPluginSettings(source.DecryptedSecureJSONData)

	return &settings, nil
}

func LoadSecretPluginSettings(source map[string]string) *SecretPluginSettings {
	return &SecretPluginSettings{
		PostgresPassword: source["postgresPassword"],
	}
}
