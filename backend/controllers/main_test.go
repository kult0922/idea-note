package controllers_test

import (
	"testing"

	"github.com/kult0922/idea-note/controllers"
	"github.com/kult0922/idea-note/controllers/testdata"
)

var iCon *controllers.IdeaController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	iCon = controllers.NewIdeaController(ser)

	m.Run()
}
