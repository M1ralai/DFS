package server

import (
	"context"
	"log"

	"github.com/M1ralai/DFS/src/internal/infrastructure/db"
	nodeHandler "github.com/M1ralai/DFS/src/internal/module/node/handler"
	nodeRepository "github.com/M1ralai/DFS/src/internal/module/node/repository"
	nodeService "github.com/M1ralai/DFS/src/internal/module/node/service"

	clientHandler "github.com/M1ralai/DFS/src/internal/module/client/handler"
	clientRepository "github.com/M1ralai/DFS/src/internal/module/client/repository"
	clientService "github.com/M1ralai/DFS/src/internal/module/client/service"

	_ "github.com/M1ralai/DFS/src/docs"

	"github.com/M1ralai/DFS/src/utils/config"
	"github.com/M1ralai/DFS/src/utils/validator"
	"github.com/gofiber/contrib/v3/swaggo"
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

	clientRepo := clientRepository.NewClientRepository(DB)
	nodeRepo := nodeRepository.NewNodeRepository(DB)
	nodeSrv := nodeService.NewNodeService(nodeRepo, clientRepo, s.cfg.NodeCfg)
	clientSrv := clientService.NewClientService(clientRepo, nodeRepo, s.cfg.NodeCfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodeSrv.StartDeadNodeChecker(ctx)

	nodeHnd := nodeHandler.NewNodeHandler(nodeSrv)
	clientHnd := clientHandler.NewClientHandler(clientSrv)

	s.app.Get("/swagger/*", swaggo.New())

	nodeHnd.RegisterRoute(s.app)
	clientHnd.RegisterRoute(s.app)

	s.app.Listen(s.cfg.Host + s.cfg.Port)

}
