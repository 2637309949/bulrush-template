/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:49:48
 * @modify date 2019-01-16 20:49:48
 * @desc [description]
 */

package openapi

import (
	openapi "github.com/2637309949/bulrush-openapi"
)

// RegisterHello for openapi
func RegisterHello(api *openapi.OpenAPI) {
	api.RegistHandler(openapi.Handler{
		Name:    "test.hello",
		Version: "1.0",
		Voke: func(*openapi.AppInfo, *openapi.CRP) (*openapi.CRPRet, error) {
			return &openapi.CRPRet{
				Body: map[string]interface{}{
					"a": "a",
					"b": "b",
				},
			}, nil
		},
	})
}
