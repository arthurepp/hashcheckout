package main

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {

	verifyCorrectError := func(t *testing.T, result, expected error) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}
	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}
	verifyCorrectProductEntity := func(t *testing.T, result, expected ProductEntity) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%v', expected '%v'", result, expected)
		}
	}

	t.Run("Empty list", func(t *testing.T) {
		productService, err := NewProductService("data/products.json")
		if err != nil {
			panic(err)
		}
		var products []ProductRequest
		result := productService.ValidateProducts(products)
		expected := fmt.Errorf("empty product list")
		verifyCorrectMessage(t, result.Error(), expected.Error())
	})

	t.Run("Invalid product", func(t *testing.T) {
		productService, err := NewProductService("data/products.json")
		if err != nil {
			panic(err)
		}
		var products []ProductRequest
		products = append(products, ProductRequest{
			ID:       1,
			Quantity: 3,
		})
		products = append(products, ProductRequest{
			ID:       127,
			Quantity: 1,
		})
		result := productService.ValidateProducts(products)
		expected := fmt.Errorf("invalid product id %v", 127)
		verifyCorrectMessage(t, result.Error(), expected.Error())
	})

	t.Run("Valid product", func(t *testing.T) {
		productService, err := NewProductService("data/products.json")
		if err != nil {
			panic(err)
		}
		var products []ProductRequest
		products = append(products, ProductRequest{
			ID:       1,
			Quantity: 3,
		})
		products = append(products, ProductRequest{
			ID:       5,
			Quantity: 1,
		})
		result := productService.ValidateProducts(products)
		verifyCorrectError(t, result, nil)
	})

	t.Run("load product correct", func(t *testing.T) {
		result, err := LoadProducts("tests/load_products.json")
		loadedProduct := ProductEntity{
			ID:          129,
			Title:       "Test load",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
			Amount:      250,
			IsGift:      false,
		}
		verifyCorrectProductEntity(t, result[0], loadedProduct)
		verifyCorrectError(t, nil, err)
	})
}
