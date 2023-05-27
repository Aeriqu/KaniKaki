import { gql } from "@apollo/client";
import { DocumentNode } from "graphql";

function queryGetKanjiByLevelRangeDocument(lowerBound: number, upperBound: number): DocumentNode {
	let mutationString = `
	query {
    getKanjiByLevelRange(lowerBound: ${lowerBound}, upperBound: ${upperBound}) {
      Character,
      WanikaniId,
      WanikaniLevel,
      Meanings,
      Onyomi,
      Kunyomi,
      Nanori
    }
  }`;
	return gql(mutationString);
}

export const kanji = {
	queryGetKanjiByLevelRangeDocument,
}