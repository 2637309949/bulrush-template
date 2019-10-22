package services

import (
	"time"

	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-addition/logger"
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/gin-contrib/cache/persistence"
)

// Services defined Services struct
type Services struct {
	MGOExt      *mgoext.Mongo
	GORMExt     *gormext.GORM
	Persistence *persistence.InMemoryStore
	Logger      *logger.Journal
}

// NewServices defined new Services
var NewServices *Services

func init() {
	NewServices = &Services{
		Persistence: persistence.NewInMemoryStore(time.Second),
		GORMExt:     addition.GORMExt,
		MGOExt:      addition.MGOExt,
		Logger:      addition.Logger,
	}
}
