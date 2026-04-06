package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sMteX/necro-prestige-planner/internal/api"
	"github.com/sMteX/necro-prestige-planner/internal/db/sqlc"
)

//go:embed all:ui/dist
var uiAssets embed.FS

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

	// API endpoints
	r.Get("/health", api.HealthHandler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/recalculate", api.RecalculateHandler)
		r.Get("/plans", apiHandler.ListPlansHandler)
		r.Get("/plans/{id}", apiHandler.GetPlanHandler)
		r.Post("/plans", apiHandler.SavePlanHandler)
		r.Delete("/plans/{id}", apiHandler.DeletePlanHandler)
	})

	// Serve React app
	uiFS, err := fs.Sub(uiAssets, "ui/dist")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(uiFS))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			fileServer.ServeHTTP(w, r)
			return
		}

		// Try to open the file to see if it exists in the embedded FS
		f, err := uiFS.Open(strings.TrimPrefix(path, "/"))
		if err == nil {
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		// Fallback to index.html for SPA routing
		content, err := fs.ReadFile(uiFS, "index.html")
		if err != nil {
			http.Error(w, "index.html not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(content)
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
