package main

import (
	"context"
	"log"
	"net/http"

	//"github.com/estacet/spy-cats/internal/repository"
	//"github.com/estacet/spy-cats/internal/service"
	"github.com/jackc/pgx/v5"
	//
	//"github.com/estacet/spy-cats/internal/handler"
)

func main() {
	ctx := context.Background()

	connStr := "postgres://postgres:pass@localhost:5433/spy_cats"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	raceRepository := repository.NewRaceRepository(conn)

	raceService := service.NewRaceService(raceRepository)

	healthCheckHandler := handler.NewHealthCheckHandler()
	raceCRUDHandler := handler.NewRaceCRUDHandler(raceService)

	mux := http.NewServeMux()

	mux.Handle("/health", healthCheckHandler)
	mux.Handle("/race", raceCRUDHandler)

	apiServer := server.NewAPIServer(mux)
	if err := apiServer.Start(); err != nil {
		log.Panic(err)
	}
}
