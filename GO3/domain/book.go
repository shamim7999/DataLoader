package domain

type Book struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	AuthorIds []string `json:"author_ids"`
}
