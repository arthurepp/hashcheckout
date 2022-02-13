package main

import (
	"testing"
	"time"
)

func TestIsBlackFriday(t *testing.T) {
	var products []ProductResponse
	products = append(products, ProductResponse{
		ID:          1,
		Quantity:    3,
		UnitAmount:  200,
		TotalAmount: 600,
		Discount:    0,
		IsGift:      false,
	})
	products = append(products, ProductResponse{
		ID:          2,
		Quantity:    2,
		UnitAmount:  100,
		TotalAmount: 200,
		Discount:    0,
		IsGift:      false,
	})
	checkout := CheckoutRespose{
		TotalAmount:             800,
		TotalAmountWithDiscount: 800,
		TotalDiscount:           0,
		Products:                products,
	}

	date, err := time.Parse("2006-01-02", time.Now().Truncate(time.Hour*24).Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	blackFridayService := NewBlackFridayService(date)
	blackFridayService.AddGiftOnBlackFriday(&checkout)

	if len(checkout.Products) < 3 {
		t.Errorf("invalid gift")
	}
}

func TestNotIsBlackFriday(t *testing.T) {
	var products []ProductResponse
	products = append(products, ProductResponse{
		ID:          1,
		Quantity:    3,
		UnitAmount:  200,
		TotalAmount: 600,
		Discount:    0,
		IsGift:      false,
	})
	products = append(products, ProductResponse{
		ID:          2,
		Quantity:    2,
		UnitAmount:  100,
		TotalAmount: 200,
		Discount:    0,
		IsGift:      false,
	})
	checkout := CheckoutRespose{
		TotalAmount:             800,
		TotalAmountWithDiscount: 800,
		TotalDiscount:           0,
		Products:                products,
	}

	date, err := time.Parse("2006-01-02", "2022-02-03")
	if err != nil {
		panic(err)
	}
	blackFridayService := NewBlackFridayService(date)
	blackFridayService.AddGiftOnBlackFriday(&checkout)

	if len(checkout.Products) > 2 {
		t.Errorf("invalid gift")
	}
}
