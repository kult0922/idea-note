package testdata

import "github.com/kult0922/idea-note/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostIdeaService(Idea models.IdeaRequest) (models.IdeaResponse, error) {
	return IdeaTestData[1], nil
}

func (s *serviceMock) GetIdeaListService(userId int) ([]models.IdeaResponse, error) {
	return IdeaTestData, nil
}

func (s *serviceMock) GetIdeaService(publicIdeaId string) (models.IdeaResponse, error) {
	return IdeaTestData[0], nil
}

func (s *serviceMock) PostUserService(User models.User) (models.User, error) {
	return UserTestData[0], nil
}
