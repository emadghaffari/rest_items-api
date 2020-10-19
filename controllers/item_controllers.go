package controllers

import (
	"fmt"
	"net/http"

	"github.com/emadghaffari/rest_items-api/services"

	"github.com/emadghaffari/go-oauth"
	"github.com/emadghaffari/rest_items-api/domain/items"
)

var (
	// ItemController var
	ItemController itemsControllerInterface = &itemsController{}
)

// itemsControllerInterface interface
type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// ItemsController struct
//  handler for items
type itemsController struct{}

// Create func
// create a new item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error
	}
	item := &items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemService.Create(item)
	if err != nil {
		// TODO; return error
	}
	fmt.Println(result)
}

// Get func
// get a item with ID
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
