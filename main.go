package main

import (
	"os"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return GetMessage()
	})
	m.Run()
}

func GetMessage() string {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello world"
	}

	return message
}
