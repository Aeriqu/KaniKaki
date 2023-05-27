import { Kanji } from '@/types/Kanji';

interface KanjiSelectableProp {
  kanji: Kanji,
  selectHandler: any,
}

export function KanjiSelectable({ kanji, selectHandler }: KanjiSelectableProp) {
  return (
    <label htmlFor={kanji.Character}>
      <input id={kanji.Character} type='checkbox' className='peer' onChange={selectHandler} hidden />
      <div className='
        group
        bg-slate-800
        hover:bg-slate-600
        p-2 m-2
        text-center
        break-words select-none
        peer-checked:text-sky-500 peer-checked:bg-slate-600'>
        <p className={'text-3xl group-hover:font-handwritten peer-checked:[.group_&]:font-handwritten'}>{kanji.Character}</p>
        <p>{kanji.Meanings[0]}</p>
        <p>Level: {kanji.WanikaniLevel}</p>
      </div>
    </label>
  )
}