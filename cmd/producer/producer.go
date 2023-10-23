package main

import (
	"log"

	"github.com/threadedstream/quicthing/internal/cmd/producer"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	producer.Main()
}
