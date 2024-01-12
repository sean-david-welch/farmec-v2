package models

import "time"

type Blog struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Date      string `json:"date"`
    MainImage string `json:"main_image"`
    Subheading string `json:"subheading"`
    Body      string `json:"body"`
}

type Exhibition struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Date      string    `json:"date"`
    Location  string    `json:"location"`
    Info      string    `json:"info"`
    Created   time.Time `json:"created"`
}
