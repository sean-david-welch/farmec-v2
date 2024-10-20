package lib

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

func boolToInt64(b bool) sql.NullInt64 {
	if b {
		return sql.NullInt64{Int64: 1, Valid: true}
	}
	return sql.NullInt64{Int64: 0, Valid: true}
}

func DeserializeSupplier(supplier types.Supplier) db.Supplier {
	return db.Supplier{
		ID:              supplier.ID,
		Name:            supplier.Name,
		LogoImage:       sql.NullString{String: supplier.LogoImage, Valid: supplier.LogoImage != ""},
		MarketingImage:  sql.NullString{String: supplier.MarketingImage, Valid: supplier.MarketingImage != ""},
		Description:     sql.NullString{String: supplier.Description, Valid: supplier.Description != ""},
		SocialFacebook:  sql.NullString{String: *supplier.SocialFacebook, Valid: supplier.SocialFacebook != nil},
		SocialInstagram: sql.NullString{String: *supplier.SocialInstagram, Valid: supplier.SocialInstagram != nil},
		SocialLinkedin:  sql.NullString{String: *supplier.SocialLinkedin, Valid: supplier.SocialLinkedin != nil},
		SocialTwitter:   sql.NullString{String: *supplier.SocialTwitter, Valid: supplier.SocialTwitter != nil},
		SocialYoutube:   sql.NullString{String: *supplier.SocialYoutube, Valid: supplier.SocialYoutube != nil},
		SocialWebsite:   sql.NullString{String: *supplier.SocialWebsite, Valid: supplier.SocialWebsite != nil},
		Created:         sql.NullString{String: supplier.Created, Valid: supplier.Created != ""},
	}
}

func DeserializeBlog(blog types.Blog) db.Blog {
	return db.Blog{
		ID:         blog.ID,
		Title:      blog.Title,
		Date:       sql.NullString{String: blog.Date, Valid: blog.Date != ""},
		MainImage:  sql.NullString{String: blog.MainImage, Valid: blog.MainImage != ""},
		Subheading: sql.NullString{String: blog.Subheading, Valid: blog.Subheading != ""},
		Body:       sql.NullString{String: blog.Body, Valid: blog.Body != ""},
		Created:    sql.NullString{String: blog.Created, Valid: blog.Created != ""},
	}
}

func DeserializeCarousel(carousel types.Carousel) db.Carousel {
	return db.Carousel{
		ID:      carousel.ID,
		Name:    carousel.Name,
		Image:   sql.NullString{String: carousel.Image, Valid: carousel.Image != ""},
		Created: sql.NullString{String: carousel.Created, Valid: carousel.Created != ""},
	}
}

func DeserializeEmployee(employee types.Employee) db.Employee {
	return db.Employee{
		ID:           employee.ID,
		Name:         employee.Name,
		Email:        employee.Email,
		Role:         employee.Role,
		ProfileImage: sql.NullString{String: employee.ProfileImage, Valid: employee.ProfileImage != ""},
		Created:      sql.NullString{String: employee.Created, Valid: employee.Created != ""},
	}
}

func DeserializeExhibition(exhibition types.Exhibition) db.Exhibition {
	return db.Exhibition{
		ID:       exhibition.ID,
		Title:    exhibition.Title,
		Date:     sql.NullString{String: exhibition.Date, Valid: exhibition.Date != ""},
		Location: sql.NullString{String: exhibition.Location, Valid: exhibition.Location != ""},
		Info:     sql.NullString{String: exhibition.Info, Valid: exhibition.Info != ""},
		Created:  sql.NullString{String: exhibition.Created, Valid: exhibition.Created != ""},
	}
}

func DeserializeLineItem(lineItem types.LineItem) db.LineItem {
	return db.LineItem{
		ID:    lineItem.ID,
		Name:  lineItem.Name,
		Price: lineItem.Price,
		Image: sql.NullString{String: lineItem.Image, Valid: lineItem.Image != ""},
	}
}

func DeserializeMachine(machine types.Machine) db.Machine {
	return db.Machine{
		ID:           machine.ID,
		SupplierID:   machine.SupplierID,
		Name:         machine.Name,
		MachineImage: sql.NullString{String: machine.MachineImage, Valid: machine.MachineImage != ""},
		Description:  sql.NullString{String: *machine.Description, Valid: machine.Description != nil},
		MachineLink:  sql.NullString{String: *machine.MachineLink, Valid: machine.MachineLink != nil},
		Created:      sql.NullString{String: machine.Created, Valid: machine.Created != ""},
	}
}

func DeserializeMachineRegistration(reg types.MachineRegistration) db.MachineRegistration {
	return db.MachineRegistration{
		ID:               reg.ID,
		DealerName:       reg.DealerName,
		DealerAddress:    sql.NullString{String: reg.DealerAddress, Valid: reg.DealerAddress != ""},
		OwnerName:        reg.OwnerName,
		OwnerAddress:     sql.NullString{String: reg.OwnerAddress, Valid: reg.OwnerAddress != ""},
		MachineModel:     reg.MachineModel,
		SerialNumber:     reg.SerialNumber,
		InstallDate:      sql.NullString{String: reg.InstallDate, Valid: reg.InstallDate != ""},
		InvoiceNumber:    sql.NullString{String: reg.InvoiceNumber, Valid: reg.InvoiceNumber != ""},
		CompleteSupply:   boolToInt64(reg.CompleteSupply),
		PdiComplete:      boolToInt64(reg.PdiComplete),
		PtoCorrect:       boolToInt64(reg.PtoCorrect),
		MachineTestRun:   boolToInt64(reg.MachineTestRun),
		SafetyInduction:  boolToInt64(reg.SafetyInduction),
		OperatorHandbook: boolToInt64(reg.OperatorHandbook),
		Date:             sql.NullString{String: reg.Date, Valid: reg.Date != ""},
		CompletedBy:      sql.NullString{String: reg.CompletedBy, Valid: reg.CompletedBy != ""},
		Created:          sql.NullString{String: reg.Created, Valid: reg.Created != ""},
	}
}

func DeserializePartsRequired(parts types.PartsRequired) db.PartsRequired {
	return db.PartsRequired{
		ID:             parts.ID,
		WarrantyID:     parts.WarrantyID,
		PartNumber:     sql.NullString{String: parts.PartNumber, Valid: parts.PartNumber != ""},
		QuantityNeeded: parts.QuantityNeeded,
		InvoiceNumber:  sql.NullString{String: parts.InvoiceNumber, Valid: parts.InvoiceNumber != ""},
		Description:    sql.NullString{String: parts.Description, Valid: parts.Description != ""},
	}
}

func DeserializePrivacy(privacy types.Privacy) db.Privacy {
	return db.Privacy{
		ID:      privacy.ID,
		Title:   privacy.Title,
		Body:    sql.NullString{String: privacy.Body, Valid: privacy.Body != ""},
		Created: sql.NullString{String: privacy.Created, Valid: privacy.Created != ""},
	}
}

func DeserializeProduct(product types.Product) db.Product {
	return db.Product{
		ID:           product.ID,
		MachineID:    product.MachineID,
		Name:         product.Name,
		ProductImage: sql.NullString{String: product.ProductImage, Valid: product.ProductImage != ""},
		Description:  sql.NullString{String: product.Description, Valid: product.Description != ""},
		ProductLink:  sql.NullString{String: product.ProductLink, Valid: product.ProductLink != ""},
	}
}

func DeserializeSparePart(sparePart types.Sparepart) db.SparePart {
	return db.SparePart{
		ID:             sparePart.ID,
		SupplierID:     sparePart.SupplierID,
		Name:           sparePart.Name,
		PartsImage:     sql.NullString{String: sparePart.PartsImage, Valid: sparePart.PartsImage != ""},
		SparePartsLink: sql.NullString{String: sparePart.SparePartsLink, Valid: sparePart.SparePartsLink != ""},
	}
}

func DeserializeTerm(term types.Terms) db.Term {
	return db.Term{
		ID:      term.ID,
		Title:   term.Title,
		Body:    sql.NullString{String: term.Body, Valid: term.Body != ""},
		Created: sql.NullString{String: term.Created, Valid: term.Created != ""},
	}
}

func DeserializeTimeline(timeline types.Timeline) db.Timeline {
	return db.Timeline{
		ID:      timeline.ID,
		Title:   timeline.Title,
		Date:    sql.NullString{String: timeline.Date, Valid: timeline.Date != ""},
		Body:    sql.NullString{String: timeline.Body, Valid: timeline.Body != ""},
		Created: sql.NullString{String: timeline.Created, Valid: timeline.Created != ""},
	}
}

func DeserializeVideo(video types.VideoRequest) db.Video {
	return db.Video{
		SupplierID: video.SupplierID,
		WebUrl:     sql.NullString{String: video.WebURL, Valid: video.WebURL != ""},
	}
}

func DeserializeWarrantyClaim(claim types.WarrantyClaim) db.WarrantyClaim {
	return db.WarrantyClaim{
		ID:             claim.ID,
		Dealer:         claim.Dealer,
		DealerContact:  sql.NullString{String: *claim.DealerContact, Valid: claim.DealerContact != nil},
		OwnerName:      *claim.OwnerName,
		MachineModel:   *claim.MachineModel,
		SerialNumber:   *claim.SerialNumber,
		InstallDate:    sql.NullString{String: *claim.InstallDate, Valid: claim.InstallDate != nil},
		FailureDate:    sql.NullString{String: *claim.FailureDate, Valid: claim.FailureDate != nil},
		RepairDate:     sql.NullString{String: *claim.RepairDate, Valid: claim.RepairDate != nil},
		FailureDetails: sql.NullString{String: *claim.FailureDetails, Valid: claim.FailureDetails != nil},
		RepairDetails:  sql.NullString{String: *claim.RepairDetails, Valid: claim.RepairDetails != nil},
		LabourHours:    sql.NullString{String: *claim.LabourHours, Valid: claim.LabourHours != nil},
		CompletedBy:    sql.NullString{String: *claim.CompletedBy, Valid: claim.CompletedBy != nil},
		Created:        sql.NullString{String: claim.Created, Valid: claim.Created != ""},
	}
}
