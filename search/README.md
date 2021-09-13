# Search

## publish subgraph to Apollo Studio

```sh
APOLLO_KEY=find-key-in-studio \
  rover subgraph publish search-demo-zb91la@current \
  --convert \
  --name search --schema ./graph/schema.graphql \
  --routing-url http://localhost:4002/query
```