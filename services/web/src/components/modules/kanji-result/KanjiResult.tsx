import { KanjiBlockLink } from "@/components/elements/kanji-block-link/KanjiBlockLink";
import { Kanji } from "@/types/Kanji";

interface KanjiResultProp {
  correctKanji: Kanji[]
  incorrectKanji: Kanji[],
}

export default function KanjiResult({ correctKanji, incorrectKanji }: KanjiResultProp) {
  let correctResult: string = '';
  let incorrectResult: string = '';

  correctKanji.forEach((kanji: Kanji) => {
    if (correctResult.length > 0) {
      correctResult = correctResult + `, ${kanji.Character}`;
    }
    else {
      correctResult = kanji.Character;
    }
  })

  incorrectKanji.forEach((kanji: Kanji) => {
    if (incorrectResult.length > 0) {
      incorrectResult = incorrectResult + `, ${kanji.Character}`;
    }
    else {
      incorrectResult = kanji.Character;
    }
  })


  return (
    <div className="w-5/6 mx-auto">
      <div className="grid grid-cols-5 h-16">
        <div className="ml-2 p-2 col-span-4 text-xl text-center rounded">Results</div>
        <div className="ml-2 p-2 col-span-1 text-xl rounded">{Math.round((correctKanji.length / (correctKanji.length + incorrectKanji.length)) * 100)}%</div>
      </div>
      <div className="w-full">
        <div className="m-2 p-2 rounded bg-red-700">
          {incorrectKanji.length} Incorrect
        </div>
        <div className="grid grid-cols-1 min-[400px]:grid-cols-3 md:grid-cols-5">
          {incorrectKanji.map((kanji) => (
            <KanjiBlockLink key={kanji.WanikaniId} kanji={kanji} />
          ))}
        </div>
      </div>
      <div className="w-full">
        <div className="m-2 p-2 rounded bg-green-700">
          {correctKanji.length} Correct
        </div>
        <div className="grid grid-cols-1 min-[400px]:grid-cols-3 md:grid-cols-5">
          {correctKanji.map((kanji) => (
            <KanjiBlockLink key={kanji.WanikaniId} kanji={kanji} />
          ))}
        </div>
      </div>
    </div>
  );
}