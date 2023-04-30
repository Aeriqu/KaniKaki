import { UsernamePasswordProp } from '@/types/UsernamePasswordProp';
import { GraphQLUtils } from '@/utils/graphql/GraphQLUtils';
import { auth } from '@/utils/graphql/schema/auth';
import { useMutation } from '@apollo/client';
import { CogIcon } from '@heroicons/react/24/solid';
import { useRouter } from 'next/router';

export default function LoginButton({ username, password, setErrorState }: UsernamePasswordProp) {
  const router = useRouter();
  const [login, { data, loading, error, reset }] = useMutation(auth.mutationLoginDocument(username, password));

  function loginHandler() {
    login();
  }

  if (error) {
    setErrorState(error.message);
  }
  else if (loading) {
    return (
      <button type='button'
        className='bg-sky-500 rounded-lg w-95 p-2 mr-1'>
        <CogIcon className='animate-spin w-4 h-4 inline mr-1'></CogIcon>Working...
      </button>
    )
  }
  else if (data) {
    if (!data.login || !data.login.token) {
      setErrorState('token was not properly received from the server')
    }
    else {
      GraphQLUtils.updateIdentity(username);
      GraphQLUtils.updateToken(data.login.token);
      reset();
      router.push('/dashboard');
    }
  }

  return (
    <button type='button'
      onClick={loginHandler}
      className='bg-sky-500 rounded-lg w-95 p-2 ml-1'>
      Log In</button>
  )
}