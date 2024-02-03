package model

import (
	"fmt"
	"regexp"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	descriptionRegex = "^[\\w\\s\\-]+$"
)

func init() {
	validate = validator.New()
	validate.RegisterValidation("descriptionValidator", validateDescription)
}

// Item represents data about a store item.
type Item struct {
	Description 	string  `json:"shortDescription" validate:"required,descriptionValidator"`
	Price       	float64 `json:"price" validate:"required,gte=0"`
}

// Custom validation function for Description field
func validateDescription(fl validator.FieldLevel) bool {
	return regexp.MustCompile(descriptionRegex).MatchString(fl.Field().String())
}

// Validate checks if the item's fields meet the specified criteria.
func (item *Item) Validate() error {
	if err := validate.Struct(item); err != nil {
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