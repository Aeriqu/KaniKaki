import { KanjiInfo } from "@/components/elements/kanji-info/KanjiInfo";
import { ReactSketchCanvas, ReactSketchCanvasRef } from "react-sketch-canvas";
import { createRef, useEffect, useRef, useState } from 'react';
import { ArrowPathIcon, ArrowUturnLeftIcon, ArrowUturnRightIcon } from '@heroicons/react/24/solid'
import { strokeOrderFont } from "@/fonts/fonts";
import { Kanji } from "@/types/Kanji";
import { KanjiListContinueProp } from "@/types/KanjiListContinueProp";

export default function KanjiReview({ kanjiList, continueHandler }: KanjiListContinueProp) {
  const canvasRef = createRef<ReactSketchCanvasRef>();
  const [answer, setAnswer] = useState('');
  const [activeKanji, setActiveKanji] = useState<Kanji>({
    Character: '',
    WanikaniId: -1,
    WanikaniLevel: -1,
    Meanings: [],
    Onyomi: [],
    Kunyomi: [],
    Nanori: [],
  });
  const nextKanjiIndex = useRef<number>();
  const reviewingKanji = useRef<Kanji[]>();
  const activeKanjiIndex = useRef<number>();
  const correctKanji = useRef<Map<Kanji, boolean>>(new Map<Kanji, boolean>());
  const incorrectKanji = useRef<Map<Kanji, boolean>>(new Map<Kanji, boolean>());

  useEffect(() => {
    nextKanjiIndex.current = 10;
    reviewingKanji.current = kanjiList.slice(0, 10);
    activeKanjiIndex.current = Math.floor(Math.random() * reviewingKanji.current.length);
    setActiveKanji(reviewingKanji.current[activeKanjiIndex.current]);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function canvasClear() {
    canvasRef.current?.clearCanvas();
  }

  function canvasUndo() {
    canvasRef.current?.undo();
  }

  function canvasRedo() {
    canvasRef.current?.redo();
  }

  function showAnswer() {
    setAnswer(activeKanji.Character);
  }

  function nextKanji(correct: boolean) {
    if (correct) {
      if (incorrectKanji.current?.has(activeKanji)) {
        // TODO: Send request to SRS to mark as incorrect
      } else {
        // TODO: Send the request to SRS to mark as correct
        correctKanji.current?.set(activeKanji, true);
      }
      reviewingKanji.current!.splice(activeKanjiIndex.current!, 1);
      if (nextKanjiIndex.current! < kanjiList.length) {
        reviewingKanji.current!.push(kanjiList[nextKanjiIndex.current!]);
        nextKanjiIndex.current!++;
      }
    }
    else {
      incorrectKanji.current?.set(activeKanji, true);
    }

    if (reviewingKanji.current!.length == 0) {
      continueHandler(correctKanji, incorrectKanji);
      return;
    }

    let newRandom = Math.floor(Math.random() * reviewingKanji.current!.length);
    while (reviewingKanji.current!.length > 1 && newRandom === activeKanjiIndex.current) {
      newRandom = Math.floor(Math.random() * reviewingKanji.current!.length);
    }
    activeKanjiIndex.current = newRandom;
    setActiveKanji(reviewingKanji.current![activeKanjiIndex.current]);

    canvasRef.current?.resetCanvas();
    setAnswer('');
  }

  return (
    <div className={`grid grid-cols-3 h-full ${strokeOrderFont.variable}`}>
      {/* Kanji Info */}
      <div className="col-span-2 m-10 p-10 rounded">
        <KanjiInfo
          kanji={activeKanji}
        />
        {answer ?
          <div className="text-center pt-5 flex justify-evenly">
            <div>
              <p className="text-xl">Typed</p>
              <p className="text-8xl">{answer}</p>
            </div>
            <div>
              <p className='font-handwritten text-xl'>Handwritten</p>
              <p className='text-9xl font-handwritten'>{answer}</p>
            </div>
            <div>
              <p className="text-xl">Stroke Order</p>
              <p className='text-9xl font-stroke-order'>{answer}</p>
            </div>
          </div>
          : <></>
        }
      </div>
      {/* Kanji Writing Area */}
      <div className="col-span-1 m-10 mt-20 rounded flex flex-col">
        <div className="grid grid-cols-2">
          <div className="inline-grid aspect-square w-40 mx-auto hover:cursor-crosshair">
            <ReactSketchCanvas
              ref={canvasRef}
              strokeWidth={2}
              strokeColor="black"
            />
          </div>

          {/* Canvas Control Buttons */}
          <div className="inline-grid grid-cols-1 ml-1">
            <button type="button"

              onClick={canvasClear}
              className="bg-violet-500 hover:bg-violet-600 rounded-t-lg p-2 w-12 flex items-center justify-center">
              <ArrowPathIcon className="h-5 inline mr-1" />
            </button>

            <button type="button"
              onClick={canvasUndo}
              className="bg-violet-500 hover:bg-violet-600 p-2 w-12 flex items-center justify-center">
              <ArrowUturnLeftIcon className="h-5 inline mr-1" />
            </button>

            <button type="button"
              onClick={canvasRedo}
              className="bg-violet-500 hover:bg-violet-600 rounded-b-lg p-2 w-12 flex items-center justify-center">
              <ArrowUturnRightIcon className="h-5 inline mr-1" />
            </button>
          </div>
        </div>
        {/* Kanji Progression Buttons */}
        <div className="grid grid-cols-2 w-2/3">
          <button type="button"
            onClick={answer ? undefined : showAnswer}
            className={answer ?
              "bg-gray-500 col-span-2 rounded-lg w-full mt-2 p-2" :
              "bg-sky-500 col-span-2 rounded-lg w-full mt-2 p-2"}>
            Show Answer
          </button>
          <button type="button"
            onClick={answer ? () => { nextKanji(false) } : undefined}
            className={answer ?
              "bg-rose-500 col-span-1 rounded-lg mt-2 mr-1 p-2" :
              "bg-gray-500 col-span-1 rounded-lg mt-2 mr-1 p-2"}>
            Incorrect
          </button>
          <button type="button"
            onClick={answer ? () => { nextKanji(true) } : undefined}
            className={answer ?
              "bg-emerald-500 col-span-1 rounded-lg mt-2 ml-1 p-2" :
              "bg-gray-500 col-span-1 rounded-lg mt-2 ml-1 p-2"}>
            Correct
          </button>
        </div>
      </div>
    </div>
  );
}