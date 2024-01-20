package types

import "time"

type Employee struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	ProfileImage string    `json:"profile_image"`
	Created      time.Time `json:"created"`
}

type Timeline struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Date    string    `json:"date"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

type Privacy struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

type Terms struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}
