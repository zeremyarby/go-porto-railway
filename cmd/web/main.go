package main

import (
	"fmt"
	"net/http"

	"github.com/zeremyarby/go-porto/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port 8080")
	http.ListenAndServe(":8080", nil)

}
