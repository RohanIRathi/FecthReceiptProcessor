package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/RohanIRathi/ReceiptProcessor/database_util"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) handleCreateReceipt(w http.ResponseWriter, r *http.Request) {
	type item struct {
		Description string `json:"shortDescription"`
		Price       string `json:"price"`
	}

	type parameters struct {
		Retailer     string `json:"retailer"`
		PurchaseDate string `json:"purchaseDate"`
		PurchaseTime string `json:"purchaseTime"`
		Items        []item `json:"items"`
		Total        string `json:"total"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("%v /receipts/process - %v: %v", r.Method, http.StatusBadRequest, err)
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %v\n", err))
		return
	}
	purchaseTimestamp, err := time.Parse(time.RFC3339, params.PurchaseDate+"T"+params.PurchaseTime+":00Z")
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %v\n", err))
		return
	}
	total, err := strconv.ParseFloat(params.Total, 64)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %v\n", err))
		return
	}

	receipt, err := apiCfg.DB.CreateReceipt(r.Context(), database_util.CreateReceiptParams{
		ID:               uuid.New().String(),
		Retailer:         params.Retailer,
		PurchaseDatetime: purchaseTimestamp,
		Total:            total,
	})
	if err != nil {
		log.Printf("%v /receipts/process - %v: %v", r.Method, http.StatusInternalServerError, err)
		respondWithJSON(w, http.StatusInternalServerError, "Error adding Receipt!")
		return
	}
	receipt_id := receipt.ID
	for _, i := range params.Items {

		price, err := strconv.ParseFloat(i.Price, 64)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %v\n", err))
			return
		}
		_, err = apiCfg.DB.AddItem(r.Context(), database_util.AddItemParams{
			ID:          uuid.New().String(),
			Description: i.Description,
			Price:       price,
			Receipt:     receipt_id,
		})
		if err != nil {
			log.Printf("%v /receipts/process - %v: %v", r.Method, http.StatusInternalServerError, err)
			respondWithJSON(w, http.StatusInternalServerError, "Error adding items")
			return
		}
	}

	type Success struct {
		ID string `json:"id"`
	}

	respondWithJSON(w, 200, Success{ID: receipt_id})
	log.Printf("%v /receipts/process - %v", r.Method, http.StatusOK)
}
