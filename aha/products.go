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

func (s *ProductsService) ListAll() ([]Product, *Pagination, error) {
	req, err := s.client.NewRequest("GET", "products", nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	productsList := new(struct {
		Products   []Product   `json:"products"`
		Pagination *Pagination `json:"pagination"`
	})

	if err := json.NewDecoder(resp.Body).Decode(productsList); err != nil {
		return nil, nil, err
	}

	return productsList.Products, productsList.Pagination, nil
}
