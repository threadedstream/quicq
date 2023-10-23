package main

import (
	"log"

	"github.com/threadedstream/quicthing/internal/cmd/consumer"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	consumer.Main()
}
