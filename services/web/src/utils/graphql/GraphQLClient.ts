import { ApolloClient, HttpLink, InMemoryCache } from "@apollo/client";

export const GraphQLClient = new ApolloClient({
  link: new HttpLink({
    uri: 'http://localhost:80/api',
  }),
  cache: new InMemoryCache(),
  defaultOptions: {
    watchQuery: {
      fetchPolicy: 'no-cache',
      errorPolicy: 'all',
    },
    query: {
      fetchPolicy: 'no-cache',
      errorPolicy: 'all',
    },
    mutate: {
      fetchPolicy: 'no-cache',
      errorPolicy: 'all',
    },
  },
});

// Set up the token with the proper link header if it's being run in the client.
if (typeof window !== "undefined") {
  const token = localStorage.getItem('graphql_token')

  GraphQLClient.setLink(new HttpLink({
    uri: 'http://localhost:80/api',
    headers: {
      'Authorization': `Bearer ${token ? token : ''}`
    },
  }));
};