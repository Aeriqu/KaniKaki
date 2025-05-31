import { handwrittenFont } from '@/fonts/fonts';
import { Kanji } from '@/types/Kanji';

interface KanjiSelectableProp {
  kanji: Kanji,
  selectHandler: any,
}

export function KanjiSelectable({ kanji, selectHandler }: Readonly<KanjiSelectableProp>) {
  return (
    <label htmlFor={kanji.Character} aria-label={`Select kanji ${kanji.Character}`}>
      <input id={kanji.Character} type='checkbox' className='peer' onChange={selectHandler} hidden />
      <div className='
        group
        bg-slate-800
        hover:bg-slate-600
        p-2 m-2
        text-center
        break-words select-none
        peer-checked:text-sky-500 peer-checked:bg-slate-600'>
        <p className={`text-4xl ${handwrittenFont.className}`}>{kanji.Character}</p>
        <p>{kanji.Meanings[0]}</p>
        <p>Level: {kanji.WanikaniLevel}</p>
      </div>
    </label>
  )
}