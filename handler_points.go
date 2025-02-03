package main

import (
	"fmt"
	"log"
	"net/http"
)

var Points_channel = make(chan int)

func (apiCfg *ApiConfig) handleGetPoints(w http.ResponseWriter, r *http.Request) {
	receipt_id := r.PathValue("id")
	if receipt_id == "" {
		log.Printf("%v %v - %v: id not provided", r.Method, r.URL, http.StatusNotFound)
		respondWithJSON(w, http.StatusNotFound, fmt.Sprintf("Receipt not found: id not found"))
		return
	}

	receipt, err := apiCfg.DB.GetReceipt(r.Context(), receipt_id)
	if err != nil {
		log.Printf("%v %v - %v: %v", r.Method, r.URL, http.StatusNotFound, err)
		respondWithJSON(w, http.StatusNotFound, fmt.Sprintf("Recipt not found: %v", err))
		return
	}

	go calculateRetailerNamePoints(receipt.Retailer)
	total := receipt.Total
	go calculateTotalValuePoints(total)

	go calculateDateTimePoints(receipt.PurchaseDatetime)

	items, err := apiCfg.DB.GetReceiptItems(r.Context(), receipt_id)
	if err != nil {
		log.Printf("%v %v - %v: %v", r.Method, r.URL, http.StatusInternalServerError, err)
		respondWithJSON(w, http.StatusInternalServerError, "Internal Server Error")
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
	log.Printf("%v %v - %v", r.Method, r.URL, http.StatusOK)
}
