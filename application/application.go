package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emadghaffari/rest_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

// StartApplication func
func StartApplication() {
	fmt.Println("servcer started")
	elasticsearch.Init()
	mapURLs()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  50 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
