package services

import "github.com/kult0922/idea-note/models"

// /Idea関連を引き受けるサービス
type IdeaServicer interface {
	PostIdeaService(Idea models.IdeaRequest) (models.IdeaResponse, error)
	GetIdeaListService(UserId int) ([]models.IdeaResponse, error)
	GetIdeaService(IdeaPublicId string) (models.IdeaResponse, error)
}

// /Userを引き受けるサービス
type UserServicer interface {
	PostUserService(User models.User) (models.User, error)
}
