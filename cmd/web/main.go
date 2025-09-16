package main

import (
	"fmt"
	"net/http"

	"github.com/zeremyarby/go-porto/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/contact", handlers.Contact)
	http.HandleFunc("/myblog", handlers.MyBlog)

	// Get Static Files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting application on port 8080")
	http.ListenAndServe(portNumber, nil)

}
