package main

import (
	"context"
	"testing"

	discount "github.com/arthurepp/hashcheckout/api/grpc"
	"google.golang.org/grpc"
)

type MockGrpcDiscountClient struct {
	discount.DiscountClient
}

func (c *MockGrpcDiscountClient) GetDiscount(ctx context.Context, in *discount.GetDiscountRequest, opts ...grpc.CallOption) (*discount.GetDiscountResponse, error) {
	out := new(discount.GetDiscountResponse)
	out.Percentage = 0.05
	return out, nil
}

func TestPrice(t *testing.T) {

	verifyTotal := func(t *testing.T, checkout CheckoutRespose) {
		t.Helper()
		totalExpected := 0
		for _, p := range checkout.Products {
			totalExpected += p.TotalAmount
		}
		if checkout.TotalAmount != totalExpected {
			t.Errorf("Total: result '%v', expected '%v'", checkout.TotalAmount, totalExpected)
		}
	}

	verifyDiscount := func(t *testing.T, checkout CheckoutRespose) {
		t.Helper()
		totalExpected := 0
		for _, p := range checkout.Products {
			totalExpected += p.Discount
		}
		if checkout.TotalDiscount != totalExpected {
			t.Errorf("Discount: result '%v', expected '%v'", checkout.TotalDiscount, totalExpected)
		}
	}

	verifyTotalAndDiscount := func(t *testing.T, checkout CheckoutRespose) {
		t.Helper()
		if checkout.TotalAmountWithDiscount != (checkout.TotalAmount - checkout.TotalDiscount) {
			t.Errorf("Total and Discount: result '%v', expected '%v'", checkout.TotalAmountWithDiscount, (checkout.TotalAmount - checkout.TotalDiscount))
		}
	}

	t.Run("Calculate Discount", func(t *testing.T) {
		var mock MockGrpcDiscountClient
		priceService := NewPriceService(&mock)
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
		priceService.CalculateDiscount(&checkout)

		verifyTotal(t, checkout)
		verifyDiscount(t, checkout)
		verifyTotalAndDiscount(t, checkout)
	})
}
