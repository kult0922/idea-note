package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kult0922/idea-note/models"
	"github.com/kult0922/idea-note/repositories"
	"github.com/kult0922/idea-note/repositories/testdata"
)

// SelectIdeaList関数のテスト
func TestSelectIdeaList(t *testing.T) {
	expectedNum := 1
	got, err := repositories.SelectIdeaList(testDB, testdata.IdeaTestData[0].UserId)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d Ideas\n", expectedNum, num)
	}
}

// SelectIdeaDetail関数のテスト
func TestSelectIdeaDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Idea
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.IdeaTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.IdeaTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectIdeaDetail(testDB, test.expected.PublicId)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
		})
	}
}

// InsertIdea関数のテスト
func TestInsertIdea(t *testing.T) {
	Idea := models.Idea{
		Title:    "insertTest",
		Contents: "testest",
	}

	newIdea, err := repositories.InsertIdea(testDB, Idea)
	if err != nil {
		t.Error(err)
	}
	if newIdea.Title != Idea.Title {
		t.Errorf("new Idea title is expected %s but got %s\n", Idea.Title, newIdea.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from Ideas
			where title = ? and contents = ?
		`
		testDB.Exec(sqlStr, Idea.Title, Idea.Contents)
	})
}
