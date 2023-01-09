import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { LoadScriptTemplate } from '../components/templates/shared/LoadScriptTemplate';

export default function App({ Component, pageProps }: AppProps) {
  return <Component {...pageProps} />;
}
