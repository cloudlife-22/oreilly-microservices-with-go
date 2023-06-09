// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Article struct {
	ArticleID   string `json:"article_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	AuthorEmail string `json:"author_email"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Bio   *string `json:"bio,omitempty"`
}
