package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taran/internal/core/domain"
	"taran/internal/core/interfaces"
)

type BooksHandler struct {
	uc interfaces.BooksUseCase
}

func NewBooksHandler(uc interfaces.BooksUseCase) *BooksHandler {
	return &BooksHandler{
		uc: uc,
	}
}

func (h *BooksHandler) GetList(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = DefaultLimit
	}
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	result, err := h.uc.GetList(uint32(limit), uint32(offset))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *BooksHandler) BulkDelete(ctx *gin.Context) {
	var request struct {
		Data []*domain.Book `json:"data"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.uc.BulkDelete(request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (h *BooksHandler) Update(ctx *gin.Context) {
	var request struct {
		Data domain.Book `json:"data"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	request.Data.ID = ctx.Param("id")

	if err := h.uc.Update(&request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"data": request.Data})
}
