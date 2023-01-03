import Head from 'next/head';
import { Inter } from '@next/font/google';
import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';

const inter = Inter({ subsets: ['latin'] });

export default function Completion() {
  const [isHost, setIsHost] = useState(false);
  const [hasJoined, setHasJoined] = useState(false);
  const event_id = useRouter().query.id;
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <h1>参加フォーム画面</h1>
        <div>event id: {event_id}</div>
        <button onClick={() => setIsHost((v) => !v)}>
          isHost : {isHost ? 'true' : 'false'}
        </button>

        {isHost ? (
          <button>募集を締め切る</button>
        ) : !hasJoined ? (
          <>
            <h2>名前</h2>
            <input></input>
            <h2>ひとこと</h2>
            <input></input>
            <button onClick={() => setHasJoined(true)}> 参加 </button>
          </>
        ) : (
          <>
            <h1>参加済みです</h1>
            <button onClick={() => setHasJoined(false)}> キャンセル </button>
          </>
        )}
      </BasicTemplate>
    </>
  );
}
