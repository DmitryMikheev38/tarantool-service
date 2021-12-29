package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"taran/infrastructure/config"
	"taran/internal/core/usecase"
	v1 "taran/internal/delivery/http/v1"
	"taran/internal/repository"
	"time"
)

func main() {
	appConfig := config.New()

	conn, err := tarantool.Connect(
		fmt.Sprintf("%s:%s", appConfig.TarantoolDB.Host, appConfig.TarantoolDB.Port),
		tarantool.Opts{
			User: appConfig.TarantoolDB.User,
			Pass: appConfig.TarantoolDB.Password,
		})

	if err != nil {
		log.Fatalf("Connection refused: ", err.Error())
	}

	defer conn.Close()

	booksRepository := repository.NewBooksRepository(conn)
	authorsRepository := repository.NewAuthorsRepository(conn)
	booksUseCase := usecase.NewBooksUseCase(booksRepository, authorsRepository)
	authorsUseCase := usecase.NewAuthorsUseCase(authorsRepository, booksRepository)

	router := gin.Default()

	v1.RegisterHTTPEndpoints(router, booksUseCase, authorsUseCase)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", appConfig.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
