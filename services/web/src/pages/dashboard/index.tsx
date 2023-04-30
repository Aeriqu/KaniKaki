import { GraphQLUtils } from '@/utils/graphql/GraphQLUtils';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

export default function Dashboard() {
  let router = useRouter();

  useEffect(() => {
    if (!GraphQLUtils.getToken()) {
      console.log('not logged in. directing to the login page.');
      router.push('login');
    }
  }, [router]);

  return (
    <>DASHBOARD GAMING</>
  );
}