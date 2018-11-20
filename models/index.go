package models
import (
	"github.com/2637309949/bulrush"
)


// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	mongo, ok := injects["Mongo"].(*bulrush.MongoGroup)
	if ok {
		user(mongo)
	}
}