import Head from 'next/head';
type TProps = {
  title: string;
  description: string;
};
export const MyHead = ({ title, description }: TProps) => {
  return (
    <Head>
      <title>{title}</title>
      <meta name="description" content={description} />
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <link rel="icon" href="/favicon.ico" />
    </Head>
  );
};
