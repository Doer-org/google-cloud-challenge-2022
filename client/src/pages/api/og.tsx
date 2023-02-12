import { ImageResponse } from '@vercel/og';
import { NextResponse } from 'next/server';
export const config = {
  runtime: 'edge',
};

export default function OGP(req: NextResponse) {
  const { searchParams } = new URL(req.url);
  const hasTitle = searchParams.has('title');
  const title = hasTitle
    ? searchParams.get('title')?.slice(0, 100)
    : 'タイトル未定';
  return new ImageResponse(
    (
      <div
        style={{
          fontSize: 30,
          background: '#267365',
          width: '100%',
          height: '100%',
          display: 'flex',
          flexFlow: 'column',
          padding: '10px',
        }}
      >
        <img
          src={`${process.env.NEXT_PUBLIC_FRONT_URL}/logo.png`}
          alt={'top'}
          style={{
            width: '10%',
          }}
        />
        <div
          style={{
            fontSize: 30,
            background: '#267365',
            width: '100%',
            height: '100%',
            textAlign: 'center',
            alignItems: 'center',
            color: 'white',
            justifyContent: 'center',
            fontFamily: 'sans-serif',
          }}
        >
          {title}
        </div>
      </div>
    ),
    {
      width: 1200,
      height: 600,
    }
  );
}
