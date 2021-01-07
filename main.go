package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go-influxdb/modules/shop"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets")))

	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fileServer.ServeHTTP(w, r)
	})

	r.Get("/", shop.GetAllProducts)
	r.Get("/product/{ID}", shop.GetProduct)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/404.html")
	})

	log.Println("👉  Server started at 127.0.0.1:3333")
	http.ListenAndServe(":3333", r)
}
