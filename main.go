package main

import (
	"log"
	"net/http"
)

func main() {

	ps := &PlayerServer{store: NewInMmemoryPlayerStore()}

	handler := http.HandlerFunc(ps.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
