package main

import (
	"fmt"
	"os"
	"system-gift-point/delivery/http"
	"system-gift-point/domain"
)

func main() {
    domain := domain.ConstructDomain()
    app := http.NewHttpDelivery(domain)
    app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
