import { handwrittenFont } from '@/fonts/fonts';
import { Kanji } from '@/types/Kanji';

interface KanjiBlockLinkProp {
  kanji: Kanji,
}

export function KanjiBlockLink({ kanji }: Readonly<KanjiBlockLinkProp>) {
  return (
    <div className='
        group
        bg-slate-800
        hover:bg-slate-600
        p-2 m-2
        text-center
        break-words select-none'>
      <p className={`text-3xl group-hover:${handwrittenFont.className}`}>{kanji.Character}</p>
      <p>{kanji.Meanings[0]}</p>
    </div>
  )
}