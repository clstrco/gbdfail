package gbdfail

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestStaticAssets(t *testing.T) {
	mux := http.NewServeMux()
	Routes(mux)
	ts := httptest.NewServer(mux)

	tests := []struct {
		label  string
		path   string
		status int
	}{
		{
			label:  "garbage",
			path:   "/laksjdf/99.txt",
			status: http.StatusNotFound,
		},
		{
			label: "top level",
			path:  "/s/00.txt",
		},
		{
			label: "nested",
			path:  "/s/a/01.txt",
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			if test.status == 0 {
				test.status = http.StatusOK
			}

			resp, err := http.Get(ts.URL + test.path)
			if err != nil {
				t.Errorf("http get: %v", err)
				return
			}

			if got, want := resp.StatusCode, test.status; got != want {
				t.Errorf("status code; got %v, want %v", got, want)
				return
			}

			if test.status != http.StatusOK {
				return
			}

			_, fn := filepath.Split(test.path)
			contents := fn[:len(fn)-len(filepath.Ext(fn))] + "\n"
			buf := &bytes.Buffer{}
			if _, err := io.Copy(buf, resp.Body); err != nil {
				t.Errorf("copy body: %v", err)
			}

			if got, want := buf.String(), contents; got != want {
				t.Errorf("file contents; got %q, want %q", got, want)
			}
		})
	}
}
