module github.com/2637309949/bulrush_template

go 1.12

// ## just for dev
replace github.com/2637309949/bulrush => ../bulrush

replace github.com/2637309949/bulrush-openapi => ../bulrush-openapi

replace github.com/2637309949/bulrush-addition => ../bulrush-addition

replace github.com/2637309949/bulrush-limit => ../bulrush-limit

// ## end

require (
	github.com/2637309949/bulrush v0.0.0-20190615094031-919971fe3950
	github.com/2637309949/bulrush-addition v0.0.0-20190608111855-725ea7bbda0e
	github.com/2637309949/bulrush-captcha v0.0.0-20190530105249-dbc3036c68e7
	github.com/2637309949/bulrush-delivery v0.0.0-20190530105510-8228540ca372
	github.com/2637309949/bulrush-identify v0.0.0-20190531073415-116f7ee8683f
	github.com/2637309949/bulrush-limit v0.0.0-20190604090049-8f1a94e98d64
	github.com/2637309949/bulrush-logger v0.0.0-20190605130005-aa61b04c2689
	github.com/2637309949/bulrush-openapi v0.0.0-00010101000000-000000000000
	github.com/2637309949/bulrush-proxy v0.0.0-20190530120941-f91dc1645076
	github.com/2637309949/bulrush-role v0.0.0-20190614132931-262fdeaab12d
	github.com/2637309949/bulrush-template v0.0.0-20190614113343-2d56eb7c18b1
	github.com/2637309949/bulrush-upload v0.0.0-20190531070413-49b362a188be
	github.com/gin-contrib/cache v1.1.1-0.20190528084033-1eca46a236ea
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/kataras/go-events v0.0.2
	github.com/mojocn/base64Captcha v0.0.0-20180423022535-7d0b78ad1685
	github.com/stretchr/testify v1.3.0
	github.com/thoas/go-funk v0.4.0
)
