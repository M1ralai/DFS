package server

import (
	"github.com/M1ralai/DFS/master/utils/config"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer() *Server {
	return &Server{
		cfg: config.LoadConfig(),
		app: fiber.New(fiber.Config{
			ServerHeader:  "Master",
			CaseSensitive: true,
		}),
	}
}

func (s *Server) Run() {

}
