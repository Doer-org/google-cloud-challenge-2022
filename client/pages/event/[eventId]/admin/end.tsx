import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../../components/atoms/TypoWrapper';
import { Button } from '../../../../components/atoms/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';

export default function End() {
  // ここは締め切るページ
  const [name, setName] = useState('');
  const [word, setWord] = useState('');
  const event_id = useRouter().query.id;
  return (
    <>
      <MyHead title="イベント締切ページ" description="" />
      <BasicTemplate className="text-center">
        <EventInfo />
        <Button className="flex m-auto my-5" onClick={() => {}}>
          締め切る
        </Button>
      </BasicTemplate>
    </>
  );
}
