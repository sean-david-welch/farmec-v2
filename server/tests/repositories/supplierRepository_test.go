package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetSupplier(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock database: %s", err)
	}
	defer db.Close()

	facebookLink := "http://facebook.com/supplier"
	instagramLink := "http://instagram.com/supplier"
	linkedinLink := "http://linkedin.com/supplier"
	twitterLink := "http://twitter.com/supplier"
	youtubeLink := "http://youtube.com/supplier"
	websiteLink := "http://supplierwebsite.com"

	suppliers := []types.Supplier{
		{
			ID:              "1",
			Name:            "Supplier 1",
			LogoImage:       "logo1.jpg",
			MarketingImage:  "marketing1.jpg",
			Description:     "Description 1",
			SocialFacebook:  &facebookLink,
			SocialInstagram: &instagramLink,
			SocialLinkedin:  &linkedinLink,
			SocialTwitter:   &twitterLink,
			SocialYoutube:   &youtubeLink,
			SocialWebsite:   &websiteLink,
			Created:         time.Now(),
		},
		{
			ID:              "2",
			Name:            "Supplier 2",
			LogoImage:       "logo2.jpg",
			MarketingImage:  "marketing2.jpg",
			Description:     "Description 2",
			SocialFacebook:  &facebookLink,
			SocialInstagram: &instagramLink,
			SocialLinkedin:  &linkedinLink,
			SocialTwitter:   &twitterLink,
			SocialYoutube:   &youtubeLink,
			SocialWebsite:   &websiteLink,
			Created:         time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "logo_image", "marketing_image", "description", "social_facebook", "social_instagram", "social_linkedin", "social_twitter", "social_youtube", "social_website", "created"})
	for _, supplier := range suppliers {
		rows.AddRow(supplier.ID, supplier.Name, supplier.LogoImage, supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram, supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube, supplier.SocialWebsite, supplier.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Supplier" ORDER BY created DESC`).WillReturnRows(rows)

	repo := repository.NewSupplierRepository(db)
	retrieved, err := repo.GetSuppliers()
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(suppliers))
		assert.Equal(test, suppliers, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateSupplier(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	facebookLink := "http://facebook.com/supplier"
	instagramLink := "http://instagram.com/supplier"
	linkedinLink := "http://linkedin.com/supplier"
	twitterLink := "http://twitter.com/supplier"
	youtubeLink := "http://youtube.com/supplier"
	websiteLink := "http://supplierwebsite.com"

	supplier := &types.Supplier{
		ID:              "1",
		Name:            "Supplier 1",
		LogoImage:       "logo1.jpg",
		MarketingImage:  "marketing1.jpg",
		Description:     "Description 1",
		SocialFacebook:  &facebookLink,
		SocialInstagram: &instagramLink,
		SocialLinkedin:  &linkedinLink,
		SocialTwitter:   &twitterLink,
		SocialYoutube:   &youtubeLink,
		SocialWebsite:   &websiteLink,
		Created:         time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Supplier" 
	\(id, name, logo_image, marketing_image, description, social_facebook, social_instagram, social_linkedin, social_twitter, social_youtube, social_website, created\) 
	VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12\)`).
		WithArgs(sqlmock.AnyArg(), supplier.Name, supplier.LogoImage,
			supplier.MarketingImage, supplier.Description, supplier.SocialFacebook, supplier.SocialInstagram,
			supplier.SocialLinkedin, supplier.SocialTwitter, supplier.SocialYoutube,
			supplier.SocialWebsite, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewSupplierRepository(db)
	err = repo.CreateSupplier(supplier)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}
