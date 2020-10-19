package elasticsearch

import (
	"context"
	"log"
	"os"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/olivere/elastic"

	"time"
)

var (
	// Client var
	// client for ELK stack
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	Index(interface{}) (*elastic.IndexResponse, *errors.ResError)
	SetClient(*elastic.Client)
}

type esClient struct {
	client *elastic.Client
}

// Init func
func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("elasticsearch:9200"),
		elastic.SetBasicAuth("elastic", "changeme"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}
	Client.SetClient(client)
}

func (c *esClient) Index(txt interface{}) (*elastic.IndexResponse, *errors.ResError) {
	ctx := context.Background()
	elk, err := c.client.Index().Do(ctx)
	if err != nil {
		return nil, errors.HandlerInternalServerError("Handler internal error in Do index", err)
	}
	return elk, nil
}

func (c *esClient) SetClient(client *elastic.Client) {
	c.client = client
}
