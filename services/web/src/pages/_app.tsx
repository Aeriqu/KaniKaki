import HeaderLayout from '@/components/layouts/HeaderLayout'
import { mainFont, handwrittenFont } from '@/fonts/fonts'
import '@/styles/globals.css'
import { GraphQLClient } from '@/utils/graphql/GraphQLClient'
import { ApolloProvider } from '@apollo/client'
import type { AppProps } from 'next/app'
import { Toaster } from 'react-hot-toast'

export default function App({ Component, pageProps }: AppProps) {
  return (
    <main className={`${mainFont.variable} ${handwrittenFont.variable} font-main`}>
      <ApolloProvider client={GraphQLClient}>
        <HeaderLayout>
          <Component {...pageProps} />
          <Toaster position='bottom-right' />
        </HeaderLayout>
      </ApolloProvider>
    </main>
  )
}
