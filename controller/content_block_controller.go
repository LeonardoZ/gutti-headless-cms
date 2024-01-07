package controller

import (
	"net/http"
	"strconv"

	"github.com/LeonardoZ/gutti-headless-cms/dto"
	"github.com/LeonardoZ/gutti-headless-cms/repository"
	"github.com/LeonardoZ/gutti-headless-cms/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	validate := validator.New()
	service := service.NewContentBlockService(repository, validate)
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
		panic(err)
	}
	c.Service.Create(createContentDto)
	ctx.Writer.WriteHeader(201)
}

func (c *ContentBlockController) Update(ctx *gin.Context) {
	updateDto := dto.UpsertContentBlockDTO{}
	err := ctx.ShouldBindJSON(&updateDto)
	if err != nil {
		panic(err)
	}

	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		panic(err)
	}
	c.Service.Update(id, updateDto)
	ctx.Writer.WriteHeader(200)
}

func (c *ContentBlockController) Delete(ctx *gin.Context) {
	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		panic(err)
	}
	c.Service.Delete(id)
	ctx.Writer.WriteHeader(200)
}

func (c *ContentBlockController) FindById(ctx *gin.Context) {
	contentId := ctx.Param("id")
	id, err := strconv.Atoi(contentId)
	if err != nil {
		panic(err)
	}
	content := c.Service.FindById(id)
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, content)
}

func (c *ContentBlockController) FindAll(ctx *gin.Context) {
	contents := c.Service.FindAll()
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, contents)
}
