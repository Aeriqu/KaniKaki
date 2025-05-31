import { Noto_Sans_JP } from 'next/font/google';
import localFont from 'next/font/local';

export const mainFont = Noto_Sans_JP({
  weight: '100',
  subsets: ['latin']
});
export const handwrittenFont = localFont({
    src: 'LeftHanded/LeftHanded.otf'
});
export const strokeOrderFont = localFont({
  src: 'KanjiStrokeOrders/KanjiStrokeOrders.ttf'
});