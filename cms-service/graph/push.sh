set -ex

APOLLO_KEY=service:search-demo-zb91la:Xx4PGP9HEUNZkyrSnKEn7g \
  rover subgraph publish search-demo-zb91la@current \
  --convert \
  --name cms --schema ./schema.graphql \
  --routing-url http://localhost:4001/query
