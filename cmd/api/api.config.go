package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/duckcoding00/blog/internal/handler"
	"github.com/gorilla/mux"
)

type Application struct {
	router *mux.Router
	config AppConfig
}

type AppConfig struct {
	handler handler.Handler
	addr    string
}

func NewApp(config AppConfig) *Application {
	return &Application{
		router: mux.NewRouter(),
		config: config,
	}
}

func (a *Application) RegisterRouter() {
	apiRouter := a.router.PathPrefix("/api/v1").Subrouter()

	// health
	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "ok",
		})
	}).Methods("GET")

	// blog
	blogRouter := apiRouter.PathPrefix("/blog").Subrouter()

	blogRouter.HandleFunc("/", a.config.handler.Blog.Create).Methods("POST")
	blogRouter.HandleFunc("/{id}", a.config.handler.Blog.Update).Methods("PATCH")
	blogRouter.HandleFunc("/{id}", a.config.handler.Blog.GetBlog).Methods("GET")
	blogRouter.HandleFunc("/{id}", a.config.handler.Blog.Delete).Methods("DELETE")
	blogRouter.HandleFunc("/", a.config.handler.Blog.GetBlogs).Methods("GET")
}

func (a *Application) Run() {
	log.Printf("server running on %s", a.config.addr)
	if err := http.ListenAndServe(a.config.addr, a.router); err != nil {
		log.Fatal("server failed : ", err)
	}
}
