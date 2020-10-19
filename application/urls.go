package application

import (
	"net/http"

	"github.com/emadghaffari/rest_items-api/controllers"
)

func mapURLs() {
	router.HandleFunc("/items", controllers.ItemController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Get).Methods(http.MethodGet)
}
