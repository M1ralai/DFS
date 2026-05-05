package server

import (
	"context"
	"log"

	"github.com/M1ralai/DFS/master/internal/infrastructure/db"
	"github.com/M1ralai/DFS/master/internal/module/node/handler"
	"github.com/M1ralai/DFS/master/internal/module/node/repository"
	"github.com/M1ralai/DFS/master/internal/module/node/service"
	"github.com/M1ralai/DFS/master/utils/config"
	"github.com/M1ralai/DFS/master/utils/validator"
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
			ServerHeader:    "Master",
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

	repo := repository.NewRepo(DB)
	service := service.NewNodeCommService(repo, s.cfg.NodeCommCfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	service.StartDeadNodeChecker(ctx)

	handler := handler.NewNodeCommHandler(service)

	handler.RegisterRoute(s.app)

	s.app.Listen(s.cfg.Host + s.cfg.Port)

}
