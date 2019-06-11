/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package plugins

import (
	openapi "github.com/2637309949/bulrush-openapi"
)

// OpenAPI Plugin init
var OpenAPI = &openapi.OpenAPI{
	URLPrefix: "/gateway",
	Auth: func(appid string) (*openapi.AppInfo, error) {
		return &openapi.AppInfo{
			AppID: "xxx",
			// PKCS1 Public Key
			PublicKey: `
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDaYoospKXldKm+/fdGh8b5Jbq4
PSrIELjHR/F6+zYExr7w0XGWC6/1Nj2Da3EWYnWo+n6URMqeyCnd09SvEmavpkUW
d/4AhjZIvO42PBy9x/wFrqgU/KUqZ2r+AWFJJZgSA9IuBEqjZxsu1XTKiT7v5U7N
WlboP5QLPPffz92sUwIDAQAB
			`,
			// JUST FOR TEST
			PrivateKey: `
MIICXAIBAAKBgQDaYoospKXldKm+/fdGh8b5Jbq4PSrIELjHR/F6+zYExr7w0XGW
C6/1Nj2Da3EWYnWo+n6URMqeyCnd09SvEmavpkUWd/4AhjZIvO42PBy9x/wFrqgU
/KUqZ2r+AWFJJZgSA9IuBEqjZxsu1XTKiT7v5U7NWlboP5QLPPffz92sUwIDAQAB
AoGAAhROHH601ap2s0rXv+QrENQ7IuXyMlV2bO9SbUlXClSaHNDhs/wIgN0zWLz9
JqlpVWKNMfw1sa1WOLZ7n+8c6yq5cJLKjlHwilm0JKRUcrKY1JEln8VslIhkhBQ5
XADsNY4bevfyn4obm0/aqgXzJjLxgdMr5frlkmYaHjJhcAECQQDtC6lV4qbvphT/
JtiLiA+PS7UYpMNJLwopLofJkWwGtx+SCcEAmwEGd/DPs8jDw0xp8T6kxzs5ewHd
kp4xnWyBAkEA69jk1VX4kHd3vmWVSmPZR6bPxlhGmrOunPxErwRIffiLwCCZjoUt
kJPj/Lz0ZC5zCJ0+BUO3XUGH3Yrgl9Y+0wJAAiQo023ItEF7zxI6wofoHBNC/4X9
fZu/K8AP2fJGV4kv79HGvIqp57UNp9Kn7ZzGA+758eMa1doWmjnI1AnoAQJAX5wl
6HZZtfc4i+8Sfn3L78goIvhWZxDAkNOT5H+QA8FmphSRK73BowLRQfw39wT9jVhx
dCDnjN6r/Zv3QJaSEQJBAK1hJlgv1pI0g2GQy/yrcb7k8QWWaKl0DgsKDVlFSjAx
9agJ0rfrBVFsjyJKlC1cpMdC/ScoqBLTIscE/12yjBg=
			`,
		}, nil
	},
}
