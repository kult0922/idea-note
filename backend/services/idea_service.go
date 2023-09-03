package services

import (
	"database/sql"
	"errors"

	"github.com/kult0922/idea-note/apperrors"
	"github.com/kult0922/idea-note/models"
	"github.com/kult0922/idea-note/repositories"
)

// PostIdeaHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostIdeaService(ideaRequest models.IdeaRequest) (models.IdeaResponse, error) {
	idea := models.Idea{
		UserId:   ideaRequest.UserId,
		Title:    ideaRequest.Title,
		Contents: ideaRequest.Contents,
	}
	newIdea, err := repositories.InsertIdea(s.db, idea)

	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.IdeaResponse{}, err
	}

	ideaResponse := models.IdeaResponse{
		PublicId:  newIdea.PublicId,
		UserId:    newIdea.UserId,
		Title:     newIdea.Title,
		Contents:  newIdea.Contents,
		WrittenAt: newIdea.WrittenAt,
	}

	return ideaResponse, nil
}

// IdeaListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *MyAppService) GetIdeaListService(UserId int) ([]models.IdeaResponse, error) {
	ideaList, err := repositories.SelectIdeaList(s.db, UserId)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(ideaList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	var ideaResponseList []models.IdeaResponse
	for i := 0; i < len(ideaList); i++ {
		ideaResponseList = append(ideaResponseList, models.IdeaResponse{
			PublicId:  ideaList[i].PublicId,
			UserId:    ideaList[i].UserId,
			Title:     ideaList[i].Title,
			Contents:  ideaList[i].Contents,
			WrittenAt: ideaList[i].WrittenAt,
		})
	}

	return ideaResponseList, nil
}

// IdeaDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func (s *MyAppService) GetIdeaService(IdeaId string) (models.IdeaResponse, error) {
	Idea, err := repositories.SelectIdeaDetail(s.db, IdeaId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.IdeaResponse{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.IdeaResponse{}, err
	}

	ideaReaponse := models.IdeaResponse{
		PublicId:  Idea.PublicId,
		UserId:    Idea.UserId,
		Title:     Idea.Title,
		Contents:  Idea.Contents,
		WrittenAt: Idea.WrittenAt,
	}

	return ideaReaponse, nil
}
