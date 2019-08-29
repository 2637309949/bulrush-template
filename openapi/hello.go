// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package openapi

import (
	openapi "github.com/2637309949/bulrush-openapi"
)

// RegisterHello for openapi
func RegisterHello(api *openapi.OpenAPI) {
	api.RegistHandler(openapi.Handler{
		Name:    "test.hello",
		Version: "1.0",
		Voke: func(*openapi.AppInfo, *openapi.Form) (*openapi.FormRet, error) {
			return &openapi.FormRet{
				Body: map[string]interface{}{
					"a": "a",
					"b": "b",
				},
			}, nil
		},
	})
}
