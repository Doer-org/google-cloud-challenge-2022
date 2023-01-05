import { useState } from 'react';
import Link from 'next/link';
import useHostApi from '../../../core/hooks/useEventHost';
import { BasicTemplate } from '../../../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../../../components/atoms/TypoWrapper';
import { Input } from '../../../components/atoms/form/Input';
import { Textarea } from '../../../components/atoms/form/Textarea';
import { FormWrapper } from '../../../components/atoms/form/FormWrapper';
import { Button } from '../../../components/atoms/Button';

export default function New() {
  const { createNewEvent } = useHostApi();
  const [createdEventId, setCreatedEventId] = useState(0);
  const createEvent = createNewEvent((v) => {
    console.log(v);
    setCreated(true);
    setCreatedEventId(v.created_event.event_id);
  }, console.log);

  const [created, setCreated] = useState(false);
  const [name, setName] = useState('');
  const [capacity, setCapacity] = useState(1);
  const [detail, setDetail] = useState('');
  const [location, setLocation] = useState('');
  return (
    <BasicTemplate className="text-center">
      <TypoWrapper size="large" line="bold">
        <h1>募集する</h1>
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
        <Input
          type="text"
          label="場所"
          content={location}
          changeContent={setLocation}
        />
        <Button
          className="flex m-auto my-5"
          onClick={() => {
            // createEvent(
            //   { user_id: 'abc' },
            //   {
            //     event_name: name,
            //     max_member: capacity,
            //     detail: detail,
            //     location: location,
            //   }
            // );
          }}
        >
          募集する
        </Button>
      </FormWrapper>
      {/* {created ?? (
        <Link href={`/event?id=${createdEventId}`}>募集したイベント</Link>
      )} */}
    </BasicTemplate>
  );
}
