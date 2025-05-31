import { UsernamePasswordProp } from '@/types/UsernamePasswordProp';
import { GraphQLUtils } from '@/utils/graphql/GraphQLUtils';
import { auth } from '@/utils/graphql/schema/auth';
import { useMutation } from '@apollo/client';
import { CogIcon } from '@heroicons/react/24/solid'
import { useRouter } from 'next/router';

export default function SignupButton({ username, password, setErrorState }: Readonly<UsernamePasswordProp>) {
  const router = useRouter();
  const [signup, { data, loading, error, reset }] = useMutation(auth.mutationSignupDocument(username, password));

  function signupHandler() {
    signup();
  }

  if (error) {
    setErrorState(error.message);
  }
  else if (loading) {
    return (
      <button type='button'
        className='bg-orange-500 rounded-lg p-2 mr-1'>
        <CogIcon className='animate-spin w-4 h-4 inline mr-1'></CogIcon>Working...
      </button>
    )
  }
  else if (data) {
    if (!data.signup?.token) {
      setErrorState('token was not properly received from the server')
    }
    else {
      GraphQLUtils.updateIdentity(username);
      GraphQLUtils.updateToken(data.signup.token);
      reset();
      router.push('/dashboard');
    }
  }

  return (
    <button type='button'
      onClick={signupHandler}
      className='bg-orange-500 rounded-lg p-2 mr-1'>
      Sign Up</button>
  )
}