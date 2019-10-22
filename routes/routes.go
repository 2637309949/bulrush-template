// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"time"

	"github.com/2637309949/bulrush"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-addition/logger"
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/services"
	"github.com/gin-contrib/cache/persistence"
	"github.com/kataras/go-events"
)

// Routes defined Routes struct
type Routes struct {
	GORMExt     *gormext.GORM
	MGOExt      *mgoext.Mongo
	SRV         *services.Services
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
var NewRoutes *Routes

func init() {
	NewRoutes = &Routes{
		Persistence: persistence.NewInMemoryStore(time.Second),
		GORMExt:     addition.GORMExt,
		MGOExt:      addition.MGOExt,
		Logger:      addition.Logger,
		SRV:         services.NewServices,
	}
}
