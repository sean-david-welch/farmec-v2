package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/services"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

// type Supplier struct {
//     ID              string    `json:"id"`
//     Name            string    `json:"name"`
//     LogoImage       string    `json:"logo_image"`
//     MarketingImage  string    `json:"marketing_image"`
//     Description     string    `json:"description"`
//     SocialFacebook  *string   `json:"social_facebook"`
//     SocialInstagram *string   `json:"social_instagram"`
//     SocialLinkedin  *string   `json:"social_linkedin"`
//     SocialTwitter   *string   `json:"social_twitter"`
//     SocialYoutube   *string   `json:"social_youtube"`
//     SocialWebsite   *string   `json:"social_website"`
//     Created         time.Time `json:"created"`
// }

func SupplierTestService(test *testing.T) (*mocks.MockSupplierRepository, *mocks.MockS3Client, services.SupplierService) {
	controller := gomock.NewController(test)
	defer controller.Finish()

	mockRepo := mocks.NewMockSupplierRepository(controller)
	mocks3 := mocks.NewMockS3Client(controller)
	service := services.NewSupplierService(mockRepo, mocks3, "test-folder")

	return mockRepo, mocks3, service
}

func TestGetSuppliers(test *testing.T) {
	mockRepo, _, service := SupplierTestService(test)

	socialFacebook := "facebook"
	socialInstagram := "instagram"
	socialLinkedin := "linkedin"
	socialTwitter := "twitter"
	socialYoutube := "youtube"
	socialWebsite := "website"

	expectedSupplier := []types.Supplier{
		{
			ID:              "id ",
			Name:            "name 1",
			LogoImage:       "image.jpg",
			MarketingImage:  "image.jpg",
			Description:     "description",
			SocialFacebook:  &socialFacebook,
			SocialInstagram: &socialInstagram,
			SocialLinkedin:  &socialLinkedin,
			SocialTwitter:   &socialTwitter,
			SocialYoutube:   &socialYoutube,
			SocialWebsite:   &socialWebsite,
		},
		{
			ID:              "id 2",
			Name:            "name 2",
			LogoImage:       "image.jpg",
			MarketingImage:  "image.jpg",
			Description:     "description",
			SocialFacebook:  &socialFacebook,
			SocialInstagram: &socialInstagram,
			SocialLinkedin:  &socialLinkedin,
			SocialTwitter:   &socialTwitter,
			SocialYoutube:   &socialYoutube,
			SocialWebsite:   &socialWebsite,
		},
	}

	mockRepo.EXPECT().GetSuppliers().Return(expectedSupplier, nil)

	suppliers, err := service.GetSuppliers()

	assert.NoError(test, err)
	assert.Equal(test, suppliers, expectedSupplier)
}

func TestCreateSupplier(test *testing.T) {
	mockRepo, mockS3, service := SupplierTestService(test)

	socialFacebook := "facebook"
	socialInstagram := "instagram"
	socialLinkedin := "linkedin"
	socialTwitter := "twitter"
	socialYoutube := "youtube"
	socialWebsite := "website"

	newSupplier := &types.Supplier{
		ID:              "id ",
		Name:            "name 1",
		LogoImage:       "image.jpg",
		MarketingImage:  "image.jpg",
		Description:     "description",
		SocialFacebook:  &socialFacebook,
		SocialInstagram: &socialInstagram,
		SocialLinkedin:  &socialLinkedin,
		SocialTwitter:   &socialTwitter,
		SocialYoutube:   &socialYoutube,
		SocialWebsite:   &socialWebsite,
	}

	mockS3.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presignedurl1", "imageurl1", nil).Times(1)
	mockS3.EXPECT().GeneratePresignedUrl(gomock.Any(), gomock.Any()).Return("presignedurl2", "imageurl2", nil).Times(1)

	mockRepo.EXPECT().CreateSupplier(newSupplier).Return(nil)

	result, err := service.CreateSupplier(newSupplier)

	assert.NoError(test, err)
	assert.NotNil(test, result)

	assert.Equal(test, "presignedurl1", result.PresignedLogoUrl)
	assert.Equal(test, "presignedurl2", result.PresignedMarketingUrl)

	assert.Equal(test, "imageurl1", result.LogoUrl)
	assert.Equal(test, "imageurl2", result.MarketingUrl)
}
