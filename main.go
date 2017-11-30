package main

import (
	"log"

	"github.com/ulule/dekiteru/cmd"
)

func main() {
	err := cmd.New().Run()
	if err != nil {
		log.Fatal(err)
	}
}
