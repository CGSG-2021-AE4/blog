package types

import "github.com/google/uuid"

// Articles

type ArticleContent struct {
	Text string `json:"text"`
}

type ArticleHeader struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type Article struct {
	Header  ArticleHeader
	Content *ArticleContent // Pointer it is supposed to be long
}

type ArticleJson struct {
	ArticleHeader
	ArticleContent
}

// Users

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
