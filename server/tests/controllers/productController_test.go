package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/sean-david-welch/farmec-v2/server/handlers"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func ProductControllerTest(test *testing.T) (*gin.Engine, *mocks.MockProductService, *handlers.ProductController, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	mockService := mocks.NewMockProductService(ctrl)
	controller := handlers.NewProductController(mockService)

	router := gin.Default()
	recorder := httptest.NewRecorder()

	return router, mockService, controller, recorder
}

func TestGetProducts(test *testing.T) {
	router, mockService, controller, recorder := ProductControllerTest(test)

	expectedProducts := []types.Product{
		{
			ID:           "1",
			MachineID:    "12",
			Name:         "Product 1",
			ProductImage: "image1.jpg",
			Description:  "description 1",
			ProductLink:  "productLink 1",
		},
		{
			ID:           "2",
			MachineID:    "12",
			Name:         "Product 2",
			ProductImage: "image1.jpg",
			Description:  "description 2",
			ProductLink:  "productLink 2",
		},
	}

	machine_id := expectedProducts[0].MachineID

	mockService.EXPECT().GetProducts(machine_id).Return(expectedProducts, nil)

	router.GET("/api/products/:id", controller.GetProducts)

	url := fmt.Sprintf("/api/products/%s", machine_id)

	mocks.PerformRequest(test, router, "GET", url, nil, recorder)

	assert.Equal(test, http.StatusOK, recorder.Code)

	var actual []types.Product
	mocks.UnmarshalResponse(test, recorder, &actual)

	assert.Equal(test, expectedProducts, actual)
}

func TestCreateProduct(test *testing.T) {
	router, mockService, controller, recorder := ProductControllerTest(test)

	newProduct := &types.Product{
		ID:           "1",
		MachineID:    "12",
		Name:         "Product 1",
		ProductImage: "image1.jpg",
		Description:  "description 1",
		ProductLink:  "productLink 1",
	}

	jsonProduct, _ := json.Marshal(newProduct)

	expectedResult := &types.ModelResult{PresignedUrl: "presigned-url", ImageUrl: "image-url"}

	mockService.EXPECT().CreateProduct(newProduct).Return(expectedResult, nil)

	router.POST("/api/products", controller.CreateProduct)
	mocks.PerformRequest(test, router, "POST", "/api/products", bytes.NewBuffer(jsonProduct), recorder)

	assert.Equal(test, http.StatusCreated, recorder.Code)

	var actuals types.ModelResult
	mocks.UnmarshalResponse(test, recorder, &actuals)

	assert.Equal(test, expectedResult, &actuals)
}
