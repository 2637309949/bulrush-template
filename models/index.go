package models
import (
	"github.com/2637309949/bulrush"
)


// Inject as Function that accept some injection parameters
func Inject (injects map[string]interface{}) {
	mongo, _ := injects["Mongo"].(*bulrush.MongoGroup)
	user(mongo)
}