import Head from 'next/head';
type TProps = {
  title: string;
  description: string;
};
export const MyHead = ({ title, description }: TProps) => {
  return (
    <Head>
      <title>{title}</title>
      <meta property="og:title" content={title} />
      <meta property="og:description" content={description} />
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <meta
        property="og:image"
        content={`${process.env.NEXT_PUBLIC_FRONT_URL}/api/og?title=${title}`}
      />
      <link rel="icon" href="/favicon.ico" />
    </Head>
  );
};
