package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type CartAPI struct {
	ProductService     Product
	PriceService       Price
	BlackFridayService BlackFriday
}

func NewCartAPI(prod Product, price Price, date BlackFriday) CartAPI {
	return CartAPI{
		ProductService:     prod,
		PriceService:       price,
		BlackFridayService: date,
	}
}

func (a *CartAPI) Checkout(c *gin.Context) {
	var checkout CheckoutRequest
	if err := c.ShouldBind(&checkout); err == nil {
		response, err := a.ProductService.CreateCheckout(checkout)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}

		var wg sync.WaitGroup
		wg.Add(2)
		go func(response *CheckoutRespose) {
			defer wg.Done()
			a.PriceService.CalculateDiscount(response)
		}(&response)
		go func(response *CheckoutRespose) {
			defer wg.Done()
			a.BlackFridayService.AddGiftOnBlackFriday(response)
		}(&response)
		wg.Wait()

		c.JSON(200, response)

	} else {
		c.JSON(400, err.Error())
	}
}
