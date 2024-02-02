package model

// item represents data about a store item.
type item struct {
	Description	string	`json:"shortDescription"`
	Price		float64	`json:"price"`
}

// albums slice to seed item data.
var items = []item{
	{Description: "Mountain Dew", Price: 3.99},
}
