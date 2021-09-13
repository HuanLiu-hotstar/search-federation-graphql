

require('dotenv').config()
const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require('@apollo/gateway');

const gateway = new ApolloGateway();
const server = new ApolloServer({ gateway, subscriptions:false, tracing:true });
server.listen();
// require('dotenv').config()
// const { ApolloServer } = require('apollo-server');
// const { ApolloGateway } = require('@apollo/gateway');

// const gateway = new ApolloGateway();

// const server = new ApolloServer({
//   gateway,
// });

// server.listen().then(({ url }) => {
//   console.log(`🚀 Gateway ready at ${url}`);
// }).catch(err => {console.error(err)});
