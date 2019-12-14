package domain

type Note struct {
	ID        string
	Title     string
	Author    User
	Content   string
	CreatedAt string
	UpdatedAt string
}

type Notes []Note
