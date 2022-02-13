package main

import "github.com/gin-gonic/gin"

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

		c.JSON(200, nil)

	} else {
		c.JSON(400, err.Error())
	}
}
