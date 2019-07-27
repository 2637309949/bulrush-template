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
	github.com/2637309949/bulrush-addition v0.0.0-20190726130144-51b2fc84c913
	github.com/2637309949/bulrush-captcha v0.0.0-20190726093941-53fcd21bcb22
	github.com/2637309949/bulrush-delivery v0.0.0-20190726093946-81d54c780c21
	github.com/2637309949/bulrush-identify v0.0.0-20190726093950-96fa883b719a
	github.com/2637309949/bulrush-limit v0.0.0-20190726093956-2d15f3b9294c
	github.com/2637309949/bulrush-logger v0.0.0-20190726094001-4037217bd745
	github.com/2637309949/bulrush-mq v0.0.0-20190726094006-45a5b16f70e2
	github.com/2637309949/bulrush-openapi v0.0.0-20190726094010-02b417a6a82a
	github.com/2637309949/bulrush-proxy v0.0.0-20190726094013-417687dbce49
	github.com/2637309949/bulrush-role v0.0.0-20190726094017-3601b3644fb2
	github.com/2637309949/bulrush-upload v0.0.0-20190726094027-afb5f4f40432
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/jinzhu/gorm v1.9.10
	github.com/kataras/go-events v0.0.2
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
)
