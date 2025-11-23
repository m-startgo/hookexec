package flog

import (
	"log/slog"
	"time"

	"github.com/m-startgo/go-utils/mcycle"
	"github.com/m-startgo/go-utils/mlog"
)

var App *mlog.Logger

func LogInit() {
	App = mlog.New(mlog.Config{
		Path:   "./logs/app",
		Name:   "app",
		Level:  slog.LevelDebug,
		Stdout: true,
	})

	cy := mcycle.New(mcycle.Options{
		Task: func() {
			App.Info("flog.LogInit", "清理一次日志文件")
			err := App.Clear(mlog.ClearOpt{}) // 清理日志文件
			if err != nil {
				App.Error("flog.LogInit", "清理出错", err)
			}
		},
		Interval: time.Hour * 24,
	})
	cy.Start()
}
