package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	discount "github.com/arthurepp/hashcheckout/api/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallCheckout(cartAPI CartAPI, request CheckoutRequest) []byte {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	cartAPI.Checkout(c)

	b, _ := ioutil.ReadAll(w.Body)
	return b

}

func TestCallApi(t *testing.T) {
	filePath := GetEnvOrFile("CART_DB_FILE_PATH")
	discoundServiceUrl := GetEnvOrFile("CART_DISCOUNT_SERVICE_URL")
	blackFridayDate := time.Now().Truncate(time.Hour * 24).Format("2006-01-02")

	productService, err := NewProductService(filePath)
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(discoundServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := discount.NewDiscountClient(conn)
	priceService := NewPriceService(client)
	date, err := time.Parse("2006-01-02", blackFridayDate)
	if err != nil {
		panic(err)
	}

	blackFridayService := NewBlackFridayService(date)

	cartAPI := NewCartAPI(&productService, &priceService, &blackFridayService)

	t.Run("Ok", func(t *testing.T) {
		request := CheckoutRequest{
			Products: []ProductRequest{
				{
					ID:       1,
					Quantity: 1,
				},
			},
		}

		result := CallCheckout(cartAPI, request)

		var response CheckoutRespose
		err = json.Unmarshal(result, &response)
		if err != nil {
			panic(err)
		}
		if len(response.Products) != 2 {
			t.Errorf("error processing correct payload")
		}
	})

	t.Run("Invalid product", func(t *testing.T) {
		request := CheckoutRequest{
			Products: []ProductRequest{
				{
					ID:       134,
					Quantity: 1,
				},
			},
		}

		result := CallCheckout(cartAPI, request)

		if string(result) != "\"invalid product id 134\"" {
			t.Errorf("error sending invalid product")
		}
	})

	t.Run("Invalid payload", func(t *testing.T) {
		request := CheckoutRequest{
			Products: []ProductRequest{
				{
					Quantity: 1,
				},
			},
		}

		result := CallCheckout(cartAPI, request)

		if string(result) != "\"Key: 'CheckoutRequest.Products[0].ID' Error:Field validation for 'ID' failed on the 'required' tag\"" {
			t.Errorf("error sending invalid product")
		}
	})
}
