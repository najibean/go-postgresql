package main

import (
	"fmt"
	submodels "go-postgres/models/subModels"
)

func main() {
	// r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 9090...")

	// log.Fatal(http.ListenAndServe(":9090", r))

	submodels.Submodel()

}
