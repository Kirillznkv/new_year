package main

import (
	"github.com/Kirillznkv/new_year/api/internal/server"
	"log"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
