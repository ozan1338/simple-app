package main

import (
	"embed"
	"favourite-book/delivery/http"
	"favourite-book/domain"
	"fmt"
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
