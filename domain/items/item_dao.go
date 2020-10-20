package items

import (
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
