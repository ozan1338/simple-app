package main

import (
	"favourite-book/delivery/http"
	"favourite-book/domain"
	"fmt"
	"os"
)

func main() {
	domain := domain.ConstructDomain()
	app := http.NewHttpDelivery(domain)
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
