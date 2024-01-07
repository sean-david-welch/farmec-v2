package types

type SupplierResult struct {
    PresignedLogoUrl string
	LogoUrl string
	PresginedMarketingUrl string
	MarketingUrl string 
}

type DeleteSupplierResult struct {
	LogoUrl string
	MarketingUrl string
}

type MachineResult struct {
	PresginedMachine string
	MachineUrl string
}