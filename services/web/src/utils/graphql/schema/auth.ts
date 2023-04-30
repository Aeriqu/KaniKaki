import { DocumentNode, gql } from "@apollo/client";

function mutationLoginDocument(identifier: string, credential: string): DocumentNode {
	let mutationString = `
		mutation {
			login(identifier: "${identifier}", credential: "${credential}") {
				token
			}
		}`;
	return gql(mutationString);
};

function mutationLogoutDocument(identifier: string): DocumentNode {
	let mutationString = `
	mutation {
		logout(identifier: "${identifier}") {
			token
		}
	}`;
	return gql(mutationString);
};

function mutationSignupDocument(identifier: string, credential: string): DocumentNode {
	let mutationString = `
		mutation {
			signup(identifier: "${identifier}", credential: "${credential}") {
				token
			}
		}`;
	return gql(mutationString)
};

function mutationRefreshTokenDocument(identifier: string): DocumentNode {
	let mutationString = `
	mutation {
		refreshToken(identifier: "${identifier}") {
			token
		}
	}`;
	return gql(mutationString);
}

export const auth = {
	mutationLoginDocument,
	mutationLogoutDocument,
	mutationSignupDocument,
	mutationRefreshTokenDocument,
}