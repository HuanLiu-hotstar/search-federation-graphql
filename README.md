# search-federation-graphql
- search with federation GraphQL

## how to run 

- cd search && go run server.go 
- cd cms-service && go run server.go
- cd gateway && node index.js


## test

- open localhost:4000 in localhost browser
- that will redirect to https://studio.apollographql.com/sandbox/explorer
- you can use apollo studio explore to send graphQL request to gateway 
- query example

```graphql
query Query{
  search(w: "yeh") {
    ContentType
    id
    result {
      id
      contentType
      contentID
      ... on Cms {
        contentID
        contentType
        id
        meta_data {
          url
          title
        }
      }
      ... on ScoreCard {
        contentID
        contentType
        id
        score_card_data {
          title
          team {
            name
            score
          }
        }
      }
    }
  }
}

## response 

{
  "data": {
    "search": [
      {
        "ContentType": "clips",
        "id": "101",
        "result": [
          {
            "id": "101",
            "contentType": "clips",
            "contentID": "101",
            "meta_data": {
              "url": "http://hotstar.com/clips/101",
              "title": "front-end developer"
            }
          }
        ]
      },
      {
        "ContentType": "clips",
        "id": "102",
        "result": [
          {
            "id": "102",
            "contentType": "clips",
            "contentID": "102",
            "meta_data": {
              "url": "http://hotstar.com/clips/102",
              "title": "back-end developer"
            }
          }
        ]
      }
    ]
  }
}
 
```


- query for score card data 

```graphql
query Query{
  search(w: "IPL") {
    ContentType
    id
    result {
      id
      contentType
      contentID
      ... on Cms {
        contentID
        contentType
        id
        meta_data {
          url
          title
        }
      }
      ... on ScoreCard {
        contentID
        contentType
        id
        score_card_data {
          title
          team {
            name
            score
          }
        }
      }
    }
  }
}
```

- response with score card fields

```js
{
  "data": {
    "search": [
      {
        "ContentType": "clips",
        "id": "matchID201",
        "result": [
          {
            "id": "",
            "contentType": "",
            "contentID": "",
            "score_card_data": {
              "title": "IPL201",
              "team": [
                {
                  "name": "player1",
                  "score": 23
                },
                {
                  "name": "player2",
                  "score": 24
                }
              ]
            }
          }
        ]
      }
    ]
  }
}

```