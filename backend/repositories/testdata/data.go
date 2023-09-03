package testdata

import "github.com/kult0922/idea-note/models"

var IdeaTestData = []models.Idea{
	models.Idea{
		Id:       1,
		PublicId: "704d9697-4c5e-c548-059e-3d67bd90e070",
		UserId:   1,
		Title:    "first idea",
		Contents: "This is first idea",
	},
	models.Idea{
		Id:       2,
		PublicId: "a52bc272-71aa-2017-4d0a-f30d21b56c2a",
		UserId:   2,
		Title:    "second idea",
		Contents: "This is second idea",
	},
}

var UserTestData = []models.User{
	models.User{
		Id:       1,
		PublicId: "80898b68-86c7-13bb-9fac-2a910dec44db",
		Name:     "1st User",
		Email:    "abc@gmail.com",
	},
	models.User{
		Id:       2,
		PublicId: "a8f92b26-a7fb-e849-11bd-c2692b1650c5",
		Name:     "1st User",
		Email:    "abc@gmail.com",
	},
}
