import { GraphQLUtils } from "@/utils/graphql/GraphQLUtils";
import { auth } from "@/utils/graphql/schema/auth";
import { useMutation } from "@apollo/client";
import { CogIcon } from "@heroicons/react/24/solid";
import Link from "next/link";
import { useRouter } from "next/router";

export default function HeaderLoginLogout() {
  let router = useRouter();
  const [logoutMutation, { data, loading, error, reset }] = useMutation(auth.mutationLogoutDocument(GraphQLUtils.getIdentity()!));

  function logout() {
    logoutMutation();
  }

  if (error) {
    console.log('error logging out', error);
  }
  else if (loading) {
    return (
      <CogIcon className='animate-spin w-4 h-4 inline mr-1'></CogIcon>
    )
  }
  else if (data) {
    GraphQLUtils.initLocalStorage();
    reset();
    router.push('/');
  }

  if (GraphQLUtils.getToken()) {
    return (
      <a href="#" onClick={logout}>Logout</a>
    )
  }

  return (
    <Link href="/login">Login</Link>
  )
}