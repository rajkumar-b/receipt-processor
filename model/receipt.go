package model

import (
	"fmt"
	"regexp"
	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

var (
	retailerRegex = "^[\\w\\s\\-]+$"
	totalRegex = "^\\d+\\.\\d{2}$"
	idRegex = "^\\S+$"
)

func init() {
	registerReceiptValidator()
}

func registerReceiptValidator() {
	validate.RegisterValidation("retailerValidator", validateRetailer)
	validate.RegisterValidation("totalValidator", validateTotal)
	validate.RegisterValidation("idValidator", validateID)
}

// Receipt represents data about a purchase receipt.
type Receipt struct {
	ID				string		`json:"id" validate:"idValidator"`
    Retailer		string		`json:"retailer" validate:"required,retailerValidator"`
    PurchaseDate	string		`json:"purchaseDate" validate:"required,datetime=2006-01-02"`
    PurchaseTime	string		`json:"purchaseTime" validate:"required,datetime=15:04"`
    PurchasedItems	[]Item		`json:"items" validate:"required,min=1,dive"`
    Total			string		`json:"total" validate:"required,totalValidator"`
	Points			int			`json:"points"`
}

// Custom validation function for Retailer field
func validateRetailer(fl validator.FieldLevel) bool {
	return regexp.MustCompile(retailerRegex).MatchString(fl.Field().String())
}

// Custom validation function for Total field
func validateTotal(fl validator.FieldLevel) bool {
	return regexp.MustCompile(totalRegex).MatchString(fl.Field().String())
}

// Custom validation function for ID field
func validateID(fl validator.FieldLevel) bool {
	id := fl.Field().String()
	return len(id) == 0 || regexp.MustCompile(idRegex).MatchString(id)
}

// Validate checks if the receipt's fields meet the specified criteria and generates ID if validated but not possess an ID.
func (receipt *Receipt) Validate() error {
	if err := validate.Struct(receipt); err != nil {
		return fmt.Errorf("validation error for record - '%s': %w", receipt.Retailer, err)
	}
	receipt.generateID()
	return nil
}

// generateID generates and sets ID only if it is not present	
func (receipt *Receipt) generateID() {
	if receipt.ID == "" {
		receipt.ID = uuid.New().String()
	}
}

// SetPoints sets the passed Points to the Receipt.	
func (receipt *Receipt) SetPoints(points int) {
    receipt.Points = points
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

// GetReceiptByID retrieves a receipt by its ID.
func GetReceiptByID(id string) (*Receipt, error) {
	if id == "" {
		return nil, fmt.Errorf("No ID passed")
	}

    for _, receipt := range Receipts {
        if receipt.ID == id {
            return &receipt, nil
        }
    }
    return nil, fmt.Errorf("receipt not found")
}