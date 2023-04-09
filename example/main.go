package main

import (
	"net/http"

	"github.com/dundunlabs/httplog"
)

func main() {
	handler := httplog.NewHandler(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
