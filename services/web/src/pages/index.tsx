/**
 * Home page contains the content of the home page.
 */

import Link from 'next/link';

export default function Home() {
  return (
    <div className='h-[calc(100vh-5rem)] flex flex-col justify-items-center justify-around text-center pb-6'>
      <div>
        <h1 className='text-4xl'>KaniKaki</h1>
      </div>
      <div>
        <p>Imagine an app that lets you practice writing kanji.</p>
      </div>
      <div>
        <p>That lets you sync with the kanji you&apos;ve learned on a certain aligator crab website.</p>
      </div>
      <div>
        <p>That has its own SRS tracking.</p>
      </div>
      <div>
        <p className='underline decoration-solid'><a href='https://github.com/Aeriqu/KaniKaki'>That&apos;s fully open source!</a></p>
      </div>
      <div>
        <p>That has a funny looking crab with a paint brush (no wani here).</p>
      </div>
      <div>
        <p>I&apos;m still imagining what it could be too.</p>
      </div>
      <div>
        <Link className='underline decoration-solid' href="/demo">You can try it out here!</Link>
      </div>
    </div>
  )
}
