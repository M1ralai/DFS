package handler

import (
	"github.com/M1ralai/DFS/src/internal/module/node/dto"
	"github.com/M1ralai/DFS/src/internal/module/node/service"
	"github.com/M1ralai/DFS/src/utils/response"
	"github.com/gofiber/fiber/v3"
)

type NodeHandler struct {
	service service.INodeService
}

func NewNodeHandler(service service.INodeService) *NodeHandler {
	return &NodeHandler{
		service: service,
	}
}

func (h *NodeHandler) RegisterRoute(app *fiber.App) {
	app.Post("/api/node", h.Save)
	app.Get("/api/node", h.FindAll)
	app.Post("/api/heartbeat", h.Heartbeat)
	app.Post("/api/acknowledgement", h.Acknowledgement)
}

// Save godoc
// @Summary      Yeni node kaydet
// @Description  Bir storage node'unu master'a kaydeder.
// @Tags         node
// @Accept       json
// @Produce      json
// @Param        request body dto.NodeSaveRequest true "Node kayıt bilgileri"
// @Success      201 {object} response.Response[string]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/node [post]
func (h *NodeHandler) Save(ctx fiber.Ctx) error {
	d := new(dto.NodeSaveRequest)
	if err := ctx.Bind().Body(d); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.Save(*d); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(201).JSON("node succesfully registered")
	return nil
}

// FindAll godoc
// @Summary      Tüm node'ları listele
// @Description  Master'a kayıtlı tüm node'ları döner.
// @Tags         node
// @Produce      json
// @Success      200 {object} response.Response[[]dto.NodeSaveRequest]
// @Failure      500 {object} response.Response[any]
// @Router       /api/node [get]
func (h *NodeHandler) FindAll(ctx fiber.Ctx) error {
	v, err := h.service.FindAll()
	if err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, v, "tüm nodelar başarıyla getirildi"))
	return nil

}

// Heartbeat godoc
// @Summary      Node heartbeat al
// @Description  Node'dan periyodik heartbeat alır, son görülme zamanını ve available space'i günceller.
// @Tags         node
// @Accept       json
// @Produce      json
// @Param        request body dto.HeartbeatRequest true "Heartbeat bilgileri"
// @Success      200 {object} response.Response[int]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/heartbeat [post]
func (h *NodeHandler) Heartbeat(ctx fiber.Ctx) error {
	d := new(dto.HeartbeatRequest)
	if err := ctx.Bind().Body(d); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.Heartbeat(*d); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, 1, "heartbeat basaiyla onaylandı"))
	return nil
}

// Acknowledgement godoc
// @Summary      Chunk ack al
// @Description  Node bir chunk'ı başarıyla kaydettiğinde master'a bildirir. Replica sayacı artırılır.
// @Tags         node
// @Accept       json
// @Produce      json
// @Param        request body dto.AckRequest true "Ack bilgileri"
// @Success      200 {object} response.Response[int]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/acknowledgement [post]
func (h *NodeHandler) Acknowledgement(ctx fiber.Ctx) error {
	d := new(dto.AckRequest)
	if err := ctx.Bind().Body(d); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.Acknowledgement(*d); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, 1, "acknowledgement basaiyla onaylandı"))
	return nil
}
