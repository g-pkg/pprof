package pprof

import (
	"net/http"
	_ "net/http/pprof"
)

//Serve serve pprof
func Serve(addr string) error {
	return http.ListenAndServe(addr, nil)
}
