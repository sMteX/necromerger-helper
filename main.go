package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sMteX/necro-prestige-planner/internal/api"
	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://necro:necro@localhost:5432/necro_prestige?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := sqlc.New(db)
	apiHandler := &api.Handler{Queries: queries}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", api.HealthHandler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/recalculate", api.RecalculateHandler)
		r.Get("/plans", apiHandler.ListPlansHandler)
		r.Get("/plans/{id}", apiHandler.GetPlanHandler)
		r.Post("/plans", apiHandler.SavePlanHandler)
		r.Delete("/plans/{id}", apiHandler.DeletePlanHandler)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}