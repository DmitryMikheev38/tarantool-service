package v1

import (
	"github.com/gin-gonic/gin"
	"taran/internal/core/interfaces"
	"taran/internal/delivery/http/v1/handler"
)

func RegisterHTTPEndpoints(router *gin.Engine, booksUC interfaces.BooksUseCase, authorsUC interfaces.AuthorsUseCase) {

	booksHandler := handler.NewBooksHandler(booksUC)
	authorsHandler := handler.NewAuthorsHandler(authorsUC)

	api := router.Group("/api/v1")
	{
		api.POST("/authors", authorsHandler.Create)
		api.GET("/authors", authorsHandler.GetList)
		api.POST("/authors/:id/books", authorsHandler.AddBook)
		api.POST("/books", booksHandler.Create)
		api.GET("/books", booksHandler.GetList)
	}
}
