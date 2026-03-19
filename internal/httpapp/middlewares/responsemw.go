package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"settings/internal/config"
	"settings/internal/types"
	"strings"
)

type ResponseMw struct {
	cfg *config.Config
}

func NewResponseMw(cfg *config.Config) ResponseMw {
	return ResponseMw{
		cfg: cfg,
	}
}

func (mw ResponseMw) Mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		urlParts := strings.Split(r.URL.Path, "/")

		if len(urlParts) > 1 && urlParts[1] == "api" {
			resp := types.NewApiResponse()
			ctx = context.WithValue(r.Context(), mw.cfg.ResponseContextKey, resp)
		}

		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)

		if urlParts[1] == "api" {
			resp, ok := req.Context().Value(mw.cfg.ResponseContextKey).(*types.ApiResponse)

			if !ok {
				resp = types.NewApiResponse()
				resp.Status = 500
				resp.ResponseData.Msg = "response middleware: error while parsing response"
			}

			respJson, err := json.MarshalIndent((*resp.ResponseData), mw.cfg.JsonPrefix, mw.cfg.JsonIndent)

			if err != nil {
				resp.Status = 500
				respJson = []byte(fmt.Sprintf("{\"msg\": \"response middleware: json marshal error: %s\", \"data\": null}", err.Error()))
			}

			w.WriteHeader(resp.Status)
			w.Write(respJson)
		}
	})
}
