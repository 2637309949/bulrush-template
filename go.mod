module github.com/2637309949/bulrush-template

go 1.12

// ## just for dev
replace github.com/2637309949/bulrush => ../bulrush

replace github.com/2637309949/bulrush-openapi => ../bulrush-openapi

replace github.com/2637309949/bulrush-addition => ../bulrush-addition

replace github.com/2637309949/bulrush-limit => ../bulrush-limit

replace github.com/2637309949/bulrush-template => ../bulrush-template

replace github.com/2637309949/bulrush-mq => ../bulrush-mq

replace github.com/2637309949/bulrush-role => ../bulrush-role

replace github.com/2637309949/bulrush-captcha => ../bulrush-captcha

replace github.com/2637309949/bulrush-delivery => ../bulrush-delivery

replace github.com/2637309949/bulrush-upload => ../bulrush-upload

replace github.com/2637309949/bulrush-logger => ../bulrush-logger

replace github.com/2637309949/bulrush-identify => ../bulrush-identify

replace github.com/2637309949/bulrush-proxy => ../bulrush-proxy

replace github.com/2637309949/bulrush-utils => ../bulrush-utils

// ## end

require (
	github.com/2637309949/bulrush v0.0.0-20190826101703-2fac7be21f72
	github.com/2637309949/bulrush-addition v0.0.0-20190831034018-427428781eb0
	github.com/2637309949/bulrush-captcha v0.0.0-20190805060015-0335411e830a
	github.com/2637309949/bulrush-delivery v0.0.0-20190805055946-c208fdca9d47
	github.com/2637309949/bulrush-identify v0.0.0-20190809180431-61d49dc64bae
	github.com/2637309949/bulrush-limit v0.0.0-20190805060044-6476475582a2
	github.com/2637309949/bulrush-logger v0.0.0-20190805060009-4b5c2b58a1b9
	github.com/2637309949/bulrush-mq v0.0.0-20190814163826-12f2e8e8b979
	github.com/2637309949/bulrush-openapi v0.0.0-20190805060026-c5958cb2f053
	github.com/2637309949/bulrush-proxy v0.0.0-20190805060003-15134ca1ade4
	github.com/2637309949/bulrush-role v0.0.0-20190815073833-05a8fe434808
	github.com/2637309949/bulrush-upload v0.0.0-20190805060057-d3f855287492
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jinzhu/gorm v1.9.10
	github.com/kataras/go-events v0.0.2
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/guregu/null.v3 v3.4.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
	gopkg.in/yaml.v2 v2.2.2
)
