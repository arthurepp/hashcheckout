package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
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
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	port := GetEnvOrFile("CHECKOUT_API_PORT")

	cartAPI := NewCartAPI(nil, nil, nil)

	shutdown := make(chan bool)
	go func() {
		setupRouter(cartAPI, port)
		shutdown <- true
	}()
	<-shutdown
}
