package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taran/internal/core/domain"
	"taran/internal/core/domain/dto"
	"taran/internal/core/interfaces"
)

type AuthorsHandler struct {
	uc interfaces.AuthorsUseCase
}

func NewAuthorsHandler(uc interfaces.AuthorsUseCase) *AuthorsHandler {
	return &AuthorsHandler{
		uc: uc,
	}
}

func (h *AuthorsHandler) Create(ctx *gin.Context) {
	var request struct {
		Data domain.Author `json:"data"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := h.uc.Create(&request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, request)
}

func (h *AuthorsHandler) AddBook(ctx *gin.Context) {
	var request struct {
		Data domain.Book `json:"data"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	authorID := ctx.Param("id")
	if err := h.uc.AddBook(authorID, &request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": request.Data})
}

func (h *AuthorsHandler) GetList(ctx *gin.Context) {
	var params dto.GetListAuthorsParams

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	result, err := h.uc.GetList(params)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
