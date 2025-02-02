package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCreateReceipt(t *testing.T) {
	type Response struct {
		ID string `json:"id"`
	}
	body := strings.NewReader(`{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [],
		"total": "$35.35"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/receipts/process", body)
	w := httptest.NewRecorder()

	apiCfg.handleCreateReceipt(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
	body = strings.NewReader(`{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
				"shortDescription": "Mountain Dew 12PK",
				"price": "$6.49"
			},{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			},{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			},{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			},{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			}
		],
		"total": "35.35"
	}`)

	req = httptest.NewRequest(http.MethodPost, "/receipts/process", body)
	w = httptest.NewRecorder()

	apiCfg.handleCreateReceipt(w, req)

	res = w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}

	body1 := strings.NewReader(`{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			},{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			},{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			},{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			},{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
				}
			],
		"total": "35.35"
	}`)

	req = httptest.NewRequest(http.MethodPost, "/receipts/process", body1)
	w = httptest.NewRecorder()

	apiCfg.handleCreateReceipt(w, req)

	res = w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	var receipt Response

	err := json.NewDecoder(res.Body).Decode(&receipt)
	if err != nil || len(receipt.ID) != 36 {
		t.Fail()
	}
	res.Body.Close()

	body2 := strings.NewReader(``)

	req = httptest.NewRequest(http.MethodPost, "/receipts/process", body2)
	w = httptest.NewRecorder()

	apiCfg.handleCreateReceipt(w, req)

	res = w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}

	body3 := strings.NewReader(`{
		"retailer": "Target",
		"purchaseDate": "202-01-01",
		"purchaseTime": "13:01",
		"items": [
			{
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			},{
					"shortDescription": "Emils Cheese Pizza",
					"price": "12.25"
			},{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			},{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			},{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			}
		],
		"total": "35.35"
	}`)

	req = httptest.NewRequest(http.MethodPost, "/receipts/process", body3)
	w = httptest.NewRecorder()

	apiCfg.handleCreateReceipt(w, req)

	res = w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}
