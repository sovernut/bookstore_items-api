package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sovernut/bookstore_items-api/src/domain/items"
	"github.com/sovernut/bookstore_items-api/src/services"
	"github.com/sovernut/bookstore_items-api/src/utils/http_utils"
	"github.com/sovernut/bookstore_oauth-go/oauth"
	"github.com/sovernut/bookstore_utils-go/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	// Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}
	// sellerId := oauth.GetCallerId(r)
	// if sellerId == 0 {
	// 	respErr := rest_errors.NewUnauthorizedError("invalid access token")
	// 	http_utils.RespondError(w, respErr)
	// 	return
	// }

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := *rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := *rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = 0 // TODO: implment func to get SellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if &createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if &err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}
