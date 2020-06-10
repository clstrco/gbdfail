package gbdfail

import (
	"net/http"

	"github.com/markbates/pkger"
)

func Routes(mux *http.ServeMux) {
	mux.Handle("/s", http.FileServer(pkger.Dir("/s")))
}
