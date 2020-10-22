package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/emadghaffari/go-oauth"
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_items-api/services"
	"github.com/emadghaffari/rest_items-api/utils/httputils"
	"github.com/gorilla/mux"

	"github.com/emadghaffari/rest_items-api/domain/items"
	"github.com/emadghaffari/rest_items-api/domain/queries"
)

var (
	// ItemController var
	ItemController itemsControllerInterface = &itemsController{}
)

// itemsControllerInterface interface
type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

// ItemsController struct
//  handler for items
type itemsController struct{}

// Create func
// create a new item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: return error
		httputils.ResponseError(w, err)
		return
	}
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
	vars := mux.Vars(r)
	id := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Get(id)
	if err != nil {
		httputils.ResponseError(w, err)
		return
	}
	httputils.ResponseJSON(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request)  {
	responseBody,err := ioutil.ReadAll(r.Body)
	if err != nil {
		resErr := errors.HandlerBadRequest("invalid Body request")
		httputils.ResponseError(w, resErr)
		return
	}
	defer r.Body.Close()
	var query queries.EsQuery
	if err := json.Unmarshal(responseBody, &query); err != nil {
		resErr := errors.HandlerBadRequest(fmt.Sprintf("error in Unmarshal requestBody %v", err))
		httputils.ResponseError(w, resErr)
		return
	}

	result,err := services.ItemService.Search(query)
	if err != nil {
		resErr := errors.HandlerBadRequest("invalid Body result")
		httputils.ResponseError(w, resErr)
		fmt.Println(resErr)
		return
	}
	httputils.ResponseJSON(w,http.StatusOK,result)
}


// Get func
// get a item with ID
func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.TrimSpace(vars["id"])

	err := services.ItemService.Delete(id)
	if err != nil {
		httputils.ResponseError(w, err)
		return
	}
	httputils.ResponseJSON(w, http.StatusOK, "Success")
}