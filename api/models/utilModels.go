package models

type Carousel struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Image string `json:"image"`
}

type User struct {
    UID  string `json:"uid"`
    Role string `json:"role"`
}


