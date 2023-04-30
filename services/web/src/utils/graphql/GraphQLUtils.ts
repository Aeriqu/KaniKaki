import { HttpLink } from "@apollo/client";
import { GraphQLClient } from "./GraphQLClient";

const GRAPHQL_IDENTITY: string = 'graphql_identity';
const GRAPHQL_TOKEN: string = 'graphql_token';

/**
 * initLocalStorage initializes the local storage with empty values to the
 * graphql_identity and graphql_token items. This can also be used to clear
 * the credentials upon the case of expired token/logout.
 */
function initLocalStorage() {
  localStorage.setItem(GRAPHQL_IDENTITY, '');
  updateToken('');
}

/**
 * getIdentity obtains the identity used for graphql from localstorage
 * @returns identity ? identity : null
 */
function getIdentity(): string | null {
  return localStorage.getItem(GRAPHQL_IDENTITY);
}

/**
 * getToken obtains the token used for graphql from localstorage
 * @returns token ? token : null
 */
function getToken(): string | null {
  return localStorage.getItem(GRAPHQL_TOKEN);
}

/**
 * updateIdentity updates the identity in local storage
 * @param identity The new identity to use
 */
function updateIdentity(identity: string) {
  localStorage.setItem(GRAPHQL_IDENTITY, identity)
}

/**
 * updateToken updates the token in local storage and in the graphql client.
 * @param token The new token to use
 */
function updateToken(token: string) {
  localStorage.setItem(GRAPHQL_TOKEN, token)
  GraphQLClient.setLink(new HttpLink({
    uri: 'http://localhost:80/api',
    headers: {
      'Authorization': `Bearer ${token}`
    },
  }));
}

export const GraphQLUtils = {
  initLocalStorage,
  getIdentity,
  getToken,
  updateIdentity,
  updateToken,
}