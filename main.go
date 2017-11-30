package main

import (
	"log"

	"github.com/ulule/dekiteru/cmd"
)

func main() {
	cmd := cmd.New()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
