package elasticsearch

import (
	"context"
	"time"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/res_errors/logger"
	"github.com/olivere/elastic"
)

var (
	// Client var
	// client for ELK stack
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	Index(string, interface{}) (*elastic.IndexResponse, errors.ResError)
	SetClient(*elastic.Client)
}

type esClient struct {
	client *elastic.Client
}

// Init func
func Init() {
	logger := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetBasicAuth("elastic", "changeme"),
		elastic.SetHealthcheckInterval(50*time.Second),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
	)
	if err != nil {
		panic(err)
	}
	Client.SetClient(client)
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, errors.ResError) {
	ctx := context.Background()
	elk, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Type(index).
		Do(ctx)
	if err != nil {
		logger.Error("error in index esClient", err)
		return nil, errors.HandlerInternalServerError("internal ELK error in Index", err)
	}
	return elk, nil
}

func (c *esClient) SetClient(client *elastic.Client) {
	c.client = client
}
