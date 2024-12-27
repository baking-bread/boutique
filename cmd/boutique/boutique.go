package main

import (
	"fmt"

	"github.com/baking-bread/boutique/internal/handlers/rest/routers"
)

func main() {
	fmt.Println("boutique!")

	router := routers.SetupRouter()

	router.Run()
}
