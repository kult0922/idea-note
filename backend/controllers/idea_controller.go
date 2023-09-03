package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kult0922/idea-note/apperrors"
	"github.com/kult0922/idea-note/controllers/services"
	"github.com/kult0922/idea-note/models"
)

type IdeaController struct {
	service services.IdeaServicer
}

func NewIdeaController(s services.IdeaServicer) *IdeaController {
	return &IdeaController{service: s}
}

// GET /hello のハンドラ
func (c *IdeaController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// POST /Idea のハンドラ
func (c *IdeaController) PostIdeaHandler(w http.ResponseWriter, req *http.Request) {

	var reqIdea models.IdeaRequest
	if err := json.NewDecoder(req.Body).Decode(&reqIdea); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	Idea, err := c.service.PostIdeaService(reqIdea)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(Idea)
}

// GET /Idea/list のハンドラ
func (c *IdeaController) IdeaListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	// クエリパラメータpageを取得
	var userId int
	if p, ok := queryMap["user_id"]; ok && len(p) > 0 {
		var err error
		userId, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be valid user_id")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		apperrors.ErrorHandler(w, req, errors.New("queryparam must be valid user_id"))
		return
	}

	IdeaList, err := c.service.GetIdeaListService(userId)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(IdeaList)
}

// GET /Idea/{id} のハンドラ
func (c *IdeaController) IdeaDetailHandler(w http.ResponseWriter, req *http.Request) {
	var publicIdeaId string
	if p, ok := mux.Vars(req)["id"]; ok && len(p) > 0 {
		publicIdeaId = p
	} else {
		apperrors.ErrorHandler(w, req, errors.New("queryparam must be valid idea_id"))
		return
	}
	Idea, err := c.service.GetIdeaService(publicIdeaId)

	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(Idea)
}
