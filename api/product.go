package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Product interface {
	CreateCheckout(cRequest CheckoutRequest) (CheckoutRespose, error)
	ValidateProducts(products []ProductRequest) error
}

type ProductService struct {
	FilePath      string
	productsCache []ProductEntity
}

func NewProductService(filePath string) (ProductService, error) {
	p, err := LoadProducts(filePath)
	return ProductService{FilePath: filePath, productsCache: p}, err
}

func (ps *ProductService) CreateCheckout(cRequest CheckoutRequest) (cResponse CheckoutRespose, err error) {
	err = ps.ValidateProducts(cRequest.Products)
	if err != nil {
		return cResponse, err
	}
	var totalAmount int = 0
	var resp []ProductResponse
	for _, p := range ps.productsCache {
		for i := range cRequest.Products {
			if cRequest.Products[i].ID == p.ID {
				totalItem := p.Amount * cRequest.Products[i].Quantity

				resp = append(resp, ProductResponse{
					ID:          p.ID,
					Quantity:    cRequest.Products[i].Quantity,
					UnitAmount:  p.Amount,
					TotalAmount: totalItem,
					Discount:    0,
					IsGift:      p.IsGift,
				})

				totalAmount += totalItem
			}
		}
	}
	cResponse.TotalAmount = totalAmount
	cResponse.TotalAmountWithDiscount = totalAmount
	cResponse.Products = resp
	return
}

func (ps *ProductService) ValidateProducts(products []ProductRequest) error {
	if len(products) == 0 {
		return fmt.Errorf("empty product list")
	}
	founded := true
	for _, p := range products {
		founded = false
		for i := range ps.productsCache {
			if ps.productsCache[i].ID == p.ID {
				founded = true
				break
			}
		}
		if !founded {
			return fmt.Errorf("invalid product id %v", p.ID)
		}
	}
	return nil
}

func LoadProducts(path string) (products []ProductEntity, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	json.Unmarshal(b, &products)
	return
}
