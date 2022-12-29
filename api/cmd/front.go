package main

import (
	"github.com/Kirillznkv/new_year/api/internal/front"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":80", front.NewServer()); err != nil {
		log.Fatal(err)
	}
}
