package models

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type PhonePrice struct {
	FinalPrice sql.NullFloat64 `json:"final_price"`
}

func CalculateFinalPrice(db *gorm.DB, phoneID int, choiceID []int) (float64, error) {
	var finalPrice PhonePrice
	idListStr := "ARRAY["
	for i, id := range choiceID {
		if i > 0 {
			idListStr += ","
		}
		idListStr += fmt.Sprintf("%d", id)
	}
	idListStr += "]"
	query := fmt.Sprintf("select * from calculate_final_price(%d,%s)", phoneID, idListStr)

	row := db.Raw(query).Row()
	if err := row.Scan(&finalPrice.FinalPrice); err != nil {
		return 0, fmt.Errorf("error while calculating final price: %v", err)
	}
	return finalPrice.FinalPrice.Float64, nil
}
