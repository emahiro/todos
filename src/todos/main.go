package main

import (
	"fmt"
	"net/http"

	"todos/handler"
)

var addr =  ":3000"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.Index)
	fmt.Printf("[START] server. port: %s\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}