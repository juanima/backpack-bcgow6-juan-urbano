package domain

type Product struct {
        Id      int     `json:"id"`
	Stock   int     `json:"stock"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Code    string  `json:"code"`
	Price   float64 `json:"price"`
	Publish *bool   `json:"publish"`
	Date    string  `json:"date"`
}

