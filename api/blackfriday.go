package main

import (
	"time"
)

type BlackFriday interface {
	AddGiftOnBlackFriday(cResponse CheckoutRespose)
}

type BlackFridayService struct {
	BlackFridayDate time.Time
}

func NewBlackFridayService(date time.Time) BlackFridayService {
	return BlackFridayService{BlackFridayDate: date}
}

func (bf *BlackFridayService) AddGiftOnBlackFriday(cResponse *CheckoutRespose) {
	isBackFriday := isBackFriday(bf.BlackFridayDate)
	if isBackFriday {
		cResponse.Products = append(cResponse.Products, ProductResponse{
			ID:          0,
			Quantity:    1,
			UnitAmount:  0,
			TotalAmount: 0,
			Discount:    0,
			IsGift:      true,
		})
	}
}

func isBackFriday(date time.Time) bool {
	return time.Now().Truncate(time.Hour*24).Format("2006-01-02") == date.Truncate(time.Hour*24).Format("2006-01-02")
}
