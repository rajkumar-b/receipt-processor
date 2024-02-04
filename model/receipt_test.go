package model

import (
	// "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptValidation(t *testing.T) {

	t.Run("Valid Receipt - single item", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "13:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "2.99",
				},
			},
			Total: "2.99",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})

	t.Run("Valid Receipt - check ID", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "13:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "2.99",
				},
			},
			Total: "2.99",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
		
		err_id := validReceipt.Validate()
		assert.NoError(t, err_id, "Validation should pass for a valid receipt")
	})

	t.Run("Valid Receipt - With ID", func(t *testing.T) {
		validReceipt := Receipt{
			ID:				"some-id",
			Retailer:		"Valid Retailer",
			PurchaseDate:	"2022-09-20",
			PurchaseTime:	"13:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "2.99",
				},
			},
			Total: "2.99",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})

	t.Run("Valid Receipt - multiple items", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "13:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "2.99",
				},
				{
					Description: "Valid Item 2",
					Price:       "1.99",
				},
			},
			Total: "4.98",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - normal time(shortened hr)", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "2.99",
				},
			},
			Total: "4.98",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - numerical retailer", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid123Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "2.99",
				},
			},
			Total: "2.99",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - underscore in retailer", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "2.99",
				},
			},
			Total: "2.99",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - zero total", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "0.00",
				},
			},
			Total: "0.00",
		}

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - zero points", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "0.00",
				},
			},
			Total: "0.00",
		}

		validReceipt.SetPoints(0)
		assert.Equal(t, 0, validReceipt.Points, "Points should be set to 0")

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - negative points", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "0.00",
				},
			},
			Total: "0.00",
		}

		validReceipt.SetPoints(-13)
		assert.Equal(t, -13, validReceipt.Points, "Points should be set to -13")

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Valid Receipt - positive points", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "0.00",
				},
			},
			Total: "0.00",
		}

		validReceipt.SetPoints(10)
		assert.Equal(t, 10, validReceipt.Points, "Points should be set to 10")

		err := validReceipt.Validate()
		assert.NoError(t, err, "Validation should pass for a valid receipt")
	})
	
	t.Run("Invalid Receipt - Zero Items", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
			},
			Total: "0.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid item count")
		assert.Contains(t, err.Error(), "min", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - Invalid Item (Description)", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Invalid Item @#$",
					Price:       "4.00",
				},
			},
			Total: "4.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid item description")
		assert.Contains(t, err.Error(), "descriptionValidator", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - Invalid ID", func(t *testing.T) {
		validReceipt := Receipt{
			ID:				"invalid id with space",
			Retailer:		"Valid_Retailer",
			PurchaseDate:	"2022-09-20",
			PurchaseTime:	"3:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "4.00",
				},
			},
			Total: "4.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid ID value")
		assert.Contains(t, err.Error(), "idValidator", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - Invalid Item (Price)", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:01",
			PurchasedItems: []Item{
				{
					Description: "Invalid Item",
					Price:       "-3.00",
				},
			},
			Total: "0.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid item price")
		assert.Contains(t, err.Error(), "priceValidator", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - negative total", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "23:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "-3.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid total")
		assert.Contains(t, err.Error(), "totalValidator", "Error message should mention the failed validator")
	})

	t.Run("Invalid Receipt - integer(as string) total", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "23:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "3",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid total")
		assert.Contains(t, err.Error(), "totalValidator", "Error message should mention the failed validator")
	})

	t.Run("Invalid Receipt - integer(as string) total", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "23:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "3",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid total")
		assert.Contains(t, err.Error(), "totalValidator", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - shortened minute", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "3:1",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "3.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid time")
		assert.Contains(t, err.Error(), "datetime", "Error message should mention the failed validator")
	})

	t.Run("Invalid Receipt - shortened month", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-9-20",
			PurchaseTime:  "03:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "3.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid date")
		assert.Contains(t, err.Error(), "datetime", "Error message should mention the failed validator")
	})

	t.Run("Invalid Receipt - shortened date", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid_Retailer",
			PurchaseDate:  "2022-09-2",
			PurchaseTime:  "03:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item",
					Price:       "3.00",
				},
			},
			Total: "3.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid date")
		assert.Contains(t, err.Error(), "datetime", "Error message should mention the failed validator")
	})
	
	t.Run("Invalid Receipt - multiple items with one invalid item", func(t *testing.T) {
		validReceipt := Receipt{
			Retailer:      "Valid Retailer",
			PurchaseDate:  "2022-09-20",
			PurchaseTime:  "13:01",
			PurchasedItems: []Item{
				{
					Description: "Valid Item 1",
					Price:       "2.99",
				},
				{
					Description: "Invalid Item 2",
					Price:       "-1.99",
				},
			},
			Total: "1.00",
		}

		err := validReceipt.Validate()
		assert.Error(t, err, "Validation should fail for an invalid item price")
		assert.Contains(t, err.Error(), "priceValidator", "Error message should mention the failed validator")
	})

}

func TestValidateReceipts(t *testing.T) {
	t.Run("Empty Items Slice", func(t *testing.T) {
		// Ensure that ValidateItems works correctly with an empty slice
		err := ValidateReceipts()
		assert.NoError(t, err, "ValidateReceipts should pass with an empty items slice")
	})

	t.Run("Valid Receipt", func(t *testing.T) {
		// Add some valid receipts to the Items slice
		validReceipts := []Receipt{
			{
				Retailer:      "Valid Retailer 1",
				PurchaseDate:  "2022-09-20",
				PurchaseTime:  "13:01",
				PurchasedItems: []Item{
					{
						Description: "Valid Item 1",
						Price:       "2.99",
					},
					{
						Description: "Valid Item 2",
						Price:       "1.99",
					},
				},
				Total: "4.98",
			},
			{
				Retailer:      "Valid Retailer",
				PurchaseDate:  "2022-09-20",
				PurchaseTime:  "13:01",
				PurchasedItems: []Item{
					{
						Description: "Valid Item 1",
						Price:       "2.99",
					},
					{
						Description: "Valid Item 2",
						Price:       "1.99",
					},
				},
				Total: "4.98",
			},
		}

		Receipts = append(Receipts, validReceipts...)

		err := ValidateReceipts()
		assert.NoError(t, err, "ValidateReceipts should pass with valid receipts")
	})
	
	t.Run("Invalid Receipt", func(t *testing.T) {
		// Add some valid receipts to the Items slice
		validReceipts := []Receipt{
			{
				Retailer:      "Valid Retailer 1",
				PurchaseDate:  "2022-09-20",
				PurchaseTime:  "13:01",
				PurchasedItems: []Item{
					{
						Description: "Valid Item 1",
						Price:       "2.99",
					},
					{
						Description: "Valid Item 2",
						Price:       "1.99",
					},
				},
				Total: "4.98",
			},
			{
				Retailer:      "Valid Retailer",
				PurchaseDate:  "2022-09-20",
				PurchaseTime:  "13:01",
				PurchasedItems: []Item{
					{
						Description: "Invalid Item 1",
						Price:       "-2.99",
					},
					{
						Description: "Valid Item 2",
						Price:       "1.99",
					},
				},
				Total: "4.98",
			},
		}

		Receipts = append(Receipts, validReceipts...)

		err := ValidateReceipts()
		assert.Error(t, err, "ValidateReceipts should fail with invalid receipts")
		assert.Contains(t, err.Error(), "validation error", "Error message should indicate a validation error")
	})
}
