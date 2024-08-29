package delivery

import (
	"strconv"
	"strings"

	"ui-rice-go/domain"
	"ui-rice-go/internal"

	"github.com/gofiber/fiber/v2"
)

type MasterAuthorWeightHandler struct {
	MawUsecase domain.MasterAuthorWeightUsecase
}

func NewMasterAuthorWeightHandler(c fiber.Router, uc domain.MasterAuthorWeightUsecase) {
	handler := &MasterAuthorWeightHandler{
		MawUsecase: uc,
	}
	c.Get("/hello", handler.hello)
	c.Get("/master-author-weights", handler.FetchAll)
	c.Get("/master-author-weight/:id", handler.FetchOneByID)
	c.Post("/master-author-weight", handler.Save)
	c.Patch("/master-author-weight/:id", handler.Edit)
	c.Delete("/master-author-weight/:id/:username", handler.SoftDelete)
}

func (mawh *MasterAuthorWeightHandler) hello(c *fiber.Ctx) error {
	return internal.ReturnTheResponse(c, false, int(200), "", "hellosss")
}

func (mawh *MasterAuthorWeightHandler) FetchAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return internal.ReturnTheResponse(c, false, int(400), "Bad Request", nil)
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		return internal.ReturnTheResponse(c, false, int(400), "Bad Request", nil)
	}
	listMaw, rowLen, err := mawh.MawUsecase.FetchAll(c, page, size)
	if rowLen <= 0 {
		return internal.ReturnTheResponse(c, true, int(404), "Record not Found", nil)
	}
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}

	return internal.ReturnTheResponse(c, false, int(200), "", listMaw)
}

func (mawh *MasterAuthorWeightHandler) FetchOneByID(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.FetchOneByID(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
	}

	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *MasterAuthorWeightHandler) Save(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.Save(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *MasterAuthorWeightHandler) Edit(c *fiber.Ctx) error {
	res, err := mawh.MawUsecase.Edit(c)
	if err != nil {
		errMessage := err.Error()
		if strings.Contains(errMessage, "not found") {
			return internal.ReturnTheResponse(c, true, int(404), err.Error(), nil)
		} else {
			return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
		}
	}
	return internal.ReturnTheResponse(c, false, int(200), "", res)
}

func (mawh *MasterAuthorWeightHandler) SoftDelete(c *fiber.Ctx) error {
	err := mawh.MawUsecase.SoftDelete(c)
	if err != nil {
		return internal.ReturnTheResponse(c, true, int(500), err.Error(), nil)
	}
	return internal.ReturnTheResponse(c, false, int(200), "Deleted succesfully", nil)
}
