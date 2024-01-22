package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/controllers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func SupplierControllerTest(test *testing.T) (*gin.Engine, *mocks.MockSupplierService, *controllers.SupplierController, *httptest.ResponseRecorder, time.Time) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockSupplierService(ctrl)
	controller := controllers.NewSupplierContoller(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	time := time.Date(2024, time.January, 1, 12, 12, 12, 12, time.UTC)

	return router, mockService, controller, recorder, time
}

func TestGetSuppliers(test *testing.T) {
	router, mockService, controller, recorder, time := SupplierControllerTest(test)

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
			Created:         time,
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
			Created:         time,
		},
	}

	mockService.EXPECT().GetSuppliers().Return(expectedSupplier, nil)

	router.GET("/api/suppliers", controller.GetSuppliers)

	mocks.PerformRequest(test, router, "GET", "/api/suppliers", nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Supplier
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedSupplier, actual)
}

func TestCreateSupplier(test *testing.T) {
	router, mockService, controller, recorder, time := SupplierControllerTest(test)

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
		Created:         time,
	}

	jsonSupplier, _ := json.Marshal(newSupplier)

	expectedResult := &types.SupplierResult{
		PresignedLogoUrl:      "presigned-logo-url",
		LogoUrl:               "logo-url",
		PresignedMarketingUrl: "presigned-marketing-url",
		MarketingUrl:          "marketing-url",
	}

	mockService.EXPECT().CreateSupplier(newSupplier).Return(expectedResult, nil)

	router.POST("/api/suppliers", controller.CreateSupplier)
	mocks.PerformRequest(test, router, "POST", "/api/suppliers", bytes.NewBuffer(jsonSupplier), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actuals types.SupplierResult
	mocks.UnmarshalResponse(test, recorder, &actuals)

	assert.Equal(test, expectedResult, &actuals)
}
