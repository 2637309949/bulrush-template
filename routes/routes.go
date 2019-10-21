// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"time"

	"github.com/2637309949/bulrush"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-addition/logger"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-contrib/cache/persistence"
	"github.com/kataras/go-events"
)

// Routes defined Routes struct
type Routes struct {
	GORMExt     *gormext.GORM
	Persistence *persistence.InMemoryStore
	Logger      *logger.Journal
}

// Plugin defined Plugin
func (r *Routes) Plugin(event events.EventEmmiter, ri *bulrush.ReverseInject) {
	ri.Register(r.testAddUser)
	ri.Register(r.testChache)
	ri.Register(r.testSeq)
	ri.Register(r.testMQHello)
	ri.Register(r.testEvent)
	ri.Register(r.testMockGormLogin)
	ri.Register(r.testMockInit)
	ri.Register(r.testGrpc)
	event.Emit("hello", "this is my payload to hello router")
}

// NewRoutes defined new Routes
func NewRoutes() *Routes {
	return &Routes{
		Persistence: persistence.NewInMemoryStore(time.Second),
		GORMExt:     addition.GORMExt,
		Logger:      addition.Logger,
	}
}
