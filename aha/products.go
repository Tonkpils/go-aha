package aha

import "encoding/json"

type ProductsService struct {
	client *Client
}

type Product struct {
	ID              string `json:"id"`
	ReferencePrefix string `json:"reference_prefix"`
	Name            string `json:"name"`
	ProductLine     bool   `json:"product_line"`
	CreatedAt       string `json:"created_at"` // TODO: Turn this into Time?
}

type ProductsResponse struct {
	Products   []*Product     `json:"products"`
	Pagination map[string]int `json:"pagination"`
}

func (s *ProductsService) ListAll() ([]*Product, error) {
	req, err := s.client.NewRequest("GET", "products", nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	productsList := new(ProductsResponse)

	if err := json.NewDecoder(resp.Body).Decode(productsList); err != nil {
		return nil, err
	}

	return productsList.Products, nil
}
