import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import { Notice } from '../components/templates/shared/Notice/Notice';
import { useNoticeStore } from '../store/noticeStore';
App.getInitialProps = async () => ({ pageProps: {} });
export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient();
  const { notice, changeNotice, resetNotice } = useNoticeStore();
  return (
    <QueryClientProvider client={queryClient}>
      <Notice type={notice.type} text={notice.text} />
      <Component {...pageProps} />
    </QueryClientProvider>
  );
}
