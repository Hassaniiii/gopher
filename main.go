package main

import (
	"net/http"

	"github.com/Hassaniiii/gopher/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
