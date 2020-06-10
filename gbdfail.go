package gbdfail

import "net/http"

func Routes(mux *http.ServeMux) {
	mux.Handle("/s/", http.StripPrefix("/s/", http.FileServer(AssetFile())))
}
