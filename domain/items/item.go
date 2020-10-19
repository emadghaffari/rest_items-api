package items

// Item struct
type Item struct{
	ID int64 `json:"id"`
	Saller string `json:"saller"`
	Title string `json:"title"`
	Description Description `json:"description"`
	Pictures []Picture `json:"pictures"`
	Video string `json:"video"`
	Price float32 `json:"price"`
	AvailableQuantity int `json:"available_quantity"`
	SoldQuantity int `json:"sold_quantity"`
	Status string `json:"status"`
}
// Description struct
type Description struct{
	PlainText string `json:"plain_text"`
	HTML string `json:"html"`
}

// Picture struct
type Picture struct{
	ID int64 `json:"id"`
	URL int64 `json:"url"`
}