package config

import (
	"os"
)

type Config struct {
	Addr string

	SettingsPath                    string
	SettingsTmpPath                 string
	SettingsBackupPath              string
	SettingsBackupTmpPath           string
	SettingsAdditionalBackupPath    string
	SettingsAdditionalBackupTmpPath string

	RollbackEndpointPath  string
	EndpointsBackupBefore []string

	JsonPrefix string
	JsonIndent string

	FilePerm   os.FileMode
	TimeFormat string

	ResponseContextKey string
}

func Load() *Config {
	return &Config{
		Addr: ":8000",

		SettingsPath:                    "/etc/moodle-monitoring/settings.json",
		SettingsTmpPath:                 "/etc/moodle-monitoring/settings.json.tmp",
		SettingsBackupPath:              "/etc/moodle-monitoring/settings.prev.json",
		SettingsBackupTmpPath:           "/etc/moodle-monitoring/settings.prev.json.tmp",
		SettingsAdditionalBackupPath:    "/etc/moodle-monitoring/settings.prev.add.json",
		SettingsAdditionalBackupTmpPath: "/etc/moodle-monitoring/settings.prev.add.json.tmp",

		RollbackEndpointPath: "/api/rollback",
		EndpointsBackupBefore: []string{
			"PATCH:/api/settings",
		},

		JsonPrefix: "",
		JsonIndent: "    ",

		FilePerm:   0644,
		TimeFormat: "02-01-2006 15:04:05",

		ResponseContextKey: "response",
	}
}
