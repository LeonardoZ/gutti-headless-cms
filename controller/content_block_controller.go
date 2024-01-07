package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonardoZ/gutti-headless-cms/dto"
	"github.com/LeonardoZ/gutti-headless-cms/error_handler"
	"github.com/LeonardoZ/gutti-headless-cms/repository"
	"github.com/LeonardoZ/gutti-headless-cms/service"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"gorm.io/gorm"
)

var (
	g          = galidator.New()
	customizer = g.Validator(dto.UpsertContentBlockDTO{})
)

type ContentBlockController struct {
	Service service.ContentBlockService
}

func NewContentBlockController(service service.ContentBlockService) *ContentBlockController {
	return &ContentBlockController{
		Service: service,
	}
}

func RegisterContentBlockRoutes(db *gorm.DB, router *gin.Engine) {
	repository := repository.NewContentBlockRepository(db)
	service := service.NewContentBlockService(repository)
	controller := NewContentBlockController(service)

	contentRouter := router.Group("/content-blocks")
	contentRouter.GET("", controller.FindAll)
	contentRouter.GET("/:id", controller.FindById)
	contentRouter.POST("", controller.Create)
	contentRouter.PUT("/:id", controller.Update)
	contentRouter.DELETE("/:id", controller.Delete)

}

func (c *ContentBlockController) Create(ctx *gin.Context) {
	createContentDto := dto.UpsertContentBlockDTO{}
	err := ctx.ShouldBindJSON(&createContentDto)
	if err != nil {
		httpErr := error_handler.NewHttpError("Invalid Body", customizer.DecryptErrors(err), http.StatusUnprocessableEntity)
		ctx.Error(httpErr)
		return
	}
	err = c.Service.Create(createContentDto)

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Writer.WriteHeader(201)
}

func (c *ContentBlockController) Update(ctx *gin.Context) {
	updateDto := dto.UpsertContentBlockDTO{}
	err := ctx.ShouldBindJSON(&updateDto)
	if err != nil {
		httpErr := error_handler.NewHttpError("Invalid Body", fmt.Sprintf("%v", err), http.StatusUnprocessableEntity)
		ctx.Error(httpErr)
		return
	}

	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		httpErr := error_handler.NewHttpError("Invalid Param Id", "", http.StatusBadRequest)
		ctx.Error(httpErr)
		return
	}
	err = c.Service.Update(id, updateDto)

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Writer.WriteHeader(200)
}

func (c *ContentBlockController) Delete(ctx *gin.Context) {
	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		httpErr := error_handler.NewHttpError("Invalid Param Id", "", http.StatusBadRequest)
		ctx.Error(httpErr)
		return
	}
	err = c.Service.Delete(id)

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Writer.WriteHeader(200)
}

func (c *ContentBlockController) FindById(ctx *gin.Context) {
	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		ctx.Error(err)
		return
	}
	content, err := c.Service.FindById(id)
	if err != nil {
		httpErr := error_handler.NewHttpError("Failed to fetch content by id", fmt.Sprintf("%v", err), http.StatusInternalServerError)
		ctx.Error(httpErr)
		return
	}
	if content == nil {
		ctx.Error(err)
		return
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, content)
}

func (c *ContentBlockController) FindAll(ctx *gin.Context) {
	contents, err := c.Service.FindAll()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, contents)
}
