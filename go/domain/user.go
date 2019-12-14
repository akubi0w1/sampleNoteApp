package domain

type User struct {
	ID        string
	Name      string
	Mail      string
	Password  string
	CreatedAt string
}

type Users []User
