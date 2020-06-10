package gbdfail

import "net/http"

func Routes(mux *http.ServeMux) {
	// just like the example at
	// https://github.com/go-bindata/go-bindata#serve-assets-with-nethttp we
	// leave the trailing slash off:
	mux.Handle("/s", http.FileServer(AssetFile()))
}
