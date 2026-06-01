package handler

import (
	"github.com/M1ralai/DFS/master/internal/module/node/dto"
	"github.com/M1ralai/DFS/master/internal/module/node/service"
	"github.com/M1ralai/DFS/master/utils/response"
	"github.com/gofiber/fiber/v3"
)

type NodeCommHandler struct {
	service service.INodeCommService
}

func NewNodeCommHandler(service service.INodeCommService) NodeCommHandler {
	return NodeCommHandler{
		service: service,
	}
}

func (h *NodeCommHandler) RegisterRoute(app *fiber.App) {
	app.Post("/api/node", h.Save)
	app.Get("/api/node", h.FindAll)
	app.Post("/api/hearthbeat", h.HearthBeat)
	app.Post("/api/acknowledgement", h.Acknowledgement)
}

func (h *NodeCommHandler) Save(ctx fiber.Ctx) error {
	dto := new(dto.NodeSaveRequest)
	if err := ctx.Bind().Body(dto); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.Save(*dto); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(201).JSON("node succesfully registered")
	return nil
}

func (h *NodeCommHandler) FindAll(ctx fiber.Ctx) error {
	v, err := h.service.FindAll()
	if err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, v, "tüm nodelar başarıyla getirildi"))
	return nil

}

func (h *NodeCommHandler) HearthBeat(ctx fiber.Ctx) error {
	dto := new(dto.HearthBeatRequest)
	if err := ctx.Bind().Body(dto); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.HearthBeat(*dto); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, 1, "hearthbeat basaiyla onaylandı"))
	return nil
}

func (h *NodeCommHandler) Acknowledgement(ctx fiber.Ctx) error {
	dto := new(dto.AckRequest)
	if err := ctx.Bind().Body(dto); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.Acknowledgement(*dto); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, 1, "acknowledgement basaiyla onaylandı"))
	return nil
}
