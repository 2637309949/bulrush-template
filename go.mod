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
	github.com/2637309949/bulrush v0.0.0-20190804083528-61e2e5a9d8c2
	github.com/2637309949/bulrush-addition v0.0.0-20190729105901-74c392aece2f
	github.com/2637309949/bulrush-captcha v0.0.0-20190802172107-c30997c91025
	github.com/2637309949/bulrush-delivery v0.0.0-20190802172034-7071a1a31c42
	github.com/2637309949/bulrush-identify v0.0.0-20190802172113-ed4c283b9b5f
	github.com/2637309949/bulrush-limit v0.0.0-20190802172131-952609dfbd69
	github.com/2637309949/bulrush-logger v0.0.0-20190802172101-4b8a609904e5
	github.com/2637309949/bulrush-mq v0.0.0-20190802172040-d0b625dc7f85
	github.com/2637309949/bulrush-openapi v0.0.0-20190802172119-9f02ece2e720
	github.com/2637309949/bulrush-proxy v0.0.0-20190802172053-f44fe65e482a
	github.com/2637309949/bulrush-role v0.0.0-20190804090357-5843cc379f79
	github.com/2637309949/bulrush-upload v0.0.0-20190802172137-073d7df26e46
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/jinzhu/gorm v1.9.10
	github.com/kataras/go-events v0.0.2
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
)
