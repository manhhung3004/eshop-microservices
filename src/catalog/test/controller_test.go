package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/aws-containers/retail-store-sample-app/catalog/model"
)

// TestCatalogList kiểm tra danh sách sản phẩm
func TestCatalogList(t *testing.T) {
	// Tạo request giả
	req := httptest.NewRequest("GET", "/catalogue", nil)
	recorder := httptest.NewRecorder()

	// Gọi trực tiếp handler
	CatalogueListHandler(recorder, req)

	// Kiểm tra mã trạng thái trả về
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Giải mã phản hồi và kiểm tra nội dung
	var response []model.Product
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(response)) // Giả định danh sách có 6 sản phẩm
}

// TestCatalogProduct kiểm tra chi tiết một sản phẩm cụ thể
func TestCatalogProduct(t *testing.T) {
	req := httptest.NewRequest("GET", "/catalogue/product/6d62d909-f957-430e-8689-b5129c0bb75e", nil)
	recorder := httptest.NewRecorder()

	// Gọi trực tiếp handler
	CatalogueProductHandler(recorder, req)

	// Kiểm tra mã trạng thái trả về
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Giải mã phản hồi và kiểm tra nội dung
	var response model.Product
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Pocket Watch", response.Name)
}

// TestCatalogProductMissing kiểm tra khi sản phẩm không tồn tại
func TestCatalogProductMissing(t *testing.T) {
	req := httptest.NewRequest("GET", "/catalogue/product/missing", nil)
	recorder := httptest.NewRecorder()

	// Gọi trực tiếp handler
	CatalogueProductHandler(recorder, req)

	// Kiểm tra mã trạng thái trả về
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}