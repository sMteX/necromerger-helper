package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sMteX/necro-prestige-planner/internal/api"
	"github.com/sMteX/necro-prestige-planner/src/templates"
)

func main() {
	// Load .env if it exists
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Default to local dev config if not provided
		dbURL = "postgres://necro:necro@localhost:5432/necro_prestige?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Printf("Warning: Could not connect to database: %v", err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static files
	workDir, _ := os.Getwd()
	// When running from cmd/server, the root is two levels up
	// However, usually we run from project root.
	// Let's ensure it works if run from project root.
	assetsDir := filepath.Join(workDir, "assets")
	FileServer(r, "/assets", http.Dir(assetsDir))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Index("World")
		component.Render(r.Context(), w)
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HTMX is working! The server responded successfully."))
	})

	// API endpoints
	r.Get("/health", api.HealthHandler)
	r.Post("/recalculate", api.RecalculateHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := rctx.RoutePattern()[:len(rctx.RoutePattern())-2]
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
