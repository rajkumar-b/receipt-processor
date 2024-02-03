package model

import (
	"fmt"
	"regexp"
	"github.com/go-playground/validator/v10"
)

var (
	validateItem *validator.Validate
	descriptionRegex = "^[\\w\\s\\-]+$"
	priceRegex = "^\\d+\\.\\d{2}$"
)

func init() {
	validateItem = validator.New()
	validateItem.RegisterValidation("descriptionValidator", validateDescription)
	validateItem.RegisterValidation("priceValidator", validatePrice)
}

// Item represents data about a store item.
type Item struct {
	Description 	string	`json:"shortDescription" validate:"required,descriptionValidator"`
	Price       	string	`json:"price" validate:"required,priceValidator"`
}

// Custom validation function for Description field
func validateDescription(fl validator.FieldLevel) bool {
	return regexp.MustCompile(descriptionRegex).MatchString(fl.Field().String())
}

// Custom validation function for Price field
func validatePrice(fl validator.FieldLevel) bool {
	return regexp.MustCompile(priceRegex).MatchString(fl.Field().String())
}

// Validate checks if the item's fields meet the specified criteria.
func (item *Item) Validate() error {
	if err := validateItem.Struct(item); err != nil {
		return fmt.Errorf("validation error for record - '%s': %w", item.Description, err)
	}
	return nil
}

// Items slice to seed item data.
var Items = []Item{}

func ValidateItems() error {
	for _, item := range Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}