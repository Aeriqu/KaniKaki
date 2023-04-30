/**
 * Login page contains login related dialogs and prompts.
 */

import { ErrorBanner } from '@/components/elements/error-banner/ErrorBanner';
import LoginButton from '@/components/elements/login-login-button/LoginButton'
import SignupButton from '@/components/elements/login-signup-button/SignupButton'
import { KeyIcon, UserIcon } from '@heroicons/react/24/solid'
import { useState } from 'react'

export default function Login() {
  const [usernameState, setUsernameState] = useState('');
  const [passwordState, setPasswordState] = useState('');
  const [errorState, setErrorState] = useState('');

  return (
    <div className='h-[calc(100vh-5rem)] flex flex-col justify-center items-center'>

      <div className='w-2/3 lg:w-1/3'>
        <ErrorBanner error={errorState}></ErrorBanner>
      </div>

      <div className='w-2/3 lg:w-1/3'>
        <label htmlFor='username-input'>ユーザーネーム</label>
        <div className='relative mb-6'>
          <div className='absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none'>
            <UserIcon className='w-5 h-5 text-gray-500
                                dark:text-gray-400'>
            </UserIcon>
          </div>
          <input type='text' id='username-input' placeholder='Username'
          onChange={event => setUsernameState(event.target.value)}
            className='bg-gray-50 border border-white-300 text-gray-900 text-sm rounded-md block w-full pl-10 p-2.5
                      dark:bg-gray-700 dark:border-white-600 dark:placeholder-gray-400 dark:text-white' />
        </div>
      </div>

      <div className='w-2/3 lg:w-1/3'>
        <label htmlFor='password-input'>パスワード</label>
        <div className='relative mb-6'>
          <div className='absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none'>
            <KeyIcon className='w-5 h-5 text-gray-500
                                dark:text-gray-400'>
            </KeyIcon>
          </div>
          <input type='password' id='password-input' placeholder='Password'
            onChange={event => setPasswordState(event.target.value)}
            className='bg-gray-50 border border-white-300 text-gray-900 text-sm rounded-md block w-full pl-10 p-2.5 
                      dark:bg-gray-700 dark:border-white-600 dark:placeholder-gray-400 dark:text-white' />
            </div>
      </div>

      <div className='grid grid-cols-2 w-2/3 lg:w-1/3'>
        <SignupButton username={usernameState} password={passwordState} setErrorState={setErrorState}></SignupButton>
        <LoginButton username={usernameState} password={passwordState} setErrorState={setErrorState}></LoginButton>
      </div>

    </div>
  )
}