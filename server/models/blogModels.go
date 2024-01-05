package models

type Blog struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Date      string `json:"date"`
    MainImage string `json:"main_image"`
    Subheading string `json:"subheading"`
    Body      string `json:"body"`
}