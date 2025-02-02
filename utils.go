package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/RohanIRathi/ReceiptProcessor/database_util"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func calculateRetailerNamePoints(retailer string) {
	count := 0
	for _, c := range retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			count++
		}
	}

	Points_channel <- count
}

func calculateTotalValuePoints(total float32) {
	count := 0
	if float32(int(total)) == total {
		count += 50
	}
	if int(total*100)%25 == 0 {
		count += 25
	}

	Points_channel <- count
}

func calculateDateTimePoints(purchaseDatetime time.Time) {
	count := 0

	start := time.Date(purchaseDatetime.Year(), purchaseDatetime.Month(), purchaseDatetime.Day(), 14, 0, 0, 0, time.UTC)
	end := time.Date(purchaseDatetime.Year(), purchaseDatetime.Month(), purchaseDatetime.Day(), 16, 0, 0, 0, time.UTC)

	if purchaseDatetime.Day()&1 == 1 {
		count += 6
	}
	if purchaseDatetime.After(start) && purchaseDatetime.Before(end) {
		count += 10
	}

	Points_channel <- count
}

func calculateItemPoints(items []database_util.Item) {
	count := 0
	if len(items) > 1 {
		count += 5 * (len(items) / 2)
	}

	for _, item := range items {
		if len(strings.Trim(item.Description, " "))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				Points_channel <- count
				return
			}
			count += int(price*0.2) + 1
		}
	}

	Points_channel <- count
}
