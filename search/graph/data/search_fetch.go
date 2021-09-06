package data

import (
	"fmt"

	"github.com/HuanLiu-hotstar/search-federation-graphql/search/graph/model"
)

func FetchSearchData(word string) ([]*model.Content, error) {
	if v, ok := searchData[word]; ok {
		// printresult(v)
		return v, nil
	}
	return nil, fmt.Errorf("no search result")
}

// func printresult(d []*model.Cid) {
// 	for _, e := range d {
// 		log.Printf("d:%+v", e)
// 	}
// }

var searchData = map[string][]*model.Content{
	"yeh": []*model.Content{
		{
			ID:          "101",
			ContentType: "clips",
		},
		{

			ID:          "102",
			ContentType: "clips",
		},
	},
	"IPL": []*model.Content{
		{

			ID:          "matchID201",
			ContentType: "clips",
		},
	},
}
