package models

import (
	"time"
)

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
    ID              string  `json:"id"`
    Name            string  `json:"name"`
    LogoImage       string  `json:"logo_image"`
    MarketingImage  string `json:"marketing_image"`
    Description     string  `json:"description"`
    SocialFacebook  *string `json:"social_facebook"`
    SocialInstagram *string `json:"social_instagram"`
    SocialLinkedin  *string `json:"social_linkedin"`
    SocialTwitter   *string `json:"social_twitter"`
    SocialYoutube   *string `json:"social_youtube"`
    SocialWebsite   *string `json:"social_website"`
    Created         time.Time `json:"created"`
}