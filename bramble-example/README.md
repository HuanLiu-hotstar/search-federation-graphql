# Bramble gateway demo

## how to run 

- cd bramble-example/first-service && make && ./first-service
- cd bramble-example/second-service && make && ./second-service
- cd bramble-example && bramble -conf config2.json


```sh
curl -H "content-type: application/json" -X POST http://localhost:8082/query -d '{"query":"{randomFoo{id gqlgen hello world}}"}' |jq

response:

{
  "data": {
    "randomFoo": {
      "id": "78",
      "gqlgen": true,
      "hello": "world_78",
      "world": "hello_78"
    }
  }
}

```

## Schema definition

- first service schema definition

```graphql
# This is the prerequsite schema required
# for fedaration by the gateway
directive @boundary on OBJECT | FIELD_DEFINITION

# The Service type provides the gateway a schema
# to merge into the graph and a name/version to
# reference the service with
type Service {
  name: String!
  version: String!
  schema: String!
}

type Query {
  # The service query is used by the gateway when
  # the service is first registered
  service: Service!

  # example Foo type
  foo(id: ID!): Foo @boundary

  # give me a Foo object
  randomFoo: Foo!
}

type Foo @boundary {
  # required by all services that add a field to
  # the Foo type
  id: ID!

  # A field provided by this service
  gqlgen: Boolean!
}

```

- second service schema definition

```graphql
# This is the prerequsite schema required
# for fedaration by the gateway
directive @boundary on OBJECT | FIELD_DEFINITION

# The Service type provides the gateway a schema
# to merge into the graph and a name/version to
# reference the service with
type Service {
  name: String!
  version: String!
  schema: String!
}

type Query {
  # The service query is used by the gateway when
  # the service is first registered
  service: Service!
  foo(id: ID!): Foo @boundary
}

type Foo @boundary {
  # required by all services that add a field to
  # the Foo type
  id: ID!

  # A field provided by this service
  hello: String!
  world: String!
}

```