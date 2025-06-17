package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.ListenAndServe(":8000", nil)
}
