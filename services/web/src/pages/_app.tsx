import HeaderLayout from '@/components/layouts/HeaderLayout'
import '@/styles/globals.css'
import { GraphQLClient } from '@/utils/graphql/GraphQLClient'
import { ApolloProvider } from '@apollo/client'
import type { AppProps } from 'next/app'
import { Noto_Sans_JP } from 'next/font/google'

const font = Noto_Sans_JP({
  weight: '400',
  subsets: ['latin'],
})

export default function App({ Component, pageProps }: AppProps) {
  return (
    <main className={font.className}>
      <ApolloProvider client={GraphQLClient}>
        <HeaderLayout>
          <Component {...pageProps} />
        </HeaderLayout>
      </ApolloProvider>
    </main>
  )
}
