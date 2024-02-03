package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemValidation(t *testing.T) {
	t.Run("Valid Item", func(t *testing.T) {
		validItem := Item{
			Description: "Valid Description",
			Price:       "10.99",
		}

		err := validItem.Validate()
		assert.NoError(t, err, "Validation should pass for a valid item")
	})

	t.Run("Valid Description - Underscore", func(t *testing.T) {
		validItem := Item{
			Description: "Valid_Description",
			Price:       "2.99",
		}

		err := validItem.Validate()
		assert.NoError(t, err, "Validation should pass for a valid item")
	})

	t.Run("Valid Description - Numbers", func(t *testing.T) {
		validItem := Item{
			Description: "Val1d D3scr1pt10n",
			Price:       "6.00",
		}

		err := validItem.Validate()
		assert.NoError(t, err, "Validation should pass for a valid item")
	})

	t.Run("Invalid Description - Symbols", func(t *testing.T) {
		invalidItem := Item{
			Description: "Invalid!@#$Description",
			Price:       "10.99",
		}

		err := invalidItem.Validate()
		assert.Error(t, err, "Validation should fail for an invalid description")
		assert.Contains(t, err.Error(), "descriptionValidator", "Error message should mention the failed validator")
	})

	t.Run("Valid Price - Zeroes", func(t *testing.T) {
		validItem := Item{
			Description: "Val1d D3scr1pt10n",
			Price:       "0.00",
		}

		err := validItem.Validate()
		assert.NoError(t, err, "Validation should pass for a valid item")
	})

	t.Run("Invalid Price - Integer as string", func(t *testing.T) {
		invalidItem := Item{
			Description: "Valid Description",
			Price:       "5",
		}

		err := invalidItem.Validate()
		assert.Error(t, err, "Validation should fail for an invalid price")
		assert.Contains(t, err.Error(), "priceValidator", "Error message should mention the failed validator")
	})

	t.Run("Invalid Price - Single Decimal", func(t *testing.T) {
		invalidItem := Item{
			Description: "Valid Description",
			Price:       "5.0",
		}

		err := invalidItem.Validate()
		assert.Error(t, err, "Validation should fail for an invalid price")
		assert.Contains(t, err.Error(), "priceValidator", "Error message should mention the failed validator")
	})

	t.Run("Invalid Price - Negative Number", func(t *testing.T) {
		invalidItem := Item{
			Description: "Valid Description",
			Price:       "-5.00",
		}

		err := invalidItem.Validate()
		assert.Error(t, err, "Validation should fail for an invalid price")
		assert.Contains(t, err.Error(), "priceValidator", "Error message should mention the failed validator")
	})

	t.Run("Invalid Price - Integer as int", func(t *testing.T) {
		invalidJSON := `{
			"shortDescription": "Valid Description",
			"price": 5
		}`
	
		var newItem Item
	
		err := json.Unmarshal([]byte(invalidJSON), &newItem)
	
		assert.Error(t, err, "Binding should fail for an invalid price")
		assert.Contains(t, err.Error(), "json", "Error message should mention JSON decoding error")
	})

	t.Run("Invalid Price - Decimal as float", func(t *testing.T) {
		invalidJSON := `{
			"shortDescription": "Valid Description",
			"price": 5.99
		}`
	
		var newItem Item
	
		err := json.Unmarshal([]byte(invalidJSON), &newItem)
	
		assert.Error(t, err, "Binding should fail for an invalid price")
		assert.Contains(t, err.Error(), "json", "Error message should mention JSON decoding error")
	})
}

func TestValidateItems(t *testing.T) {
	t.Run("Empty Items Slice", func(t *testing.T) {
		// Ensure that ValidateItems works correctly with an empty slice
		err := ValidateItems()
		assert.NoError(t, err, "ValidateItems should pass with an empty items slice")
	})

	t.Run("Valid Items", func(t *testing.T) {
		// Add some valid items to the Items slice
		validItems := []Item{
			{Description: "Item1", Price: "10.99"},
			{Description: "Item2", Price: "5.00"},
		}

		Items = append(Items, validItems...)

		err := ValidateItems()
		assert.NoError(t, err, "ValidateItems should pass with valid items")
	})

	t.Run("Invalid Items", func(t *testing.T) {
		// Add some invalid items to the Items slice
		invalidItems := []Item{
			{Description: "Invalid-Item!@#", Price: "5.00"},
			{Description: "Item3", Price: "8.99"},
		}

		Items = append(Items, invalidItems...)

		err := ValidateItems()
		assert.Error(t, err, "ValidateItems should fail with invalid items")
		assert.Contains(t, err.Error(), "validation error", "Error message should indicate a validation error")
	})
}
