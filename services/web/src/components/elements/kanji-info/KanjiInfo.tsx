import { KanjiInfoProp } from "@/types/KanjiInfoProp";

export function KanjiInfo({ kanji }: Readonly<KanjiInfoProp>) {
  let meaning, otherMeanings, onyomi, kunyomi, nanori = <span></span>;

  // Main Meaning
  meaning =
    <div className="col-span-3 text-2xl text-center">
      <h2>{kanji.Meanings[0]}</h2>
    </div>;
  // Other Meanings
  if (kanji.Meanings.length > 1) {
    otherMeanings =
      <div className="col-span-3 text-xl text-center">
        {kanji.Meanings.slice(1, kanji.Meanings.length).join(', ')}
      </div>;
  }
  // Onyomi
  if (kanji.Onyomi.length > 0) {
    onyomi =
      <div className="grid grid-cols-5">
        <div className="col-span-1">音読み:</div>
        <div className="col-span-4">{kanji.Onyomi.join(', ')}</div>
      </div>;
  }
  // Kunyomi
  if (kanji.Kunyomi.length > 0) {
    kunyomi =
      <div className="grid grid-cols-5">
        <div className="col-span-1">訓読み:</div>
        <div className="col-span-4">{kanji.Kunyomi.join(', ')}</div>
      </div>;
  }
  // Nanori
  if (kanji.Nanori.length > 0) {
    nanori =
      <div className="grid grid-cols-5">
        <div className="col-span-1">名乗り:</div>
        <div className="col-span-4">{kanji.Nanori.join(', ')}</div>
      </div>;
  }

  return (
    <div className="">
      {meaning}
      {otherMeanings}
      {onyomi}
      {kunyomi}
      {nanori}
    </div>
  )
}