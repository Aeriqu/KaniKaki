import { XCircleIcon } from '@heroicons/react/24/solid'
import toast from 'react-hot-toast'

function error(text: string) {
  toast.custom(
    <div className='bg-red-500 p-2 rounded flex items-center select-none'>
      <XCircleIcon className='h-5' />
      <p>
        {text}
      </p>
    </div>
  )
}

export const Notification = {
  error,
}