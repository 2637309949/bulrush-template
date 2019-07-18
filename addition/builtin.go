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
	"github.com/globalsign/mgo"
)

// Logger defined bulrush or system log global proxy
// Two Transport has been added
// 1: RotateFile
// 2: Console
var Logger = addition.RushLogger.
	AppendTransports([]*logger.Transport{
		&logger.Transport{
			Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "error"),
			Level:   logger.ERROR,
			Maxsize: logger.Maxsize,
		},
		&logger.Transport{
			Dirname: path.Join(path.Join(".", utils.Some(conf.Cfg.Log.Path, "logs").(string)), "combined"),
			Level:   logger.SILLY,
			Maxsize: logger.Maxsize,
		},
	}...).
	Init(func(j *logger.Journal) {
		j.SetFlags((logger.LstdFlags | logger.Lshortfile))
	})

// GORMExt defined ext for gorm
// .API use default api routes
// .Plugin use as bulrush plugin
var GORMExt = gormext.
	New().
	Init(func(ext *gormext.GORM) {
		cfg := &gormext.Config{}
		if err := conf.Cfg.Unmarshal("sql", cfg); err != nil {
			panic(err)
		}
		ext.Conf(cfg)
		// 建议在数据库创建时指定CHARSET, 这里设置后看gorm log并不起效
		ext.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
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

// MGOExt defined ext for mgo
// .API use default api routes
// .Plugin use as bulrush plugin
var MGOExt = mgoext.
	New().
	Init(func(ext *mgoext.Mongo) {
		cfg := &mgo.DialInfo{}
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

// APIDoc defined apidoc for sys
// .Doc confg system api
var APIDoc = apidoc.
	New().
	Doc(path.Join("", "doc")).
	Init(func(api *apidoc.APIDoc) {
		api.Prefix = "/docs"
		api.GORMExt = GORMExt
		api.MGOExt = MGOExt
	})

// I18N defined i18n for sys
// .I18NLocale to get i18n asserts
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
