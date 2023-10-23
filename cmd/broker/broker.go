package main

import (
	"log"

	"github.com/threadedstream/quicthing/internal/cmd/broker"
	_ "github.com/threadedstream/quicthing/internal/config"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	broker.Main()
}
