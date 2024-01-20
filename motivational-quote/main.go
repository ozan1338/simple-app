package main

import (
	"embed"
	"fmt"
	"motivational-reminder/delivery/http"
	"motivational-reminder/domain"
	httpLib "net/http"
	"os"

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
