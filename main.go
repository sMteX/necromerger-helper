package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sMteX/necro-prestige-planner/internal/api"
	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
	"github.com/swaggo/swag"

	_ "github.com/sMteX/necro-prestige-planner/docs"
)

//	@title			Necro Prestige Planner API
//	@version		1.0
//	@description	API for planning prestige resets and experiment spending in Necromerger

//	@host		localhost:8085
//	@BasePath	/

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
		r.Post("/resource-cap/{threshold}", api.ResourceCapHandler)
	})

	r.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		doc, err := swag.ReadDoc()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(doc))
	})

	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		doc, err := swag.ReadDoc()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecContent: doc,
			DarkMode:    true,
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Necro Prestige Planner API",
			},
			ShowSidebar: true,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(htmlContent))
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
