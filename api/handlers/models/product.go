package models

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       string `json:"price"`
	ContactInfo string `json:"contact_info"`
}

type AllProducts struct {
	Products []*Product `json:"products"`
}
