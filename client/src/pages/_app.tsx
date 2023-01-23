import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import { Notice } from '../components/templates/shared/Notice/Notice';
App.getInitialProps = async () => ({ pageProps: {} });
export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <Notice type={'Error'} text="errorです" />
      <Component {...pageProps} />
    </QueryClientProvider>
  );
}
