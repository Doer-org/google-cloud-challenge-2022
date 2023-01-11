import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { useCopyToClipboard } from 'usehooks-ts';

export default function Participate() {
  // ここはイベントを作成したときにリンクをコピーする画面
  // TODO:URLコピー機能
  const [value, copy] = useCopyToClipboard();
  const router = useRouter();
  const [name, setName] = useState('');
  const [word, setWord] = useState('');
  const event_id = useRouter().query.id;
  return (
    <>
      <MyHead title="イベントURLコピー" description="" />
      <BasicTemplate className="text-center">
        <EventInfo />
        <Button
          className="flex m-auto my-5"
          onClick={() => {
            copy(router.pathname); // とりあえずcurrenturlをコピー
          }}
        >
          URLをコピー
        </Button>
      </BasicTemplate>
    </>
  );
}
