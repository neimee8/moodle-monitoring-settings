package middlewares

import (
	"fmt"
	"net/http"
	"settings/internal/config"
	"settings/internal/engine"
	"settings/internal/types"
	"strings"
	"time"
)

type BackupMw struct {
	cfg   *config.Config
	cmdCh chan engine.Cmd
}

func NewBackupMw(cfg *config.Config, cmdCh chan engine.Cmd) BackupMw {
	return BackupMw{
		cfg:   cfg,
		cmdCh: cmdCh,
	}
}

func (mw BackupMw) Mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method
		backedUp := false

		for _, endpoint := range mw.cfg.EndpointsBackupBefore {
			if endpoint == fmt.Sprintf("%s:%s", strings.ToUpper(method), path) ||
				endpoint == "*:"+path {

				backedUp = true

				cmd := engine.CmdBackup(mw.cfg)
				resCh := cmd.RespCh
				mw.cmdCh <- cmd
				res := <-resCh

				if res.Status != 200 {
					resp, ok := r.Context().Value(mw.cfg.ResponseContextKey).(*types.ApiResponse)

					if !ok {
						return
					}

					resp.Status = res.Status
					resp.ResponseData.Msg = res.Msg
					resp.ResponseData.Data = res.Data
				}
			}
		}

		if backedUp {
			fmt.Printf("🕰️ [%s]\n💿 Settings backed up\n\n", time.Now().Format(mw.cfg.TimeFormat))
		}

		next.ServeHTTP(w, r)
	})
}
