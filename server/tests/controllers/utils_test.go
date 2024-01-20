package controllers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func PerformRequest(test *testing.T, router *gin.Engine, method, path string, body io.Reader, recorder *httptest.ResponseRecorder) {
    req, err := http.NewRequest(method, path, body)
    if err != nil {
        test.Fatal("Error creating request:", err)
    }
    if method == "POST" {
        req.Header.Set("Content-Type", "application/json")
    }
    router.ServeHTTP(recorder, req)
}

func UnmarshalResponse(test *testing.T, recorder *httptest.ResponseRecorder, target interface{}) {
    err := json.Unmarshal(recorder.Body.Bytes(), target)
    if err != nil {
        test.Fatal("Error unmarshalling response:", err)
    }
}