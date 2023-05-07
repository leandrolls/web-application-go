package main

import (
	"net/http"

	"web-application-project/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
