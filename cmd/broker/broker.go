package main

import (
	"github.com/threadedstream/quicthing/internal/cmd/broker"
	_ "github.com/threadedstream/quicthing/internal/config"
)

func main() {
	broker.Main()
}
