package handler

import (
	"github.com/M1ralai/DFS/src/internal/module/client/dto"
	"github.com/M1ralai/DFS/src/internal/module/client/service"
	"github.com/M1ralai/DFS/src/utils/response"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ClientHandler struct {
	service service.IClientService
}

func NewClientHandler(service service.IClientService) *ClientHandler {
	return &ClientHandler{
		service: service,
	}
}

func (h *ClientHandler) RegisterRoute(app *fiber.App) {
	app.Post("/api/upload", h.Upload)
	app.Get("/api/file/:id", h.GetFile)
	app.Delete("/api/file", h.DeleteFile)
	app.Get("/api/user/:user_id/files", h.GetUserFiles)
}

// Upload godoc
// @Summary      Dosya yükleme planı oluştur
// @Description  Bir dosya için chunk placement üretir. Hangi chunk'ın hangi node'lara yazılacağını döner.
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        request body dto.UploadRequest true "Yüklenecek dosya bilgileri"
// @Success      201 {object} response.Response[dto.UploadResponse]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/upload [post]
func (h *ClientHandler) Upload(ctx fiber.Ctx) error {
	d := new(dto.UploadRequest)
	if err := ctx.Bind().Body(d); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	v, err := h.service.UploadFile(*d)
	if err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(201).JSON(response.NewResponse(true, v, "dosya organize edildi"))
	return nil
}

// GetFile godoc
// @Summary      Dosya metadata getir
// @Description  Verilen file_id için dosya bilgisi ve chunk placement'ı döner.
// @Tags         client
// @Produce      json
// @Param        id path string true "File UUID"
// @Success      200 {object} response.Response[dto.FileResponse]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/file/{id} [get]
func (h *ClientHandler) GetFile(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, "geçersiz file id"))
		return err
	}
	v, err := h.service.GetFile(id)
	if err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, v, "dosya başarıyla getirildi"))
	return nil
}

// DeleteFile godoc
// @Summary      Dosyayı sil
// @Description  Verilen file_id'ye sahip dosyayı ve tüm chunk'larını siler.
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        request body dto.DeleteRequest true "Silme isteği"
// @Success      200 {object} response.Response[int]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/file [delete]
func (h *ClientHandler) DeleteFile(ctx fiber.Ctx) error {
	d := new(dto.DeleteRequest)
	if err := ctx.Bind().Body(d); err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	if err := h.service.DeleteFile(*d); err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, 1, "dosya silindi"))
	return nil
}

// GetUserFiles godoc
// @Summary      Kullanıcının dosyalarını listele
// @Description  Verilen user_id'ye ait tüm dosyaları ve chunk placement'larını döner.
// @Tags         client
// @Produce      json
// @Param        user_id path string true "User UUID"
// @Success      200 {object} response.Response[dto.UserFilesResponse]
// @Failure      400 {object} response.Response[any]
// @Failure      500 {object} response.Response[any]
// @Router       /api/user/{user_id}/files [get]
func (h *ClientHandler) GetUserFiles(ctx fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Params("user_id"))
	if err != nil {
		ctx.Status(400).JSON(response.NewResponse[any](false, nil, "geçersiz user id"))
		return err
	}
	v, err := h.service.GetUserFiles(userID)
	if err != nil {
		ctx.Status(500).JSON(response.NewResponse[any](false, nil, err.Error()))
		return err
	}
	ctx.Status(200).JSON(response.NewResponse(true, v, "kullanıcı dosyaları getirildi"))
	return nil
}
