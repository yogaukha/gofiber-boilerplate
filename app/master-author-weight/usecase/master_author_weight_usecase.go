package usecase

import (
	"strconv"

	"ui-rice-go/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type masterAuthorWeightUsecase struct {
	masterAuthorWeightRepo domain.MasterAuthorWeightRepository
}

func NewMasterAuthorWeightUsecase(mar domain.MasterAuthorWeightRepository) domain.MasterAuthorWeightUsecase {
	return &masterAuthorWeightUsecase{
		masterAuthorWeightRepo: mar,
	}
}

func (mawu *masterAuthorWeightUsecase) FetchAll(c *fiber.Ctx, page int, size int) (paged paginate.Page, len int64, err error) {
	paged, len, err = mawu.masterAuthorWeightRepo.FetchAll(c)
	if err != nil {
		return paginate.Page{}, 0, err
	}

	return
}

func (mawu *masterAuthorWeightUsecase) FetchOneByID(c *fiber.Ctx) (res domain.MasterAuthorWeight, err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	res, err = mawu.masterAuthorWeightRepo.FetchOneByID(c, uint(id))

	return
}

func (mawu *masterAuthorWeightUsecase) Save(c *fiber.Ctx) (res domain.MasterAuthorWeight, err error) {
	res, err = mawu.masterAuthorWeightRepo.Save(c)
	return
}

func (mawu *masterAuthorWeightUsecase) Edit(c *fiber.Ctx) (res domain.MasterAuthorWeight, err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	res, err = mawu.masterAuthorWeightRepo.Edit(c, uint(id))
	return
}

func (mawu *masterAuthorWeightUsecase) SoftDelete(c *fiber.Ctx) (err error) {
	id, err1 := strconv.Atoi(c.Params("id"))
	if err1 != nil {
		return
	}
	err = mawu.masterAuthorWeightRepo.SoftDelete(c, uint(id), c.Params("username"))
	return
}
