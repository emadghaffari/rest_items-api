package items

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_items-api/clients/elasticsearch"
)

const (
	indexES = "items"
	docType = "_doc"
)

// Save method
// store new item
func (i *Item) Save() errors.ResError {
	result, err := elasticsearch.Client.Index(indexES, docType, i)
	if err != nil {
		return err
	}
	i.ID = result.Id
	return nil
}

// Get func
// get item with id
func (i *Item) Get() errors.ResError {
	itemID := i.ID
	result, err := elasticsearch.Client.Get(indexES, docType, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.HandlerNotFoundError(fmt.Sprintf("item not found %s", i.ID))
		}
		return err
	}
	if !result.Found {
		return errors.HandlerNotFoundError(fmt.Sprintf("item not found %s", i.ID))
	}

	bytes, marshalErr := result.Source.MarshalJSON()
	if marshalErr != nil {
		return errors.HandlerInternalServerError(fmt.Sprintf("error in MarshalJSON from DB %s", i.ID), err)
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return errors.HandlerInternalServerError(fmt.Sprintf("error in unmarshal data %s", i.ID), err)
	}
	i.ID = itemID
	return nil
}
