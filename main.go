package main

import (
	"log"
	"net/http"
	"time"

	_ "./docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Ejemplo Web 001
// @version 1.0
// @description Ejemplo uso Chi Go para pruebas de concepto
// @termsOfService http://swagger.io/terms/

// @contact.name Zeke
// @contact.url http://www.zeke.cl
// @contact.email correomandos@zeke.cl

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

//Routes Acciones principales para el ruteo Chi
func Routes() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(120 * time.Second))
	return r
}

func main() {

	r := Routes()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	r.Get("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	log.Fatal(http.ListenAndServe(":3000", r))

}
