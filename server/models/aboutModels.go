package models

import "time"

type Employee struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Email        string `json:"email"`
    Role         string `json:"role"`
    Bio          *string `json:"bio"`
    ProfileImage string `json:"profile_image"`
    Created      time.Time `json:"created"`
    Phone        string `json:"phone"`
}

type Timeline struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Date    string `json:"date"`
    Body    string `json:"body"`
    Created string `json:"created"`
}