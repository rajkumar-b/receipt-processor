package model

// Item represents data about a store item.
type Item struct {
	Description string  `json:"shortDescription"`
	Price       float64 `json:"price"`
}

// Items slice to seed item data.
var Items = []Item{
	{Description: "Mountain Dew", Price: 3.99},
}
