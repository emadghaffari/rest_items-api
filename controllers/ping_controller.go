package controllers

import "net/http"

var (
	// PingController var
	PingController pingControllerInterface = &pingController{}
)

// pingControllerInterface interface
type pingControllerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// pingController struct
//  handler for ping
type pingController struct{}

// Get func
// get a item with ID
func (c *pingController) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}
