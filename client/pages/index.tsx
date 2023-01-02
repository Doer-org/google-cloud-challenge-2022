import { useState } from 'react';
import Link from 'next/link';
import { MyHead } from '../components/templates/shared/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';

export default function Home() {
  const [authed, setAuthed] = useState(false);
  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <h1>サービス名未定</h1>
        <a href="/debug"> {'> '} debug page</a>
        <div className="my-10 grid grid-cols-1 gap-2">
          {!authed ? (
            <>
              <button onClick={() => setAuthed(true)}>ログイン</button>
              <button>サインイン</button>
            </>
          ) : (
            <>
              <button>
                <Link href="/event/create">募集する</Link>
              </button>
              <button onClick={() => setAuthed(false)}>log out</button>
            </>
          )}
        </div>
      </BasicTemplate>
    </>
  );
}
