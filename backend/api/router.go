package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kult0922/idea-note/api/middlewares"
	"github.com/kult0922/idea-note/controllers"
	"github.com/kult0922/idea-note/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	iCon := controllers.NewIdeaController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/idea", iCon.PostIdeaHandler).Methods(http.MethodPost)
	r.HandleFunc("/idea/list", iCon.IdeaListHandler).Methods(http.MethodGet)
	r.HandleFunc("/idea/{id}", iCon.IdeaDetailHandler).Methods(http.MethodGet)

	r.Use(middlewares.LoggingMiddleware)

	return r
}
