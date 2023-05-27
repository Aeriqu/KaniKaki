import { Noto_Sans_JP } from 'next/font/google';
import localFont from 'next/font/local';

export const mainFont = Noto_Sans_JP(
  {
    weight: '100',
    subsets: ['latin'],
    variable: '--font-main',
  }
);
export const handwrittenFont = localFont(
  {
    src: 'LeftHanded/LeftHanded.otf',
    variable: '--font-handwritten',
  }
);
export const strokeOrderFont = localFont(
  {
    src: 'KanjiStrokeOrders/KanjiStrokeOrders.ttf',
    variable: '--font-stroke-order',
  }
);