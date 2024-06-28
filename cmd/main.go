package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"spy-cats/internal/handler"
	"spy-cats/internal/repository"
	"spy-cats/internal/service"
)

func main() {
	ctx := context.Background()

	connStr := "postgres://postgres:pass@localhost:5436/spy_cats_agency"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	r := gin.Default()

	spyCatRepository := repository.NewSpyCatRepository(conn)
	spyCatService := service.NewSpyCatService(spyCatRepository)
	spyCatHandler := handler.NewSpyCatCRUDHandler(spyCatService)
	spyCatHandler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
