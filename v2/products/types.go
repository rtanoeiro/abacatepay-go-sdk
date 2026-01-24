package products

import v2 "github.com/AbacatePay/go-types/v2"

type Product = v2.APIProduct

// TODO: Add this in go-types
type DeleteProductData = v2.APIResponse[v2.APIProduct]
type DeleteProductBody struct {
	ID string `json:"id"`
}

type CreateProductBody = v2.RESTPostCreateProductBody

type CreateProductData = v2.RESTPostCreateProductData

type GetProductData = v2.RESTGetProductData

type ListProductsData = v2.RESTGetListProductsData
