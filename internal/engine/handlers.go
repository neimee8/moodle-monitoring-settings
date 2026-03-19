package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"settings/internal/config"
	"settings/internal/utils"
)

type cmdHandler func(cfg *config.Config, params any, respCh chan Resp)

func handleGet(cfg *config.Config, _ any, respCh chan Resp) {
	data, err := os.ReadFile(cfg.SettingsPath)

	if err != nil {
		respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
		return
	}

	respCh <- NewResp("ok", json.RawMessage(data))
}

func handleModify(cfg *config.Config, params any, respCh chan Resp) {
	if params == nil {
		respCh <- NewResp("params not given", nil, 500)
		return
	}

	paramsCasted, ok := params.(map[string]json.RawMessage)

	if !ok {
		respCh <- NewResp("error while parsing params: map[string]string expected", nil, 400)
		return
	}

	data, err := os.ReadFile(cfg.SettingsPath)

	if err != nil {
		respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
		return
	}

	var settings map[string]json.RawMessage
	err = json.Unmarshal(data, &settings)

	if err != nil {
		respCh <- NewResp("unmarshal json error: "+err.Error(), nil, 500)
		return
	}

	for k, v := range paramsCasted {
		if _, ok := settings[k]; !ok {
			continue
		}

		if !json.Valid([]byte(v)) {
			respCh <- NewResp(fmt.Sprintf("invalid json for key %s: %s", k, v), nil, 400)
			return
		}

		settings[k] = v
	}

	newData, err := json.MarshalIndent(
		settings,
		cfg.JsonPrefix,
		cfg.JsonIndent,
	)

	if err != nil {
		respCh <- NewResp("marshal json error: "+err.Error(), nil, 500)
		return
	}

	err = utils.AtomicWrite(
		cfg.SettingsPath,
		cfg.SettingsTmpPath,
		newData,
		cfg.FilePerm,
	)

	if err != nil {
		respCh <- NewResp("error while writing a file: "+err.Error(), nil, 500)
		return
	}

	respCh <- NewResp("ok", json.RawMessage(newData))
}

func handleBackup(cfg *config.Config, _ any, respCh chan Resp) {
	data, err := os.ReadFile(cfg.SettingsPath)

	if err != nil {
		respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
		return
	}

	err = utils.AtomicWrite(
		cfg.SettingsBackupPath,
		cfg.SettingsBackupTmpPath,
		data,
		cfg.FilePerm,
	)

	if err != nil {
		respCh <- NewResp("error while writing a file: "+err.Error(), nil, 500)
		return
	}

	respCh <- NewResp("ok", json.RawMessage(data))
}

func handleRollback(cfg *config.Config, params any, respCh chan Resp) {
	paramsCasted, ok := params.(string)

	if !ok {
		respCh <- NewResp("error while parsing params: string expected", nil, 400)
		return
	}

	if paramsCasted == cfg.RollbackEndpointPath {
		data, err := os.ReadFile(cfg.SettingsPath)

		if err != nil {
			respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
			return
		}

		err = utils.AtomicWrite(
			cfg.SettingsAdditionalBackupPath,
			cfg.SettingsAdditionalBackupTmpPath,
			data,
			cfg.FilePerm,
		)

		if err != nil {
			respCh <- NewResp("error while writing a file: "+err.Error(), nil, 500)
			return
		}
	}

	data, err := os.ReadFile(cfg.SettingsBackupPath)

	if err != nil {
		respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
		return
	}

	err = utils.AtomicWrite(
		cfg.SettingsPath,
		cfg.SettingsTmpPath,
		data,
		cfg.FilePerm,
	)

	if err != nil {
		respCh <- NewResp("error while writing a file: "+err.Error(), nil, 500)
		return
	}

	if paramsCasted == cfg.RollbackEndpointPath {
		data, err := os.ReadFile(cfg.SettingsAdditionalBackupPath)

		if err != nil {
			respCh <- NewResp("error while reading a file: "+err.Error(), nil, 500)
			return
		}

		err = utils.AtomicWrite(
			cfg.SettingsBackupPath,
			cfg.SettingsBackupTmpPath,
			data,
			cfg.FilePerm,
		)

		if err != nil {
			respCh <- NewResp("error while writing a file: "+err.Error(), nil, 500)
			return
		}

		os.Remove(cfg.SettingsAdditionalBackupPath)
	}

	respCh <- NewResp("ok", json.RawMessage(data))
}
