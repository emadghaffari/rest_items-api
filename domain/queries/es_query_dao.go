package queries

import "github.com/olivere/elastic"

// Build meth
func (ch *EsQuery) Build() elastic.Query {
	boolQuery := elastic.NewBoolQuery()
	queries := make([]elastic.Query,0)
	for _, item := range ch.Equals {
		queries = append(queries, elastic.NewMatchQuery(item.Field,item.Value))
	}
	boolQuery.Must(queries...)

	return boolQuery
}