package aha

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductsService_ListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"products": [{
			"id": "131414752",
			"reference_prefix": "PRJ1",
			"name": "Project 1",
			"product_line": false,
			"created_at": "2016-07-04T19:18:49Z"
		}], "pagination": {
			"total_records": 1, "total_pages": 1, "current_page": 1
		}}`)
	})

	products, pagination, err := client.Products.ListAll()
	require.NoError(t, err)

	expected := []Product{{
		ID:              "131414752",
		ReferencePrefix: "PRJ1",
		Name:            "Project 1",
		ProductLine:     false,
		CreatedAt:       "2016-07-04T19:18:49Z",
	}}
	assert.Equal(t, expected, products)

	paginationExpected := &Pagination{TotalRecords: 1, TotalPages: 1, CurrentPage: 1}
	assert.Equal(t, paginationExpected, pagination)

}
