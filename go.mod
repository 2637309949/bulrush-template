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
	github.com/2637309949/bulrush v0.0.0-20190930081059-e979ad9b6c52
	github.com/2637309949/bulrush-addition v0.0.0-20190927143255-352750ae197b
	github.com/2637309949/bulrush-captcha v0.0.0-20190923102136-af2f3030b733
	github.com/2637309949/bulrush-delivery v0.0.0-20190923102053-8d2c53577dbe
	github.com/2637309949/bulrush-identify v0.0.0-20190923102129-444091e36e80
	github.com/2637309949/bulrush-limit v0.0.0-20190923102225-2cb72c17b190
	github.com/2637309949/bulrush-logger v0.0.0-20190923102217-900277928acb
	github.com/2637309949/bulrush-mq v0.0.0-20190923102247-af31ddf46e92
	github.com/2637309949/bulrush-openapi v0.0.0-20190923102232-c7eee70cda74
	github.com/2637309949/bulrush-proxy v0.0.0-20190923102240-74307b36eca9
	github.com/2637309949/bulrush-role v0.0.0-20190923102203-c1acfe6d1657
	github.com/2637309949/bulrush-upload v0.0.0-20190923102107-5f2d3990d9e1
	github.com/2637309949/bulrush-utils v0.0.0-20190926030656-d38f3aa4b69a
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/kataras/go-events v0.0.2
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.4.0
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc
	golang.org/x/net v0.0.0-20191003171128-d98b1b443823
	google.golang.org/grpc v1.24.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/guregu/null.v3 v3.4.0
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5
	gopkg.in/yaml.v2 v2.2.4
)
