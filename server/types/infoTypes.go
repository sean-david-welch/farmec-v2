package types

type Blog struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	MainImage  string `json:"main_image"`
	Subheading string `json:"subheading"`
	Body       string `json:"body"`
	Created    string `json:"created"`
}

type Exhibition struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Info     string `json:"info"`
	Created  string `json:"created"`
}

type Employee struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image"`
	Created      string `json:"created"`
}

type Timeline struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Body    string `json:"body"`
	Created string `json:"created"`
}
