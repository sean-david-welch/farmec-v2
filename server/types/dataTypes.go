package types

type SupplierResult struct {
	PresignedLogoUrl      string `json:"presignedLogoUrl"`
	LogoUrl               string `json:"logoUrl"`
	PresignedMarketingUrl string `json:"presignedMarketingUrl"`
	MarketingUrl          string `json:"marketingUrl"`
}

type ModelResult struct {
	PresignedUrl string `json:"presignedUrl"`
	ImageUrl     string `json:"imageUrl"`
}

type PartsModelResult struct {
	PresignedUrl     string `json:"presignedUrl"`
	ImageUrl         string `json:"imageUrl"`
	PresignedLinkUrl string `json:"linkUrl"`
}

type EmailData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type UserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
