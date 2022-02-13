package main

type Price interface {
	CalculateDiscount(cResponse CheckoutRespose)
}
