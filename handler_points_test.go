package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RohanIRathi/ReceiptProcessor/database_util"
	"github.com/google/uuid"
)

func TestHandlerGetPoints(t *testing.T) {
	type Response struct {
		Points int `json:"points"`
	}

	req := httptest.NewRequest(http.MethodGet, "/receipts/e1ce4eb1-a248-4b18-8e66-efc4606b9447/points", nil)
	w := httptest.NewRecorder()
	id1 := uuid.MustParse("e1ce4eb1-a248-4b18-8e66-efc4606b9447")
	receipt1 := database_util.CreateReceiptParams{
		ID:               id1,
		Retailer:         "Target",
		PurchaseDatetime: time.Date(2022, 1, 1, 13, 1, 0, 0, time.UTC),
		Total:            "35.35",
	}
	apiCfg.DB.CreateReceipt(req.Context(), receipt1)

	items1 := []database_util.AddItemParams{
		{ID: uuid.New(), Description: "Mountain Dew 12PK", Price: "6.49", Receipt: id1},
		{ID: uuid.New(), Description: "Emils Cheese Pizza", Price: "12.25", Receipt: id1},
		{ID: uuid.New(), Description: "Knorr Creamy Chicken", Price: "1.26", Receipt: id1},
		{ID: uuid.New(), Description: "Doritos Nacho Cheese", Price: "3.35", Receipt: id1},
		{ID: uuid.New(), Description: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00", Receipt: id1},
	}
	for _, item := range items1 {
		apiCfg.DB.AddItem(req.Context(), item)
	}

	req.SetPathValue("id", id1.String())
	apiCfg.handleGetPoints(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	response := Response{}
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil || response.Points != 28 {
		t.Fail()
	}

	res.Body.Close()

	req = httptest.NewRequest(http.MethodGet, "/receipts/90e3ad56-b8ff-48aa-a4d5-8dd3654110f4/points", nil)
	w = httptest.NewRecorder()

	id2 := uuid.MustParse("90e3ad56-b8ff-48aa-a4d5-8dd3654110f4")
	receipt2 := database_util.CreateReceiptParams{
		ID:               id2,
		Retailer:         "M&M Corner Market",
		PurchaseDatetime: time.Date(2022, 3, 20, 14, 33, 0, 0, time.UTC),
		Total:            "9.00",
	}
	apiCfg.DB.CreateReceipt(req.Context(), receipt2)
	items2 := []database_util.AddItemParams{
		{ID: uuid.New(), Description: "Gatorade", Price: "2.25", Receipt: id2},
		{ID: uuid.New(), Description: "Gatorade", Price: "2.25", Receipt: id2},
		{ID: uuid.New(), Description: "Gatorade", Price: "2.25", Receipt: id2},
		{ID: uuid.New(), Description: "Gatorade", Price: "2.25", Receipt: id2},
	}
	for _, item := range items2 {
		apiCfg.DB.AddItem(req.Context(), item)
	}

	req.SetPathValue("id", id2.String())
	apiCfg.handleGetPoints(w, req)

	res = w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil || response.Points != 109 {
		t.Fail()
	}

	res.Body.Close()

	req = httptest.NewRequest(http.MethodGet, "/receipts/90e3ad56-b8ff-48aa-a4d5-8dd3654110f4/points", nil)
	w = httptest.NewRecorder()

	apiCfg.handleGetPoints(w, req)

	res = w.Result()

	if res.StatusCode != http.StatusNotFound {
		t.Fail()
	}

	req = httptest.NewRequest(http.MethodGet, "/receipts/90e3ad56-b8ff-48aa-a4d5-8dd3654110f5/points", nil)
	req.SetPathValue("id", "90e3ad56-b8ff-48aa-a4d5-8dd3654110f5")
	w = httptest.NewRecorder()

	apiCfg.handleGetPoints(w, req)

	res = w.Result()

	if res.StatusCode != http.StatusNotFound {
		t.Fail()
	}
}
