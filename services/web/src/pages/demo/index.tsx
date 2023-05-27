import KanjiReview from "@/components/modules/kanji-review/KanjiReview";
import KanjiSelect from "@/components/modules/kanji-select/KanjiSelect";
import { Kanji } from "@/types/Kanji";
import { kanji } from "@/utils/graphql/schema/kanji";
import { useQuery } from "@apollo/client";
import { MutableRefObject, useRef, useState } from 'react';
import { handwrittenFont } from "@/fonts/fonts";
import KanjiResult from "@/components/modules/kanji-result/KanjiResult";

export default function Demo() {
  const { loading, error, data } = useQuery(kanji.queryGetKanjiByLevelRangeDocument(1, 3));
  const [demoState, setDemoState] = useState('select');
  const [kanjiList, setKanjiList] = useState<Kanji[]>();
  const kanjiResultsCorrect = useRef<Kanji[]>([]);
  const kanjiResultsIncorrect = useRef<Kanji[]>([]);

  function kanjiSelectContinueHandler(selectedKanjiMap: Map<string, boolean>) {
    setKanjiList(data['getKanjiByLevelRange'].filter((kanji: Kanji) => {
      return selectedKanjiMap.has(kanji.Character);
    }));
    // setDemoState('learn');
    setDemoState('review');
  }

  function kanjiReviewContinueHandler(correctKanji: MutableRefObject<Map<Kanji, boolean>>, incorrectKanji: MutableRefObject<Map<Kanji, boolean>>) {
    kanjiResultsCorrect.current = Array.from(correctKanji.current.keys());
    kanjiResultsIncorrect.current = Array.from(incorrectKanji.current.keys());
    setDemoState('result');
  }

  if (loading) {
    return (
      <div className="p-10 rounded text-center">
        Loading...
      </div>
    );
  }
  else if (error) {
    return (
      <div className="p-10 rounded text-center">
        <p>Error obtaining kanji</p>
        <p>({error.message})</p>
      </div>
    );
  } else {
    data['getKanjiByLevelRange'].sort((a: Kanji, b: Kanji): number => {
      if (a.WanikaniLevel != b.WanikaniLevel) {
        return a.WanikaniLevel - b.WanikaniLevel
      }
      return a.Character.charCodeAt(0) - b.Character.charCodeAt(0);
    });
  }

  // Kanji Select State
  if (demoState === 'select') {
    return (
      <div className="w-4/5 mx-auto text-center">
        <h2 className={'text-3xl ' + handwrittenFont.className}>Select the kanji to test in the demo</h2>
        <p className={'text-lg ' + handwrittenFont.className}>※ This only displays the free levels of WaniKani (1-3)</p>
        <KanjiSelect
          kanjiList={data['getKanjiByLevelRange']}
          continueHandler={kanjiSelectContinueHandler} />
      </div>
    );
  }

  // Learn State
  else if (demoState === 'learn') {
    return (
      <>Learn Area</>
    );
  }

  // Review State
  else if (demoState === 'review') {
    return (
      <div>
        <KanjiReview kanjiList={kanjiList!} continueHandler={kanjiReviewContinueHandler} />
      </div>
    );
  }

  return (
    <KanjiResult correctKanji={kanjiResultsCorrect.current} incorrectKanji={kanjiResultsIncorrect.current}></KanjiResult>
  );
}