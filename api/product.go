package main

type Product interface {
	CreateCheckout(cRequest CheckoutRequest) (CheckoutRespose, error)
	ValidateProducts(products []ProductRequest) error
}
