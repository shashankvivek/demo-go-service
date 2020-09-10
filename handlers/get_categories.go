package handlers

import (
	"demo-svc/models"
	"demo-svc/restapi/operations/menu"
	"github.com/go-openapi/runtime/middleware"
)

type menuImpl struct {
}

func NewMenuCategoryHandler() menu.CategoryListHandler {
	return &menuImpl{}
}

func (impl *menuImpl) Handle(param menu.CategoryListParams) middleware.Responder {

	retVal := models.Categories{}
	retVal = append(retVal, &models.Category{
		BcID:          1,
		BcIsActive:    true,
		BcImageURL:    "image.png",
		BcName:        "Category1",
		SubCategories: nil,
	})

	return menu.NewCategoryListOK().WithPayload(retVal)
}
