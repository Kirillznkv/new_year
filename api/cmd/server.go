package main

import (
	"github.com/Kirillznkv/new_year/api/internal/front"
	"github.com/Kirillznkv/new_year/api/internal/server"
	"log"
)

func StartApi(ch chan<- error) {
	if err := server.Start(); err != nil {
		ch <- err
	}
}

func StartFront(ch chan<- error) {
	if err := front.Start(); err != nil {
		ch <- err
	}
}

func main() {
	ch := make(chan error)

	go StartApi(ch)
	go StartFront(ch)

	for {
		select {
		case err := <-ch:
			log.Println(err)
		}
	}
}
