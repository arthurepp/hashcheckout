package main

import (
	"context"

	discount "github.com/arthurepp/hashcheckout/api/grpc"
)

type Price interface {
	CalculateDiscount(cResponse *CheckoutRespose)
}

type PriceService struct {
	Client discount.DiscountClient
}

func NewPriceService(c discount.DiscountClient) PriceService {
	return PriceService{Client: c}
}

func (p *PriceService) CalculateDiscount(cResponse *CheckoutRespose) {
	totalDiscount := 0
	for i := range cResponse.Products {
		d, err := getDiscount(p.Client, cResponse.Products[i].ID)
		if err != nil {
			break
		}
		cResponse.Products[i].Discount = int(d * float32(cResponse.Products[i].UnitAmount))
		totalDiscount += cResponse.Products[i].Discount
	}
	cResponse.TotalDiscount = totalDiscount
	cResponse.TotalAmountWithDiscount = cResponse.TotalAmount - cResponse.TotalDiscount
}

func getDiscount(client discount.DiscountClient, productId int) (float32, error) {
	response, err := client.GetDiscount(context.Background(), &discount.GetDiscountRequest{ProductID: int32(productId)})
	if err != nil {
		return 0, err
	}
	return response.Percentage, err
}
