package types

type Machine struct {
	ID           string  `json:"id"`
	SupplierID   string  `json:"supplier_id"`
	Name         string  `json:"name"`
	MachineImage string  `json:"machine_image"`
	Description  *string `json:"description"`
	MachineLink  *string `json:"machine_link"`
	Created      string  `json:"created"`
}

type Product struct {
	ID           string `json:"id"`
	MachineID    string `json:"machine_id"`
	Name         string `json:"name"`
	ProductImage string `json:"product_image"`
	Description  string `json:"description"`
	ProductLink  string `json:"product_link"`
}

type Sparepart struct {
	ID             string `json:"id"`
	SupplierID     string `json:"supplier_id"`
	Name           string `json:"name"`
	PartsImage     string `json:"parts_image"`
	SparePartsLink string `json:"spare_parts_link"`
}

type Supplier struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	LogoImage       string  `json:"logo_image"`
	MarketingImage  string  `json:"marketing_image"`
	Description     string  `json:"description"`
	SocialFacebook  *string `json:"social_facebook"`
	SocialInstagram *string `json:"social_instagram"`
	SocialLinkedin  *string `json:"social_linkedin"`
	SocialTwitter   *string `json:"social_twitter"`
	SocialYoutube   *string `json:"social_youtube"`
	SocialWebsite   *string `json:"social_website"`
	Created         string  `json:"created"`
}

type Video struct {
	ID           string  `json:"id"`
	SupplierID   string  `json:"supplier_id"`
	WebURL       string  `json:"web_url"`
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	VideoID      *string `json:"video_id"`
	ThumbnailURL *string `json:"thumbnail_url"`
	Created      string  `json:"created"`
}

type VideoRequest struct {
	SupplierID string `json:"supplier_id"`
	WebURL     string `json:"web_url"`
}
