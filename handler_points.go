package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

var Points_channel = make(chan int)

func (apiCfg *ApiConfig) handleGetPoints(w http.ResponseWriter, r *http.Request) {
	receipt_id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respondWithJSON(w, 404, fmt.Sprintf("Receipt not found: %v", err))
		return
	}

	receipt, err := apiCfg.DB.GetReceipt(r.Context(), receipt_id)
	if err != nil {
		respondWithJSON(w, 404, err)
		return
	}

	go calculateRetailerNamePoints(receipt.Retailer)
	total, err := strconv.ParseFloat(receipt.Total, 32)
	if err != nil {
		respondWithJSON(w, 500, "Internal Server Error")
		return
	}
	go calculateTotalValuePoints(float32(total))

	go calculateDateTimePoints(receipt.PurchaseDatetime)

	items, err := apiCfg.DB.GetReceiptItems(r.Context(), receipt_id)
	if err != nil {
		respondWithJSON(w, 500, "Internal Server Error")
		return
	}
	go calculateItemPoints(items)

	count := 0
	count += <-Points_channel
	count += <-Points_channel
	count += <-Points_channel
	count += <-Points_channel

	type Reward struct {
		Points int `json:"points"`
	}

	respondWithJSON(w, 200, Reward{Points: count})
}
