package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emadghaffari/go-oauth"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_items-api/services"
	"github.com/emadghaffari/rest_items-api/utils/httputils"

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
		fmt.Println("err")
		httputils.ResponseError(w, err)
		return
	}
	fmt.Println(r.Header.Get("X-Caller-Id"))
	fmt.Println(r.Header.Get("X-Client-Id"))
	saller := oauth.GetCallerID(r)
	if saller == 0 {
		resErr := errors.HandlerBadRequest("invalid caller ID")
		httputils.ResponseError(w, resErr)
		return
	}

	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resErr := errors.HandlerBadRequest("invalid Body request")
		httputils.ResponseError(w, resErr)
		return
	}
	defer r.Body.Close()

	var item items.Item
	if err := json.Unmarshal(responseBody, &item); err != nil {
		resErr := errors.HandlerBadRequest(fmt.Sprintf("error in Unmarshal requestBody %v", err))
		httputils.ResponseError(w, resErr)
		return
	}
	item.Seller = saller

	result, createdErr := services.ItemService.Create(item)
	if createdErr != nil {
		// TODO; return error
		httputils.ResponseError(w, createdErr)
		return
	}
	fmt.Println(result)
	httputils.ResponseJSON(w, http.StatusCreated, result)
}

// Get func
// get a item with ID
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
