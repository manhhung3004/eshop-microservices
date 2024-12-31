package test

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aws-containers/retail-store-sample-app/catalog/model"
)

// CatalogueListHandler trả về danh sách sản phẩm
func CatalogueListHandler(w http.ResponseWriter, r *http.Request) {
	// Dữ liệu mẫu
	products := []model.Product{
		{Name: "Pocket Watch"},
		{Name: "Smartphone"},
		{Name: "Laptop"},
		{Name: "Tablet"},
		{Name: "Headphones"},
		{Name: "Camera"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CatalogueProductHandler trả về thông tin chi tiết một sản phẩm
func CatalogueProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// Dữ liệu mẫu
	if id == "6d62d909-f957-430e-8689-b5129c0bb75e" {
		product := model.Product{Name: "Pocket Watch"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
		return
	}

	// Trả về lỗi nếu không tìm thấy
	http.NotFound(w, r)
}