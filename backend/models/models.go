package models

import "time"

type User struct {
	Id       int    `json:"idea_id"`
	PublicId string `json:"public_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Idea struct {
	Id        int       `json:"id"`
	PublicId  string    `json:"public_id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	WrittenAt time.Time `json:"written_at"`
}

type IdeaRequest struct {
	UserId   int    `json:"user_id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

type IdeaResponse struct {
	PublicId  string    `json:"public_id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	WrittenAt time.Time `json:"written_at"`
}
