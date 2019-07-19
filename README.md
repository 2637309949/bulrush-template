## Usage

```go
// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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

### 1. For Dev
```shell
# Hot-start 
$ make -f Makefile.dep
$ make -f Makefile.dev
```

```shell
# Not Hot-start 
$ go run $(ls -1 *.go | grep -v _test.go)
```

### 2. For Prod

```shell
$ make
```
    // Remove those dev lines in go.mod file
    // ## just for dev
    replace github.com/2637309949/bulrush => ../bulrush
    replace github.com/2637309949/bulrush-openapi => ../bulrush-openapi
    replace github.com/2637309949/bulrush-addition => ../bulrush-addition
    replace github.com/2637309949/bulrush-limit => ../bulrush-limit
    ...
    // ## end


### 3. For Apidoc
![Bulrush flash](./assets/apidoc.png)

```shell
$ make -f Makefile.api
```

### 4. Run Test

```shell
/usr/local/go/bin/go test -timeout 30s github.com/2637309949/bulrush-template -run "^(TestCache)$"
```
Or run with log
```shell
/usr/local/go/bin/go test -timeout 30s github.com/2637309949/bulrush-template -run "^(TestCache)$" -v
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