package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"

	"ui-rice-go/domain"
)

type masterAuthorWeightRepository struct {
	DB *gorm.DB
}

func NewMasterAuthorWeightRepository(DB *gorm.DB) domain.MasterAuthorWeightRepository {
	return &masterAuthorWeightRepository{DB}
}

func (mawr *masterAuthorWeightRepository) FetchAll(c *fiber.Ctx) (paged paginate.Page, rowLen int64, err error) {
	var mawdomain []domain.MasterAuthorWeight
	pg := paginate.New()
	result := mawr.DB.Where("is_deleted = ?", "0").Find(&mawdomain)
	pager := pg.With(result).Request(c.Request()).Response(&mawdomain)
	return pager, result.RowsAffected, result.Error
}

func (mawr *masterAuthorWeightRepository) FetchOneByID(c *fiber.Ctx, id uint) (res domain.MasterAuthorWeight, err error) {
	result := mawr.DB.Where("is_deleted = ?", "0").First(&res, id)
	return res, result.Error
}

func (mawr *masterAuthorWeightRepository) Save(c *fiber.Ctx) (res domain.MasterAuthorWeight, err error) {
	maw := new(domain.MasterAuthorWeight)
	if err1 := c.BodyParser(maw); err1 != nil {
		return
	}
	mawr.DB.Create(&maw)
	return *maw, nil
}

func (mawr *masterAuthorWeightRepository) Edit(c *fiber.Ctx, id uint) (res domain.MasterAuthorWeight, err error) {
	var maw domain.MasterAuthorWeight
	result := mawr.DB.Where("is_deleted = ?", "0").First(&maw, id)
	if result.Error != nil {
		return domain.MasterAuthorWeight{}, result.Error
	}
	if err1 := c.BodyParser(&maw); err1 != nil {
		return domain.MasterAuthorWeight{}, err1
	}
	mawr.DB.Save(&maw)
	return maw, nil
}

func (mawr *masterAuthorWeightRepository) SoftDelete(c *fiber.Ctx, id uint, dltdby string) (err error) {
	var maw domain.MasterAuthorWeight
	result := mawr.DB.Where("is_deleted = ? AND id = ?", "0", id).First(&maw)
	if result.Error != nil {
		return result.Error
	}
	maw.IsDeleted = "1"
	maw.UpdatedBy = dltdby
	mawr.DB.Save(&maw)
	return
}
