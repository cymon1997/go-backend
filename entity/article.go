package entity

type Article struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Content     string `json:"content" db:"content"`
	Timestamp
}

type InsertArticleResponse struct {
	ID string `json:"id"`
	Response
}

type GetArticleRequest struct {
	ID string `json:"id"`
}

type GetArticleResponse struct {
	Data Article `json:"data"`
	Response
}
