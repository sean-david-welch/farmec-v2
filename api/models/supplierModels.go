package models

import (
	"database/sql"
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
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	LogoImage       string `json:"logo_image"`
	MarketingImage  sql.NullString `json:"marketing_image"`
	SocialFacebook  sql.NullString `json:"social_facebook"`
	SocialInstagram sql.NullString `json:"social_instagram"`
	SocialLinkedin  sql.NullString `json:"social_linkedin"`
	SocialTwitter   sql.NullString `json:"social_twitter"`
	SocialYoutube   sql.NullString `json:"social_youtube"`
	SocialWebsite   sql.NullString `json:"social_website"`
	Created         time.Time `json:"created"`
}