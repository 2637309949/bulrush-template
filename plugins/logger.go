// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package plugins

import (
	"fmt"
	"time"

	logger "github.com/2637309949/bulrush-logger"
	"github.com/gin-gonic/gin"
)

// Logger Plugin init
var Logger = &logger.Logger{
	Path: "logs",
	Format: func(p *logger.Payload, ctx *gin.Context) string {
		if p.Type == logger.INT {
			startTime := time.Unix(p.StartUnix, 0).Format("2006/01/02 15:04:05")
			return fmt.Sprintf("[%v bulrush] => %s %6s %s", startTime, p.IP, p.Method, p.URL)
		} else if p.Type == logger.OUT {
			endOfTime := time.Unix(p.EndUnix, 0).Format("2006/01/02 15:04:05")
			latency := float64(time.Unix(p.EndUnix, 0).Sub(time.Unix(p.StartUnix, 0)) / time.Millisecond)
			return fmt.Sprintf("[%v bulrush] <= %.2fms %s %6s %s", endOfTime, latency, p.IP, p.Method, p.URL)
		}
		return "FROMAT ERROR"
	},
}
