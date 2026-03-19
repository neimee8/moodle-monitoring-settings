package api

import (
	"encoding/json"
	"net/http"
	"settings/internal/engine"
	"settings/internal/types"
)

func (h ApiHandlers) HandleSettings(w http.ResponseWriter, r *http.Request) {
	resp, ok := r.Context().Value(h.cfg.ResponseContextKey).(*types.ApiResponse)

	if !ok {
		return
	}

	var cmd engine.Cmd

	switch r.Method {
	case http.MethodGet:
		cmd = engine.CmdGet(h.cfg)

	case http.MethodPatch:
		var params map[string]json.RawMessage
		err := json.NewDecoder(r.Body).Decode(&params)

		if err != nil {
			resp.Status = 400
			resp.ResponseData.Msg = "error while parsing json: " + err.Error()

			return
		}

		cmd = engine.CmdModify(h.cfg, params)

	default:
		resp.Status = 405
		resp.ResponseData.Msg = "method not allowed"

		return
	}

	resCh := cmd.RespCh
	h.cmdCh <- cmd
	res := <-resCh

	resp.Status = res.Status
	resp.ResponseData.Msg = res.Msg
	resp.ResponseData.Data = res.Data
}
