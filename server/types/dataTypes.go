package types

type SupplierResult struct {
    PresignedLogoUrl string
	LogoUrl string
	PresginedMarketingUrl string
	MarketingUrl string 
}

type ModelResult struct {
	PresginedUrl string
	ImageUrl string
}

type EmailData struct {
    Name string
    Email string
    Message string
}