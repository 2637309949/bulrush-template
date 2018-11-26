package middles

import (
	"path"
	"github.com/2637309949/bulrush/middles"
)

func static (injects map[string]interface{}) {
	serive := middles.Serve {
		URLPrefix: "/public",
		Fs: middles.LocalFile(path.Join("assets/public", ""), false),
	}

	serive.Inject(injects)
}