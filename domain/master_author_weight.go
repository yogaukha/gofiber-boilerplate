package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

type MasterAuthorWeight struct {
	// do not use gorm.Model, but type it manually
	ID          uint      `json:"id"`
	CreatedAt   time.Time `gorm:"column:created_time;<-:create" json:"created_time"`
	UpdatedAt   time.Time `gorm:"column:updated_time;<-:update" json:"updated_time"`
	Description string    `json:"description"`
	Value       float32   `json:"value"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `gorm:"default:null" json:"updated_by"`
	IsDeleted   string    `gorm:"default:'0'" json:"is_deleted"`
}

type MasterAuthorWeightUsecase interface {
	FetchAll(c *fiber.Ctx, page int, size int) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx) (MasterAuthorWeight, error)
	Save(c *fiber.Ctx) (MasterAuthorWeight, error)
	Edit(c *fiber.Ctx) (MasterAuthorWeight, error)
	SoftDelete(c *fiber.Ctx) error
}

type MasterAuthorWeightRepository interface {
	FetchAll(c *fiber.Ctx) (paginate.Page, int64, error)
	FetchOneByID(c *fiber.Ctx, id uint) (MasterAuthorWeight, error)
	Save(c *fiber.Ctx) (MasterAuthorWeight, error)
	Edit(c *fiber.Ctx, id uint) (MasterAuthorWeight, error)
	SoftDelete(c *fiber.Ctx, id uint, dltdby string) error
}
