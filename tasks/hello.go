// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tasks

import (
	"github.com/2637309949/bulrush_template/addition"
	"github.com/kataras/go-events"
	"gopkg.in/robfig/cron.v2"
)

// RegisterHello defined hello routes
// 复杂的业务推荐调度中心, task只完成系统性调度，如：发邮件，推送系统信息等
// Jobrunner 结合 emmiter，高程度解耦

type reminderEmails struct {
	message string
}

func (e reminderEmails) Run() {
	addition.Logger.Info("Every 5 sec send reminder emails \n")
}

// RegisterHello defined hello task
func RegisterHello(schedule *cron.Cron, emmiter events.EventEmmiter) {
	schedule.AddFunc("@every 30s", func() {
		addition.Logger.Info("Every 30 sec send reminder emails \n")
	})
	schedule.Start()
}
