package engine

import (
	"settings/internal/config"
)

type Cmd struct {
	handler cmdHandler
	cfg     *config.Config
	params  any
	RespCh  chan Resp
}

func CmdGet(cfg *config.Config) Cmd {
	return Cmd{
		handler: handleGet,
		cfg:     cfg,
		params:  nil,
		RespCh:  make(chan Resp, 1),
	}
}

func CmdModify(cfg *config.Config, params any) Cmd {
	return Cmd{
		handler: handleModify,
		cfg:     cfg,
		params:  params,
		RespCh:  make(chan Resp, 1),
	}
}

func CmdBackup(cfg *config.Config) Cmd {
	return Cmd{
		handler: handleBackup,
		cfg:     cfg,
		params:  nil,
		RespCh:  make(chan Resp, 1),
	}
}

func CmdRollback(cfg *config.Config, params any) Cmd {
	return Cmd{
		handler: handleRollback,
		cfg:     cfg,
		params:  params,
		RespCh:  make(chan Resp, 1),
	}
}

func (cmd Cmd) Handle() {
	cmd.handler(
		cmd.cfg,
		cmd.params,
		cmd.RespCh,
	)
}
