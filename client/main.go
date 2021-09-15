package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/graphql"
)

type Team struct {
	Name  graphql.String
	Score graphql.Int
}
type ScoreCardData struct {
	Title graphql.String
	Team  []Team `graphql:"team"`
}
type ScoreCard struct {
	// ContentID     graphql.String `graphql:"contentID"`
	// ContentType   graphql.String `graphql:"contentType"`
	ScoreCardData ScoreCardData `graphql:"score_card_data"`
}
type MetaData struct {
	Url   graphql.String
	Title graphql.String
}
type Cms struct {
	// ContentID   graphql.String `graphql:"contentID"`
	// ContentType graphql.String `graphql:"contentType"`
	MetaData MetaData `graphql:"meta_data"`
}

type Result struct {
	// ContentType graphql.String `graphql:"contentType"`
	// ContentID   graphql.String `graphql:"contentID"`
	Cms       `graphql:"... on Cms"`
	ScoreCard `graphql:"... on ScoreCard"`
}

func searchReq() interface{} {
	var q struct {
		Search []struct {
			ContentType graphql.String `graphql:"ContentType"`
			Id          graphql.String `graphql:"id"`
			Result      []Result       `graphql:"result"`
			// Result      []struct {
			// ContentType graphql.String `graphql:"contentType"`
			// ContentID   graphql.String `graphql:"contentID"`
			// Cms         struct {
			// 	ContentID   graphql.String `graphql:"contentID"`
			// 	ContentType graphql.String `graphql:"contentType"`
			// 	MetaData    struct {
			// 		Url   graphql.String
			// 		Title graphql.String
			// 	} `graphql:"meta_data"`
			// } `graphql:"... on Cms"`
			// Cms       `graphql:"... on Cms"`
			// ScoreCard `graphql:"... on ScoreCard"`
			// ScoreCard struct {
			// 	ContentID     graphql.String `graphql:"contentID"`
			// 	ContentType   graphql.String `graphql:"contentType"`
			// 	ScoreCardData `graphql:"score_card_data"`
			// ScoreCardData struct {
			// 	Title graphql.String
			// 	Team  []struct {
			// 		Name  graphql.String
			// 		Score graphql.Int
			// 	} `graphql:"team"`
			// } `graphql:"score_card_data"`
			// } `graphql:"... on ScoreCard"`
			// } `graphql:"result"`
		} `graphql:"search(w: $w)"`
	}
	// return q
	client := graphql.NewClient("http://localhost:4000/query", nil)
	// q := searchReq()
	m := map[string]interface{}{
		"w": graphql.String("IPL"),
	}
	err := client.Query(context.Background(), &q, m)
	if err != nil {
		log.Fatal("err:", err)
	}
	json.NewEncoder(os.Stdout).Encode(q)
	return q
}
func sendsearch(w string) {
	searchReq()
	return
	client := graphql.NewClient("http://localhost:4000/query", nil)
	q := searchReq()
	m := map[string]interface{}{
		"w": w,
	}
	err := client.Query(context.Background(), &q, m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", q)

}
func main() {
	sendsearch("IPL")
	return
	client := graphql.NewClient("http://localhost:8082/query", nil)
	// Use client...
	var q struct {
		RandomFoo struct {
			// Title  graphql.String `graphql:"title"`
			Hello   graphql.String `graphql:"hello"`
			World   graphql.String
			ID      graphql.String
			Someone graphql.String
			Gqlgen  graphql.Boolean
			//Height graphql.Float `graphql:"height(unit: METER)"`
		} `graphql:"randomFoo"`
	}
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("result:%+v", q)
	fmt.Printf("%+v\n", q)
}
