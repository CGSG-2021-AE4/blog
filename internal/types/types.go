package types

import (
	"time"

	"github.com/google/uuid"
)

// Articles
type ArticleDescr struct {
	Title          string    `json:"title"`
	AuthorId       uuid.UUID `json:"authorId"`
	AuthorUsername string    `json:"authorUsername"` // Made just for confort
}

type Article struct {
	Id uuid.UUID `json:"id"`
	ArticleDescr
	Created   time.Time `json:"created"`
	ContentId uuid.UUID `json:"contentId"`
}

var NilArticle Article

// Users

type User struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	// RegTime  time.Time `json:"regTime"`
	// LastSeen time.Time `json:"lastSeen`
	Username string `json:"username"`
	Password string `json:"password"`
}

var NilUser User
