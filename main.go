package main

import (
	"fmt"
	"net/http"
	"router/myrouter"
)

func main() {
	// Create a new router
	router := myrouter.NewRouter()

	// Register middleware for specific routes
	router.Handle("GET", "/", myrouter.HandleHome, myrouter.LogMiddleware)
	router.Handle("GET", "/about", myrouter.HandleAbout, myrouter.LogMiddleware)
	router.Handle("GET", "/users/:id", myrouter.HandleUser,myrouter.LogMiddleware)

	// Start the server on port 8080 with our router
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
