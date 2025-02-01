package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RohanIRathi/ReceiptProcessor/database_util"
	"github.com/google/uuid"
)

type item struct {
	Description string `json:"shortDescription"`
	Price       string `json:"price"`
}

func (apiCfg *ApiConfig) handleCreateReceipt(w http.ResponseWriter, r *http.Request) {
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
		respondWithJSON(w, 400, fmt.Sprintf("Error parsing json: %v\n", err))
		return
	}
	purchaseTimestamp, err := time.Parse(time.RFC3339, params.PurchaseDate+"T"+params.PurchaseTime+":00Z")
	if err != nil {
		respondWithJSON(w, 400, fmt.Sprintf("Error parsing json: %v\n", err))
		return
	}

	receipt, err := apiCfg.DB.CreateReceipt(r.Context(), database_util.CreateReceiptParams{
		ID:               uuid.New(),
		Retailer:         params.Retailer,
		PurchaseDatetime: purchaseTimestamp,
		Total:            params.Total,
	})
	if err != nil {
		log.Printf("Internal Error adding receipt: %v", err)
		respondWithJSON(w, 500, "Error adding Receipt!")
		return
	}
	receipt_id := receipt.ID
	for _, i := range params.Items {
		_, err := apiCfg.DB.AddItem(r.Context(), database_util.AddItemParams{
			ID:          uuid.New(),
			Description: i.Description,
			Price:       i.Price,
			Receipt:     receipt_id,
		})
		if err != nil {
			log.Printf("Error adding items: %v", err)
			respondWithJSON(w, 500, "Error adding items")
			return
		}
	}

	type Success struct {
		ID uuid.UUID `json:"id"`
	}

	respondWithJSON(w, 200, Success{ID: receipt_id})
}
