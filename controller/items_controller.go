package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"meli-code-challenge/service"
	"net/http"
)

type ItemsController struct {
	tagsService service.ItemsService
}

func NewItemsController(service service.ItemsService) *ItemsController {
	return &ItemsController{
		tagsService: service,
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
//func (controller *ItemsController) Create(ctx *gin.Context) {
//	log.Info().Msg("create tags")
//	createTagsRequest := request.CreateTagsRequest{}
//	err := ctx.ShouldBindJSON(&createTagsRequest)
//	helper.ErrorPanic(err)
//
//	controller.tagsService.Create(createTagsRequest)
//	webResponse := response.Response{
//		Code:   http.StatusOK,
//		Status: "Ok",
//		Data:   nil,
//	}
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
//func (controller *ItemsController) Update(ctx *gin.Context) {
//	log.Info().Msg("update tags")
//	updateTagsRequest := request.UpdateTagsRequest{}
//	err := ctx.ShouldBindJSON(&updateTagsRequest)
//	helper.ErrorPanic(err)
//
//	tagId := ctx.Param("tagId")
//	id, err := strconv.Atoi(tagId)
//	helper.ErrorPanic(err)
//	updateTagsRequest.Id = id
//
//	controller.tagsService.Update(updateTagsRequest)
//
//	webResponse := response.Response{
//		Code:   http.StatusOK,
//		Status: "Ok",
//		Data:   nil,
//	}
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
//func (controller *ItemsController) Delete(ctx *gin.Context) {
//	log.Info().Msg("delete tags")
//	tagId := ctx.Param("tagId")
//	id, err := strconv.Atoi(tagId)
//	helper.ErrorPanic(err)
//	controller.tagsService.Delete(id)
//
//	webResponse := response.Response{
//		Code:   http.StatusOK,
//		Status: "Ok",
//		Data:   nil,
//	}
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
//func (controller *ItemsController) FindById(ctx *gin.Context) {
//	log.Info().Msg("findbyid tags")
//	tagId := ctx.Param("tagId")
//	id, err := strconv.Atoi(tagId)
//	helper.ErrorPanic(err)
//
//	tagResponse := controller.tagsService.FindById(id)
//
//	webResponse := response.Response{
//		Code:   http.StatusOK,
//		Status: "Ok",
//		Data:   tagResponse,
//	}
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
//func (controller *ItemsController) FindAll(ctx *gin.Context) {
//	log.Info().Msg("findAll tags")
//	tagResponse := controller.tagsService.FindAll()
//	webResponse := response.Response{
//		Code:   http.StatusOK,
//		Status: "Ok",
//		Data:   tagResponse,
//	}
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//
//}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *ItemsController) ProcessItemsFile(ctx *gin.Context) {
	log.Info().Msg("Processing items file")
	//tagResponse := controller.tagsService.FindAll()
	//webResponse := response.Response{
	//	Code:   http.StatusOK,
	//	Status: "Ok",
	//	Data:   tagResponse,
	//}
	//ctx.Header("Content-Type", "application/json")
	//ctx.JSON(http.StatusOK, webResponse)

	// Retrieve the uploaded file from the request
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errProcess := controller.tagsService.ProcessItemsFile(file)
	if errProcess != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errProcess.Error()})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "CSV file uploaded and records inserted successfully"})
}
