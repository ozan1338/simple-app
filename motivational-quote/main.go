package main

import (
	"fmt"
	"motivational-reminder/delivery/http"
	"motivational-reminder/domain"
	"os"
)

func main() {
    domain := domain.ConstructDomain()
    app := http.NewHttpDelivery(domain)
    app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
