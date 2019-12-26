package domain

type Note struct {
	ID        string
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
	Author    User
}

type Notes []Note
