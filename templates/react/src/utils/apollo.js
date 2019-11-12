import { typeDefs } from '../models/typeDefs'
import resolvers from '../controllers/resolvers'
import auth from '../services/auth'

import { ApolloClient } from 'apollo-client'
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import { ApolloLink } from 'apollo-link'
import { RestLink } from 'apollo-link-rest'
import { setContext } from 'apollo-link-context'

const authLink = setContext(async (_, { headers }) => {
  const { jwt } = await auth.getToken()
  return {
    headers: {
      ...headers,
      authorization: jwt ? `Bearer ${jwt}` : '',
    },
  }
})
const restLink = new RestLink({ uri: 'https://swapi.co/api/' })
const httpLink = createHttpLink({ uri: process.env.REACT_APP_GRAPHQL_HOST })
const cache = new InMemoryCache()

// authLink should go before restLink should go before httpLink in the array,
// as httpLink will swallow any calls that should be routed through rest!
const client = new ApolloClient({
  link: ApolloLink.from([authLink, restLink, httpLink]),
  cache,
  resolvers,
  typeDefs,
})

// TODO: defaults deprecated to direct write cache
// ref: https://www.apollographql.com/docs/react/essentials/local-state/
// const defaults = {
//   session: {
//     __typename: 'Session',
//     access_token: null,
//     userGroup: null,
//     userId: null,
//     orgId: null,
//   },
// }
// cache.writeData(defaults)

// const client = new ApolloClient({
//   uri: config.GRAPHQL_HOST,
//   // cache: new InMemoryCache({ addTypename: false }),
//   request: async operation => {
//     const { jwt } = await auth.getToken()
//     operation.setContext({
//       headers: {
//         authorization: jwt ? `Bearer ${jwt}` : '',
//       },
//     })
//   },
//   clientState: {
//     defaults,
//     resolvers,
//     typeDefs,
//   },
// })

export default client
