package model

import (
	"fmt"
	"regexp"
	"github.com/go-playground/validator/v10"
)

var (
	validateReceipt *validator.Validate
	retailerRegex = "^[\\w\\s\\-]+$"
	totalRegex = "^\\d+\\.\\d{2}$"
)

func init() {
	validateReceipt = validator.New()
	validateReceipt.RegisterValidation("retailerValidator", validateRetailer)
	validateReceipt.RegisterValidation("totalValidator", validateTotal)
}

// Receipt represents data about a purchase receipt.
type Receipt struct {
    Retailer		string		`json:"retailer" validate:"required,retailerValidator"`
    PurchaseDate	string		`json:"purchaseDate" validate:"required,date"`
    PurchaseTime	string		`json:"purchaseTime" validate:"required,time"`
    items			[]item		`json:"items" validate:"required,min=1,dive"`
    Total			float64		`json:"total" validate:"required,totalValidator"`
}

// Custom validation function for Retailer field
func validateRetailer(fl validator.FieldLevel) bool {
	return regexp.MustCompile(retailerRegex).MatchString(fl.Field().String())
}

// Custom validation function for Total field
func validateTotal(fl validator.FieldLevel) bool {
	return regexp.MustCompile(totalRegex).MatchString(fl.Field().String())
}


// Validate checks if the receipt's fields meet the specified criteria.
func (receipt *Receipt) Validate() error {
	if err := validateReceipt.Struct(receipt); err != nil {
		return fmt.Errorf("validation error for record - '%s': %w", receipt.Retailer, err)
	}
	return nil
}

// Receipts slice to seed receipt data.
var Receipts = []Receipt{}

func ValidateReceipts() error {
	for _, receipt := range Receipts {
		if err := receipt.Validate(); err != nil {
			return err
		}
	}
	return nil
}