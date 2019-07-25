## Usage
```go
package main
import (
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/kataras/go-events"
)

func main() {
	app := InitApp()
	app.Use(func(event events.EventEmmiter) {
		event.On(bulrush.EventSysBulrushPluginRunImmediately, func(message ...interface{}) {
			addition.Logger.Info("EventSysBulrushPluginRunImmediately %v", message)
		})
	})
	app.RunImmediately()
}
```

### For Dev
```shell
# Hot-start 
$ make -f Makefile.dep
$ make -f Makefile.dev
```

```shell
# Not Hot-start 
$ go run $(ls -1 *.go | grep -v _test.go)
```

### For Prod
```shell
$ make
```

### For Apidoc
![Bulrush flash](./assets/apidoc.png)

```shell
$ make -f Makefile.api
```

### Run Test
```shell
/usr/local/go/bin/go test -timeout 30s github.com/2637309949/bulrush-template -run "^(TestCache)$"
```
Or run with log
```shell
/usr/local/go/bin/go test -timeout 30s github.com/2637309949/bulrush-template -run "^(TestCache)$" -v
```

##  Standard
### Http status and Http reponse
#### Reponse Success
```go
c.JSON(http.StatusOK, users)
```

#### Reponse Error
```go
c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	"message": "Internal Server Error",
	"stack":   "no id found",
})
```

### Global object
```go
//./addition
func Request(method string, url string, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
```

### Logger
```go
func RegisterUser(r *gin.RouterGroup, role *role.Role) {
	addition.MGOExt.API.List(r, "User").Pre(func(c *gin.Context) {
		addition.Logger.Info("after")
	}).Auth(func(c *gin.Context) bool {
		return true
	})
	addition.MGOExt.API.Feature("feature").List(r, "User")
	addition.MGOExt.API.One(r, "User", role.Can("r1,r2@p1,p3,p4;r4"))
	addition.MGOExt.API.Create(r, "User")
	addition.MGOExt.API.Update(r, "User")
	addition.MGOExt.API.Delete(r, "User")
}
```

## MIT License

Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.