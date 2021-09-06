package data

import (
	"fmt"
	"log"

	"github.com/HuanLiu-hotstar/search-federation-graphql/cms-service/graph/model"
)

func FetchCmsData(id string) ([]model.SearchResultBase, error) {
	res := []model.SearchResultBase{}
	log.Printf("fetch cms data id:%s", id)
	for _, e := range cmsdata {
		if e.ID == id {
			res = append(res, *e)
		}
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return res, nil
}

func FetchScoreData(matchID string) ([]model.SearchResultBase, error) {
	if v, ok := m[matchID]; ok {
		return []model.SearchResultBase{model.ScoreCard{ScoreCardData: v}}, nil
	}
	return nil, fmt.Errorf("not found matchID %s", matchID)

}

var m = map[string]*model.ScoreData{
	"matchID201": &model.ScoreData{
		Title: "IPL201",
		Team: []*model.Team{
			{
				Name:  "player1",
				Score: 23,
			},
			{
				Name:  "player2",
				Score: 24,
			},
		},
	},
}

var cmsdata = []*model.Cms{
	{

		ID:          "101",
		ContentID:   "101",
		ContentType: "clips",
		MetaData: &model.MetaData{
			URL:   "http://hotstar.com/clips/101",
			Title: "front-end developer",
		},
	},
	{

		ID:          "102",
		ContentID:   "102",
		ContentType: "clips",
		MetaData: &model.MetaData{
			URL:   "http://hotstar.com/clips/102",
			Title: "back-end developer",
		},
	},
}
