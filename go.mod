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
	github.com/2637309949/bulrush v0.0.0-20190725143958-5a43e012d374
	github.com/2637309949/bulrush-addition v0.0.0-20190726090813-d1b08681a36b
	github.com/2637309949/bulrush-captcha v0.0.0-20190702143725-3743f3a68df2
	github.com/2637309949/bulrush-delivery v0.0.0-20190705025530-46053d43000d
	github.com/2637309949/bulrush-identify v0.0.0-20190725144021-d513740ba708
	github.com/2637309949/bulrush-limit v0.0.0-20190706024522-5f4075c64746
	github.com/2637309949/bulrush-logger v0.0.0-20190723022945-059d00205728
	github.com/2637309949/bulrush-mq v0.0.0-20190712173825-54404462465f
	github.com/2637309949/bulrush-openapi v0.0.0-20190725144026-d1134b71f094
	github.com/2637309949/bulrush-proxy v0.0.0-20190702143820-a3c56976a0a0
	github.com/2637309949/bulrush-role v0.0.0-20190702143825-37a00040b1fb
	github.com/2637309949/bulrush-upload v0.0.0-20190725144035-b6ac7a7f2464
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/jinzhu/gorm v1.9.9
	github.com/kataras/go-events v0.0.2
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
)

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
