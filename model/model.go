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
	Description string  `json:"shortDescription" validate:"required,descriptionValidator"`
	Price       float64 `json:"price" validate:"required,gte=0"`
}

// Custom validation function for Description field
func validateDescription(fl validator.FieldLevel) bool {
	return regexp.MustCompile(descriptionRegex).MatchString(fl.Field().String())
}

// Validate checks if the item's fields meet the specified criteria.
func (item *Item) Validate() error {
	return validate.Struct(item)
}

// Items slice to seed item data.
var Items = []Item{
	// {Description: "Mountain Dew", Price: 3.99},
	{Description: "ThisIs-AValidString", Price: 0},
}

func ValidateItems() error {
	for _, item := range Items {
		if err := item.Validate(); err != nil {
			return fmt.Errorf("validation error for record - '%s': %w", item.Description, err)
		}
	}
	return nil
}