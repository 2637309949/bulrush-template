package addition

import (
	"path"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/2637309949/bulrush-addition/apidoc"
	gormext "github.com/2637309949/bulrush-addition/gorm"
	"github.com/2637309949/bulrush-addition/logger"
	mgoext "github.com/2637309949/bulrush-addition/mgo"
	"github.com/2637309949/bulrush-addition/redis"
	"github.com/2637309949/bulrush-template/conf"
	"github.com/2637309949/bulrush-template/utils"
	"github.com/gin-gonic/gin"
)

// Logger application logger
var Logger = addition.RushLogger.AppendTransports([]*logger.Transport{
	// only for error
	&logger.Transport{
		Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "error"),
		Level:   logger.ERROR,
		Maxsize: logger.Maxsize,
	},
	// combined all level
	&logger.Transport{
		Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "combined"),
		Level:   logger.SILLY,
		Maxsize: logger.Maxsize,
	},
	// console level
	// &logger.Transport{
	// 	Level: logger.SILLY,
	// },
}...)

// GORMExt application mongo store
var GORMExt = gormext.
	New().
	Init(func(ext *gormext.GORM) {
		var cfg = &gormext.Config{}
		if err := conf.Cfg.Unmarshal("sql", cfg); err != nil {
			panic(err)
		}
		ext.Conf(cfg)
		ext.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
		ext.DB.LogMode(true)
		ext.API.Opts.Prefix = "/template/gorm"
		ext.API.Opts.RouteHooks = &gormext.RouteHooks{
			List: &gormext.ListHook{
				Pre: func(c *gin.Context) {
					Logger.Info("all gormext before")
				},
			},
		}
	})

// MGOExt application mongo store
var MGOExt = mgoext.
	New().
	Init(func(ext *mgoext.Mongo) {
		var cfg = &mgoext.Config{}
		if err := conf.Cfg.Unmarshal("mongo", cfg); err != nil {
			panic(err)
		}
		ext.Conf(cfg)
		ext.API.Opts.Prefix = "/template/mgo"
		ext.API.Opts.RouteHooks = &mgoext.RouteHooks{
			List: &mgoext.ListHook{
				Pre: func(c *gin.Context) {
					Logger.Info("all mgo before")
				},
			},
		}
	})

// APIDoc defined http rest api
var APIDoc = apidoc.
	New().
	Doc(path.Join("", "doc")).
	Init(func(api *apidoc.APIDoc) {
		api.Prefix = "/docs"
		api.GORMExt = GORMExt
		api.MGOExt = MGOExt
	})

// I18N defined global I18N
var I18N = addition.
	NewI18N().
	Init(func(i18n *addition.I18N) func() {
		return func() {
			i18n.AddLocales(addition.M{
				"zh_CN": addition.M{
					"sys_hello": "你好",
				},
				"en_US": addition.M{
					"sys_hello": "hello",
				},
				"zh_TW": addition.M{
					"sys_hello": "妳好",
				},
			})
		}
	})

// Redis application redis store
var Redis = redis.New(conf.Cfg)
