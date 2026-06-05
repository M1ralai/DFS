// Package main DFS Node entrypoint.
package main

import (
	"github.com/M1ralai/DFS/node/src/internal/server"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	srv := server.NewServer()
	srv.Run()
}
