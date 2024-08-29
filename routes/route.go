package routes

import (
	maw_delivery "ui-rice-go/app/master-author-weight/delivery"
	maw_repo "ui-rice-go/app/master-author-weight/repository"
	maw_usecase "ui-rice-go/app/master-author-weight/usecase"
	"ui-rice-go/configs"

	"github.com/gofiber/fiber/v2"
)

func RouteRegister(app *fiber.App, config configs.Config) {
	// api versioning, check config.yaml file to set api versioning
	ver := app.Group("/api/" + config.ApiVersion)

	// call each modules to register to fiber routing
	// TODO, init and assign repo interface
	mawr := maw_repo.NewMasterAuthorWeightRepository(configs.DBConn)
	mawu := maw_usecase.NewMasterAuthorWeightUsecase(mawr)
	maw_delivery.NewMasterAuthorWeightHandler(ver, mawu)
}
