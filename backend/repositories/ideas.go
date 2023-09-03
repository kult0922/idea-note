package repositories

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/kult0922/idea-note/models"
)

const (
	IdeaNumPerPage = 5
)

// 新規投稿をDBにinsertする関数
func InsertIdea(db *sql.DB, Idea models.Idea) (models.Idea, error) {
	const sqlStr = `
	insert into ideas (public_id, user_id, title, contents, written_at, created_at) values
	(?, ?, ?, ?, ?, now());
	`

	u, err := uuid.NewRandom()
	if err != nil {
		return models.Idea{}, err
	}
	uuid_ := u.String()

	// 特定のタイムゾーンでの現在時刻を取得
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return models.Idea{}, err
	}
	localTime := time.Now()
	tokyoTime := localTime.In(location)

	var newIdea models.Idea
	result, err := db.Exec(sqlStr, uuid_, Idea.UserId, Idea.Title, Idea.Contents, tokyoTime)
	if err != nil {
		return models.Idea{}, err
	}
	id, _ := result.LastInsertId()

	newIdea.Title, newIdea.Contents, newIdea.PublicId, newIdea.WrittenAt = Idea.Title, Idea.Contents, uuid_, tokyoTime

	newIdea.Id = int(id)

	return newIdea, nil
}

// Ideaのタイトル一覧をDBから取得する関数
func SelectIdeaList(db *sql.DB, userId int) ([]models.Idea, error) {
	const sqlStr = `
		select public_id, user_id, title, contents, written_at
		from ideas where user_id = ?;
	`

	rows, err := db.Query(sqlStr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	IdeaArray := make([]models.Idea, 0)
	for rows.Next() {
		var Idea models.Idea
		rows.Scan(&Idea.PublicId, &Idea.UserId, &Idea.Title, &Idea.Contents, &Idea.WrittenAt)

		IdeaArray = append(IdeaArray, Idea)
	}

	return IdeaArray, nil
}

// 投稿IDを指定して、記事データを取得する関数
func SelectIdeaDetail(db *sql.DB, IdeaPublicId string) (models.Idea, error) {
	const sqlStr = `
		select user_id, public_id, title, contents, written_at
		from ideas
		where public_id = ?;
	`
	row := db.QueryRow(sqlStr, IdeaPublicId)
	if err := row.Err(); err != nil {
		return models.Idea{}, err
	}

	var Idea models.Idea
	err := row.Scan(&Idea.UserId, &Idea.PublicId, &Idea.Title, &Idea.Contents, &Idea.WrittenAt)
	if err != nil {
		return models.Idea{}, err
	}

	return Idea, nil
}
