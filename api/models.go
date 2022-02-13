package main

type ProductRequest struct {
	ID       int `json:"id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type CheckoutRequest struct {
	Products []ProductRequest `json:"products" binding:"required,dive"`
}

type ProductEntity struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	IsGift      bool   `json:"is_gift" binding:"required"`
}

type ProductResponse struct {
	ID          int  `json:"id" binding:"required"`
	Quantity    int  `json:"quantity" binding:"required"`
	UnitAmount  int  `json:"unit_amount" binding:"required"`
	TotalAmount int  `json:"total_amount" binding:"required"`
	Discount    int  `json:"discount" binding:"required"`
	IsGift      bool `json:"is_gift" binding:"required"`
}

type CheckoutRespose struct {
	TotalAmount             int               `json:"total_amount" binding:"required"`
	TotalAmountWithDiscount int               `json:"total_amount_with_discount" binding:"required"`
	TotalDiscount           int               `json:"total_discount" binding:"required"`
	Products                []ProductResponse `json:"products" binding:"required"`
}
