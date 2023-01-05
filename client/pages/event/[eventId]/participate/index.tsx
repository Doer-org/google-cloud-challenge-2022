import { Inter } from '@next/font/google';
import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../../components/atoms/TypoWrapper';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { Button } from '../../../../components/atoms/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';

export default function Participate() {
  // ここは参加者がみるただの参加フォーム
  // TODO : SSRで実装してリンクを貼った時にOGPを表示させるようにする
  const [name, setName] = useState('');
  const [word, setWord] = useState('');
  const event_id = useRouter().query.id;
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <BasicTemplate className="text-center">
        <EventInfo />
        <FormWrapper>
          <Input
            type="text"
            label="名前"
            content={name}
            changeContent={setName}
            required={true}
          />
          <Input
            type="text"
            label="ひとこと"
            content={word}
            changeContent={setWord}
            required={true}
          />
          <Button className="flex m-auto my-5" onClick={() => {}}>
            参加する
          </Button>
        </FormWrapper>
      </BasicTemplate>
    </>
  );
}
