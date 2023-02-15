package main

import (
	"log"

	"github.com/TomChv/csc-0847/project_1/backend/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
