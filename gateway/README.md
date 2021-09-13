

# Gateway

- fetch supergraphql for federation graph


```sh 
APOLLO_KEY=find-key-in-studio rover graph fetch  search-demo-zb91la@current > supergraph.graphql

```



## Compose supergraph in local with rover

- subgrap configuration to compose
- reference docs https://www.apollographql.com/docs/rover/supergraphs/

```yml
# subgraphs.yml
subgraphs:
  search:
    routing_url: http://localhost:4002/query
    schema:
      file: ../search/graph/schema.graphql
  cms:
    routing_url: http://localhost:4001/query
    schema:
      file: ../cms-service/graph/schema.graphql

```

- compose a supergraph

```sh

# compose supergraph and use it in main.js
rover supergraph compose --config ./supergraph.yaml > supergraph.gql

```

- start the gateway using `supergraph`

```js
// in main.js
const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');
const { readFileSync } = require('fs');

const supergraphSdl = readFileSync('./supergraph.gql').toString();

// Initialize an ApolloGateway instance and pass it
// the supergraph schema
const gateway = new ApolloGateway({
  supergraphSdl
});

// Pass the ApolloGateway to the ApolloServer constructor
const server = new ApolloServer({
  gateway
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});


```

- run service
- start the search service & start cms-service
- start gateway service `node main.js`

## Use Apollo Studio supergraph

```sh

APOLLO_KEY=service:search-demo-zb91la:Xx4PGP9HEUNZkyrSnKEn7g \
APOLLO_GRAPH_REF=search-demo-zb91la@current \
node remote-index.js

```
