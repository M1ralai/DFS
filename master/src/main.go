// Package main DFS Master entrypoint.
//
// @title           DFS Master API
// @version         1.0
// @description     Distributed File System master service: node orchestration ve client metadata yönetimi.
// @host            localhost:3030
// @BasePath        /
package main

import (
	"github.com/M1ralai/DFS/src/internal/server"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	srv := server.NewServer()
	srv.Run()
}
