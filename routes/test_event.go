// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"

	"github.com/2637309949/bulrush"
	"github.com/kataras/go-events"
)

func onHello(event events.EventEmmiter) {
	event.On("hello", func(payload ...interface{}) {
		message := payload[0].(string)
		fmt.Println(message)
	})
}

// RegisterEvent defined test routes
func RegisterEvent(ri *bulrush.ReverseInject) {
	ri.Register(onHello)
}
