package server

import (
	"log"

	"github.com/M1ralai/DFS/node/src/internal/infrastructure/db"
	"github.com/M1ralai/DFS/node/src/utils/config"
	"github.com/M1ralai/DFS/node/src/utils/response"
	"github.com/M1ralai/DFS/node/src/utils/validator"
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
			ServerHeader:    "Node",
			CaseSensitive:   true,
			StructValidator: validator.NewValidator(),
		}),
	}
}

func (s *Server) Run() {
	DB, err := db.NewPostgresDB(s.cfg.DBCfg)
	if err != nil {
		log.Fatal("db connection cannot established", err)
	}

	if err := db.Migrate(DB, true); err != nil {
		log.Fatal("db migration is failed.", err)
	}

	s.app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(200).JSON(response.NewResponse(true, "node is alive", ""))
	})

	s.app.Listen(s.cfg.Host + s.cfg.Port)
}
