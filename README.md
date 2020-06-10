# demonstrates a bug

When I try to use the [go-bindata instructions on serving assets using `net/http`](https://github.com/go-bindata/go-bindata#serve-assets-with-nethttp) I only get `404 page not found` responses.

This repository reproduces the bug and hopefully will help folks fix `go-bindata`

## demonstrate bug:

Just run `make test`:

```bash
$ make test                                                                                                               
go mod vendor
go build -o build/go-bindata ./vendor/github.com/go-bindata/go-bindata/go-bindata
./build/go-bindata -o static.go -fs -pkg gbdfail -prefix "s/" s/...
go test
--- FAIL: TestStaticAssets (0.00s)
    --- FAIL: TestStaticAssets/top_level (0.00s)
        gbdfail_test.go:50: status code; got 404, want 200
    --- FAIL: TestStaticAssets/nested (0.00s)
        gbdfail_test.go:50: status code; got 404, want 200
FAIL
exit status 1
FAIL    clstr.co/gbdfail        0.007s
Makefile:22: recipe for target 'test' failed
make: *** [test] Error 1
```
