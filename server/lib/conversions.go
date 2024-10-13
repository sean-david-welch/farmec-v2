package lib

import (
	"database/sql"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

func int64ToBool(i sql.NullInt64) bool {
	val := i.Int64 != 0
	return val
}

func convertSupplier(supplier db.Supplier) types.Supplier {
	return types.Supplier{
		ID:              supplier.ID,
		Name:            supplier.Name,
		LogoImage:       supplier.LogoImage.String,
		MarketingImage:  supplier.MarketingImage.String,
		Description:     supplier.Description.String,
		SocialFacebook:  &supplier.SocialFacebook.String,
		SocialInstagram: &supplier.SocialInstagram.String,
		SocialLinkedin:  &supplier.SocialLinkedin.String,
		SocialTwitter:   &supplier.SocialTwitter.String,
		SocialYoutube:   &supplier.SocialYoutube.String,
		SocialWebsite:   &supplier.SocialWebsite.String,
		Created:         supplier.Created.String,
	}
}

func convertBlog(blog db.Blog) types.Blog {
	return types.Blog{
		ID:         blog.ID,
		Title:      blog.Title,
		Date:       blog.Date.String,
		MainImage:  blog.MainImage.String,
		Subheading: blog.Subheading.String,
		Body:       blog.Body.String,
		Created:    blog.Created.String,
	}
}

func convertCarousel(carousel db.Carousel) types.Carousel {
	return types.Carousel{
		ID:      carousel.ID,
		Name:    carousel.Name,
		Image:   carousel.Image.String,
		Created: carousel.Created.String,
	}
}

func convertEmployee(employee db.Employee) types.Employee {
	return types.Employee{
		ID:           employee.ID,
		Name:         employee.Name,
		Email:        employee.Email,
		Role:         employee.Role,
		ProfileImage: employee.ProfileImage.String,
		Created:      employee.Created.String,
	}
}

func convertExhibition(exhibition db.Exhibition) types.Exhibition {
	return types.Exhibition{
		ID:       exhibition.ID,
		Title:    exhibition.Title,
		Date:     exhibition.Date.String,
		Location: exhibition.Location.String,
		Info:     exhibition.Info.String,
		Created:  exhibition.Created.String,
	}
}

func convertLineItem(lineItem db.LineItem) types.LineItem {
	return types.LineItem{
		ID:    lineItem.ID,
		Name:  lineItem.Name,
		Price: lineItem.Price,
		Image: lineItem.Image.String,
	}
}

func convertMachine(machine db.Machine) types.Machine {
	return types.Machine{
		ID:           machine.ID,
		SupplierID:   machine.SupplierID,
		Name:         machine.Name,
		MachineImage: machine.MachineImage.String,
		Description:  &machine.Description.String,
		MachineLink:  &machine.MachineLink.String,
		Created:      machine.Created.String,
	}
}

func convertMachineRegistration(reg db.MachineRegistration) types.MachineRegistration {
	return types.MachineRegistration{
		ID:               reg.ID,
		DealerName:       reg.DealerName,
		DealerAddress:    reg.DealerAddress.String,
		OwnerName:        reg.OwnerName,
		OwnerAddress:     reg.OwnerAddress.String,
		MachineModel:     reg.MachineModel,
		SerialNumber:     reg.SerialNumber,
		InstallDate:      reg.InstallDate.String,
		InvoiceNumber:    reg.InvoiceNumber.String,
		CompleteSupply:   int64ToBool(reg.CompleteSupply),
		PdiComplete:      int64ToBool(reg.PdiComplete),
		PtoCorrect:       int64ToBool(reg.PtoCorrect),
		MachineTestRun:   int64ToBool(reg.MachineTestRun),
		SafetyInduction:  int64ToBool(reg.SafetyInduction),
		OperatorHandbook: int64ToBool(reg.OperatorHandbook),
		Date:             reg.Date.String,
		CompletedBy:      reg.CompletedBy.String,
		Created:          reg.Created.String,
	}
}

func convertPartsRequired(parts db.PartsRequired) types.PartsRequired {
	return types.PartsRequired{
		ID:             parts.ID,
		WarrantyID:     parts.WarrantyID,
		PartNumber:     parts.PartNumber.String,
		QuantityNeeded: parts.QuantityNeeded,
		InvoiceNumber:  parts.InvoiceNumber.String,
		Description:    parts.Description.String,
	}
}

func convertPrivacy(privacy db.Privacy) types.Privacy {
	return types.Privacy{
		ID:      privacy.ID,
		Title:   privacy.Title,
		Body:    privacy.Body.String,
		Created: privacy.Created.String,
	}
}

func convertProduct(product db.Product) types.Product {
	return types.Product{
		ID:           product.ID,
		MachineID:    product.MachineID,
		Name:         product.Name,
		ProductImage: product.ProductImage.String,
		Description:  product.Description.String,
		ProductLink:  product.ProductLink.String,
	}
}

func convertSparePart(sparePart db.SparePart) types.Sparepart {
	return types.Sparepart{
		ID:             sparePart.ID,
		SupplierID:     sparePart.SupplierID,
		Name:           sparePart.Name,
		PartsImage:     sparePart.PartsImage.String,
		SparePartsLink: sparePart.SparePartsLink.String,
	}
}

func convertTerm(term db.Term) types.Terms {
	return types.Terms{
		ID:      term.ID,
		Title:   term.Title,
		Body:    term.Body.String,
		Created: term.Created.String,
	}
}

func convertTimeline(timeline db.Timeline) types.Timeline {
	return types.Timeline{
		ID:      timeline.ID,
		Title:   timeline.Title,
		Date:    timeline.Date.String,
		Body:    timeline.Body.String,
		Created: timeline.Created.String,
	}
}

func convertVideo(video db.Video) types.Video {
	return types.Video{
		ID:           video.ID,
		SupplierID:   video.SupplierID,
		WebURL:       video.WebUrl.String,
		Title:        &video.Title.String,
		Description:  &video.Description.String,
		VideoID:      &video.VideoID.String,
		ThumbnailURL: &video.ThumbnailUrl.String,
		Created:      video.Created.String,
	}
}

func convertWarrantyClaim(claim db.WarrantyClaim) types.WarrantyClaim {
	return types.WarrantyClaim{
		ID:             claim.ID,
		Dealer:         claim.Dealer,
		DealerContact:  &claim.DealerContact.String,
		OwnerName:      &claim.OwnerName,
		MachineModel:   &claim.MachineModel,
		SerialNumber:   &claim.SerialNumber,
		InstallDate:    &claim.InstallDate.String,
		FailureDate:    &claim.FailureDate.String,
		RepairDate:     &claim.RepairDate.String,
		FailureDetails: &claim.FailureDetails.String,
		RepairDetails:  &claim.RepairDetails.String,
		LabourHours:    &claim.LabourHours.String,
		CompletedBy:    &claim.CompletedBy.String,
		Created:        claim.Created.String,
	}
}
