import { KanjiSelectable } from "@/components/elements/kanji-selectable/KanjiSelectable";
import { KanjiListContinueProp } from "@/types/KanjiListContinueProp";
import { Notification } from "@/utils/notification/Notification";

export default function KanjiSelect({ kanjiList, continueHandler }: KanjiListContinueProp) {

  let selectedKanjiMap = new Map<string, boolean>();

  function selectHandler(event: Event) {
    let eventElement = (event.target as HTMLInputElement)
    if (eventElement.checked) {
      selectedKanjiMap.set(eventElement.id, true);
    }
    else {
      selectedKanjiMap.delete(eventElement.id);
    }
    console.log(selectedKanjiMap);
  }

  function continueClickHandler() {
    if (selectedKanjiMap.size > 0) {
      continueHandler(selectedKanjiMap);
    } else {
      Notification.error('Please select at least 1 kanji');
    }
  }

  return (
    <div>
      <div className="grid grid-cols-1 min-[400px]:grid-cols-3 md:grid-cols-5">
        {kanjiList.map((kanji) => (
          <KanjiSelectable key={kanji.WanikaniId} kanji={kanji} selectHandler={selectHandler} />
        ))}
      </div>
      <button type='button'
        onClick={continueClickHandler}
        className='bg-sky-500 rounded-lg w-full p-2 ml-1 mb-2'>
        Select Kanji</button>
    </div>
  )
}