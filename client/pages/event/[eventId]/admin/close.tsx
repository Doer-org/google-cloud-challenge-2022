import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../../components/atoms/text/TypoWrapper';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { closeEvent } from '../../../../core/api/event/close';

export default function Close() {
  // ここは締め切るページ
  // TODO:event締切hooksをonClickへ
  const [result, setResult] = useState('not closed');
  const close = closeEvent(
    (ok) => {
      setResult('ok!');
    },
    (err) => {
      setResult('error!');
    }
  );
  const [name, setName] = useState('');
  const [word, setWord] = useState('');

  const tmp = useRouter().query.id;
  const event_id = typeof tmp === 'undefined' || Array.isArray(tmp) ? '' : tmp;

  return (
    <>
      <MyHead title="イベント締切ページ" description="" />
      <BasicTemplate className="text-center">
        <EventInfo />
        <Button className="flex m-auto my-5" onClick={() => close(event_id)}>
          締め切る
        </Button>
        <>{result}</>
      </BasicTemplate>
    </>
  );
}
