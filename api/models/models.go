package models

type Blog struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Date      string `json:"date"`
    MainImage string `json:"main_image"`
    Subheading string `json:"subheading"`
    Body      string `json:"body"`
}

type Carousel struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Image string `json:"image"`
}

type Employee struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Email        string `json:"email"`
    Role         string `json:"role"`
    Bio          string `json:"bio"`
    ProfileImage string `json:"profile_image"`
    Created      string `json:"created"`
    Phone        string `json:"phone"`
}

type Machine struct {
    ID            string `json:"id"`
    SupplierID    string `json:"supplierId"`
    Name          string `json:"name"`
    MachineImage  string `json:"machine_image"`
    Description   string `json:"description"`
    MachineLink   string `json:"machine_link"`
}

type Product struct {
    ID            string `json:"id"`
    MachineID     string `json:"machineId"`
    Name          string `json:"name"`
    ProductImage  string `json:"product_image"`
    Description   string `json:"description"`
    ProductLink   string `json:"product_link"`
}

type Sparepart struct {
    ID              string `json:"id"`
    SupplierID      string `json:"supplierId"`
    Name            string `json:"name"`
    PartsImage      string `json:"parts_image"`
    SparePartsLink  string `json:"spare_parts_link"`
    PdfLink         string `json:"pdf_link"`
}

type Supplier struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	LogoImage       string `json:"logo_image"`
	MarketingImage  string `json:"marketing_image"`
	SocialFacebook  string `json:"social_facebook"`
	SocialInstagram string `json:"social_instagram"`
	SocialLinkedin  string `json:"social_linkedin"`
	SocialTwitter   string `json:"social_twitter"`
	SocialYoutube   string `json:"social_youtube"`
	SocialWebsite   string `json:"social_website"`
	Created         string `json:"created"`
}

type Timeline struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Date    string `json:"date"`
    Body    string `json:"body"`
    Created string `json:"created"`
}

type User struct {
    UID  string `json:"uid"`
    Role string `json:"role"`
}

type YoutubeVideo struct {
    ID      string `json:"id"`
    Snippet struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Thumbnails  struct {
            Medium struct {
                URL string `json:"url"`
            } `json:"medium"`
        } `json:"thumbnails"`
    } `json:"snippet"`
}

type YoutubeApiResponse struct {
    Data struct {
        Items []YoutubeVideo `json:"items"`
    } `json:"data"`
}

type Video struct {
    ID            string `json:"id"`
    SupplierID    string `json:"supplierId"`
    WebURL        string `json:"web_url"`
    Title         string `json:"title"`
    Description   string `json:"description"`
    VideoID       string `json:"video_id"`
    ThumbnailURL  string `json:"thumbnail_url"`
}
