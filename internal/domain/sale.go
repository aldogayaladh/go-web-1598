package domain

// Sale.
type Sale struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IdSeller    int    `json:"id_seller"`
	IdProduct   int    `json:"id_product"`
}
