package main

import (
	"embed"
	"fmt"
	httpLib "net/http"
	"os"
	"system-gift-point/delivery/http"
	"system-gift-point/domain"

	"github.com/gofiber/template/html/v2"
)

//go:embed resource/*
var resourcefs embed.FS

func main() {
    domain := domain.ConstructDomain()
    engine := html.NewFileSystem(httpLib.FS(resourcefs), ".html")
    app := http.NewHttpDelivery(domain, engine)
    app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
