import { useState } from 'react';
import { Button } from '../../../../components/atoms/Button';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { MapForm } from '../../../../components/atoms/form/MapForm';
import { Textarea } from '../../../../components/atoms/form/Textarea';
import { TypoWrapper } from '../../../../components/atoms/TypoWrapper';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';

export default function Edit() {
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState(1);
  const [detail, setDetail] = useState('');
  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1 className="mt-5">編集する</h1>
      </TypoWrapper>

      <FormWrapper>
        <Input
          type="text"
          label="イベント名"
          content={name}
          changeContent={setName}
          required={true}
        />
        <Input
          type="number"
          label="募集人数"
          content={capacity}
          changeContent={setCapacity}
          required={true}
        />
        <Textarea
          label="詳細"
          content={detail}
          changeContent={setDetail}
          required={true}
        />
        <MapForm />
        <Button className="flex m-auto my-5" onClick={() => {}}>
          編集完了する
        </Button>
      </FormWrapper>

      {/* {created ?? (
        <Link href={`/event?id=${createdEventId}`}>募集したイベント</Link>
      )} */}
    </BasicTemplate>
  );
}
