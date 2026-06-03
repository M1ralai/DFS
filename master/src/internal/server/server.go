package server

import (
	"context"
	"log"

	"github.com/M1ralai/DFS/src/internal/infrastructure/db"
	nodeHandler "github.com/M1ralai/DFS/src/internal/module/node/handler"
	nodeRespository "github.com/M1ralai/DFS/src/internal/module/node/repository"
	nodeService "github.com/M1ralai/DFS/src/internal/module/node/service"
	"github.com/M1ralai/DFS/src/utils/config"
	"github.com/M1ralai/DFS/src/utils/validator"
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

	nodeRepo := nodeRespository.NewRepo(DB)
	nodeSrv := nodeService.NewNodeCommService(nodeRepo, s.cfg.NodeCommCfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodeSrv.StartDeadNodeChecker(ctx)

	nodeHnd := nodeHandler.NewNodeCommHandler(nodeSrv)

	nodeHnd.RegisterRoute(s.app)

	s.app.Listen(s.cfg.Host + s.cfg.Port)

}
