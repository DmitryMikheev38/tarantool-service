package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taran/internal/core/domain"
	"taran/internal/core/interfaces"
)

const DefaultLimit = 10

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

	request.Data.AuthorID = ctx.Param("id")
	if err := h.uc.AddBook(&request.Data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": request.Data})
}

func (h *AuthorsHandler) GetList(ctx *gin.Context) {
	booksLimit, _ := strconv.Atoi(ctx.Query("booksLimit"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = DefaultLimit
	}
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	result, err := h.uc.GetList(booksLimit, uint32(limit), uint32(offset))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *AuthorsHandler) GetBooksList(ctx *gin.Context) {
	authorID := ctx.Param("id")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = DefaultLimit
	}
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	result, err := h.uc.GetBooksList(authorID, uint32(limit), uint32(offset))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
