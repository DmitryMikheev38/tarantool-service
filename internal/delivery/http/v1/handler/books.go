package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *BooksHandler) Create(ctx *gin.Context) {
	var request struct {
		Data domain.Book `json:"data"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := h.uc.Create(&request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, request)
}

func (h *BooksHandler) GetList(ctx *gin.Context) {
	result, err := h.uc.GetList()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
