package types

type SupplierResult struct {
	PresignedLogoUrl      string
	LogoUrl               string
	PresignedMarketingUrl string
	MarketingUrl          string
}

type ModelResult struct {
	PresignedUrl string
	ImageUrl     string
}

type EmailData struct {
	Name    string
	Email   string
	Message string
}

type UserData struct {
	Email    string
	Password string
	Role     string
}

type LoginAuth struct {
	Username string
	Password string
}
