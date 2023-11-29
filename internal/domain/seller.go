package domain

type Seller struct {
	Id              int    `json:"id"`
	Description     string `json:"description"`
	CodSeller       string `json:"cod_seller"`
	IsAuthorization bool   `json:"is_authorization"`
}
