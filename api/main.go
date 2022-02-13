package main

import (
	"log"
	"time"

	discount "github.com/arthurepp/hashcheckout/api/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setupRouter(cartAPI CartAPI, port string) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/")
	{
		v1.POST("/", cartAPI.Checkout)

	}
	r.Run(":" + port)
	return r
}

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	port := GetEnvOrFile("CHECKOUT_API_PORT")
	filePath := GetEnvOrFile("CART_DB_FILE_PATH")
	discoundServiceUrl := GetEnvOrFile("CART_DISCOUNT_SERVICE_URL")
	blackFridayDate := GetEnvOrFile("CART_BLACK_FRIDAY_DATE")

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

	shutdown := make(chan bool)
	go func() {
		setupRouter(cartAPI, port)
		shutdown <- true
	}()
	<-shutdown
}
