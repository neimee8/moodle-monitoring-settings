package api

import (
	"net/http"
	"settings/internal/engine"
	"settings/internal/types"
)

func (h ApiHandlers) HandleRollback(w http.ResponseWriter, r *http.Request) {
	resp, ok := r.Context().Value(h.cfg.ResponseContextKey).(*types.ApiResponse)

	if !ok {
		return
	}

	cmd := engine.CmdRollback(h.cfg, r.URL.Path)
	resCh := cmd.RespCh
	h.cmdCh <- cmd
	res := <-resCh

	resp.Status = res.Status
	resp.ResponseData.Msg = res.Msg
	resp.ResponseData.Data = res.Data
}
