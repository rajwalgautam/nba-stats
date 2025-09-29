package main

import (
	"log"

	"github.com/rajwalgautam/nba-stats/internal/pkg/gamestatsjob"
)

func main() {
	err := gamestatsjob.Run()
	if err != nil {
		log.Fatalf("error running gamestatsjob: %v", err)
	}
}
