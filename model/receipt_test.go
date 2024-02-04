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
	
	t.Run("Valid Receipt - multiple items with one invalid item", func(t *testing.T) {
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
