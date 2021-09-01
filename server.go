package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/horri1520/hori-api/config"
	"github.com/horri1520/hori-api/db"
	"github.com/horri1520/hori-api/handler"
	"github.com/horri1520/hori-api/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

type Server struct {
	db           *sqlx.DB
	router       *mux.Router
	envVariables *config.EnvVariables
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(envVariables *config.EnvVariables) error {
	s.envVariables = envVariables

	cs := db.NewPostgreSQL(envVariables.DatabaseUrl)
	dbcon, err := cs.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}

	s.db = dbcon
	s.router = s.Route()

	return nil
}

func (s *Server) Route() *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins: []string{s.envVariables.AccessControlAllowOrigin},
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		corsMiddleware.Handler,
	)

	r := mux.NewRouter()

	v1r := r.PathPrefix("/v1").Subrouter()

	markdownPostUsecase := usecase.NewMarkdownPostUsecase(s.db)
	markdownPostHandler := handler.NewMarkdownPostHandler(markdownPostUsecase)

	v1r.Methods(http.MethodGet, http.MethodOptions).Path("/markdown_posts/{markdown_post_id}").Handler(commonChain.Then(AppHandler{markdownPostHandler.Show}))
	v1r.Methods(http.MethodGet, http.MethodOptions).Path("/markdown_posts").Handler(commonChain.Then(AppHandler{markdownPostHandler.Index}))
	v1r.Methods(http.MethodPost, http.MethodOptions).Path("/markdown_posts").Handler(commonChain.Then(AppHandler{markdownPostHandler.Create}))

	return r
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}
